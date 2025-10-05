package server

import (
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
}

func Init() {
	r := mux.NewRouter()
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler)
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
		panic(err)
	}
	http.Redirect(w, r, url, http.StatusFound)
}
