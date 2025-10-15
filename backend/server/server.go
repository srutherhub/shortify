package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shortify/db"
	"shortify/models"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const RedirectEndpoint = "http://localhost:8080/GO/"
const HankoURL = "https://7e09ca18-a6ff-46fc-b43a-4d35c0aa2cf2.hanko.io"

var HankoValidator = NewHankoSessionValidator(HankoURL)

var Routes = []models.Route{
	{Path: "/", Method: http.MethodGet, Handler: homeHandler},
	{Path: "/{route}/{id}", Method: http.MethodGet, Handler: urlRedirectHandler},
	{Path: "/api/url/shortify", Method: http.MethodPost, Handler: apiURLShortifyHandler},
	{Path: "/api/auth/validate", Method: http.MethodGet, Handler: AuthMiddleware(HankoValidator)(apiAuthValidateSessionHandler)},
}

func Init() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	var err error
	w.Header().Set("Content-Type", "application/json")

	switch d := data.(type) {
	case models.ServerErrorResponse:
		w.WriteHeader(d.Status)
	default:
		w.WriteHeader(http.StatusOK)
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusBadRequest)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from shortify")
}

func urlRedirectHandler(w http.ResponseWriter, r *http.Request) {
	inlineParams := mux.Vars(r)
	id := inlineParams["id"]
	route := inlineParams["route"]
	data, err := db.GetURL(id)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	if !strings.EqualFold(data.RouteName, route) {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	if !db.IsUrlActive(data.CreatedAt, data.ExpiresAt) {
		http.Redirect(w, r, "/", http.StatusNotFound)
		go db.DeleteExpireUrl(id, route)
		return
	}

	go db.IncrementURLClickCount(id)
	http.Redirect(w, r, data.Url, http.StatusFound)
}

func apiURLShortifyHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody models.ShortifyRequestBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&requestBody)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	if requestBody.URL == "" {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: "url is required", Status: http.StatusBadRequest})
		return
	}

	var id string
	id, err = db.InsertNewURL(requestBody)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	response := models.ShortifyResponse{URL: RedirectEndpoint + id}
	JSONResponse(w, response)
}

func apiAuthValidateSessionHandler(w http.ResponseWriter, r *http.Request) {

}
