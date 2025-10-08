package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shortify/db"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var Routes = []Route{
	{Path: "/", Method: http.MethodGet, Handler: homeHandler},
	{Path: "/go/{id}", Method: http.MethodGet, Handler: urlRedirectHandler},
	{Path: "/api/url/shortify", Method: http.MethodPost, Handler: apiURLShortifyHandler},
}

func Init() {
	r := mux.NewRouter()
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from shortify")
}

func urlRedirectHandler(w http.ResponseWriter, r *http.Request) {
	inlineParams := mux.Vars(r)
	id := inlineParams["id"]
	url, err := db.GetURLFromID(id)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return

	}
	go db.IncrementURLClickCount(id)
	http.Redirect(w, r, url, http.StatusFound)
}

type ShortifyRequestBody struct {
	URL string `json:"url"`
}

type ShortifyResponse struct {
	URL string `json:"shorturl"`
}

func apiURLShortifyHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody ShortifyRequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "Request body: Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.URL == "" {
		http.Error(w, "Request body: url cannot be empty", http.StatusBadRequest)
		return
	}

	var id string
	id, err = db.InsertNewURL(requestBody.URL)

	if err != nil {
		http.Error(w, "Failed to shortify: "+err.Error(), http.StatusBadRequest)
		return
	}

	response := ShortifyResponse{URL: "http://localhost:8080/go/" + id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusBadRequest)
	}

}
