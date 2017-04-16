package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("POST /register")
	fmt.Fprintf(w, "Register new user")
}
