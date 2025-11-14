package server

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"shortify/db"
	"shortify/models"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

var _ = godotenv.Load()

var ServerUrl = os.Getenv("SERVER_URL")
var RedirectEndpoint = ServerUrl + "/s/" + "go/"

const HankoURL = "https://7e09ca18-a6ff-46fc-b43a-4d35c0aa2cf2.hanko.io"

var HankoValidator = NewHankoSessionValidator(HankoURL)

var Routes = []models.Route{
	{Path: "/s/{route}/{id}", Method: http.MethodGet, Handler: urlRedirectHandler},
	{Path: "/api/url/shortify", Method: http.MethodPost, Handler: AuthMiddleware(HankoValidator)(apiURLShortifyHandler), IsPublic: true},
	{Path: "/api/url/getuserurls", Method: http.MethodGet, Handler: AuthMiddleware(HankoValidator)(apiURLGetUserUrlsHandler)},
	{Path: "/api/url/getquickstats", Method: http.MethodGet, Handler: AuthMiddleware(HankoValidator)(apiURLGetQuickStatsHandler)},
	{Path: "/api/auth/validate", Method: http.MethodGet, Handler: AuthMiddleware(HankoValidator)(apiAuthValidateSessionHandler)},
}

func Init() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:5173", "http://192.168.1.66:5173", "https://192.168.1.66:5173", "https://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	log.Fatal(http.ListenAndServe(":5555", handler))
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
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
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

	var username string
	username, err = GetUserFromCookie(r)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	var id string
	id, err = db.InsertNewURL(requestBody, username)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	response := models.ShortifyResponse{URL: RedirectEndpoint + id}
	JSONResponse(w, response)
}

func apiAuthValidateSessionHandler(w http.ResponseWriter, r *http.Request) {

}

func apiURLGetUserUrlsHandler(w http.ResponseWriter, r *http.Request) {
	var username string
	var err error
	username, err = GetUserFromCookie(r)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	urls, err := db.GetUserUrls(username)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	JSONResponse(w, urls)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func apiURLGetQuickStatsHandler(w http.ResponseWriter, r *http.Request) {
	username, err := GetUserFromCookie(r)

	if err != nil {
		JSONResponse(w, models.ServerErrorResponse{Error: models.EInvalidRequest, Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	var g errgroup.Group
	var totalUrlCount int64
	var totalClickCount int64

	g.Go(func() error {
		var err error
		totalUrlCount, err = db.GetTotalNumUrls(username)
		return err
	})

	g.Go(func() error {
		var err error
		totalClickCount, err = db.GetTotalNumClickCount(username)
		return err
	})

	if err := g.Wait(); err != nil {
		JSONResponse(w, models.ServerErrorResponse{
			Error:   models.EAppError,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	var avgClicksPerUrl float64
	if totalUrlCount == 0 {
		avgClicksPerUrl = 0
	} else {
		avgClicksPerUrl = math.Round((float64(totalClickCount) / float64(totalUrlCount) * 100)) / 100
	}

	res := models.QuickStatsResponse{Total_Urls: totalUrlCount, Total_Click_Count: totalClickCount, Average_Clicks_Per_Url: avgClicksPerUrl}

	JSONResponse(w, res)
}
