package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("GET /")
	fmt.Fprint(w, "KeyCentral!\n")
}

func NewAPIServer() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/register", Register)

	return router
}
