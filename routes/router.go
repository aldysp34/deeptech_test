package routes

import (
	"github.com/gorilla/mux"
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

}
