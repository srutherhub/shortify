package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"shortify/db"

	"github.com/gorilla/mux"
)

type ServerError string

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type ServerErrorResponse struct {
	Error   ServerError `json:"error"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

type ShortifyRequestBody struct {
	URL string `json:"url"`
}

type ShortifyResponse struct {
	URL string `json:"shorturl"`
}

const (
	EInvalidRequest ServerError = "INVALID_REQUEST"
	EAppError       ServerError = "APP_ERROR"
	EDatabaseError  ServerError = "DATABASE_ERROR"
)

const RedirectEndpoint = "http://localhost:8080/go/"

var Routes = []Route{
	{Path: "/", Method: http.MethodGet, Handler: homeHandler},
	{Path: "/go/{id}", Method: http.MethodGet, Handler: urlRedirectHandler},
	{Path: "/api/url/shortify", Method: http.MethodPost, Handler: apiURLShortifyHandler},
}

// func MiddleWareExample(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Hello from middle ware")
// 		handler(w, r)
// 	}
// }

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

func apiURLShortifyHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody ShortifyRequestBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&requestBody)

	if err != nil {
		JSONResponse(w, ServerErrorResponse{Error: EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	if requestBody.URL == "" {
		JSONResponse(w, ServerErrorResponse{Error: EInvalidRequest, Message: "url is required", Status: http.StatusBadRequest})
		return
	}

	var id string
	id, err = db.InsertNewURL(requestBody.URL)

	if err != nil {
		JSONResponse(w, ServerErrorResponse{Error: EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	response := ShortifyResponse{URL: RedirectEndpoint + id}
	JSONResponse(w, response)
}

func JSONResponse(w http.ResponseWriter, data any) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	dataType := reflect.TypeOf(data).Name()

	if dataType == "ServerError" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusBadRequest)
		return
	}
}
