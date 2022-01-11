package app

import (
	"poke_api/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", controller.GetPokemon)
	return r
}
