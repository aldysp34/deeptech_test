package routes

import (
	"log"
	"net/http"

	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Routes struct {
	Router *mux.Router
	Port   string
}

func (a *Routes) Init() {
	a.Router = mux.NewRouter()
	a.Router.StrictSlash(false)
	a.Router.PathPrefix("/").Subrouter()

	a.ListRouter()

	negroWare := negroni.New()

	//* PANIC Recovery
	negroWare.Use(negroni.NewRecovery())

	negroWare.UseHandler(a.Router)

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000

	origins := muxHandlers.AllowedOrigins([]string{"*"})                                                // All origins
	methods := muxHandlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"}) // Allowing only get, just an example
	headers := muxHandlers.AllowedHeaders([]string{"access_token", "x-diug-ofni", "device_id", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "data_type"})

	log.Fatal(http.ListenAndServe(":"+a.Port, muxHandlers.CORS(origins, methods, headers)(negroWare)))
}
