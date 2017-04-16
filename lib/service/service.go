package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/keycentral/keycentral/lib/api"
)

type KeyCentral struct {
	apiRouter *httprouter.Router
}

// NewKeyCentral instantiates all services
func NewKeyCentral() *KeyCentral {
	return &KeyCentral{
		apiRouter: api.NewAPIServer(),
	}
}

func (keycentral *KeyCentral) Start() {
	var domain = "localhost"
	var port = 8000
	var addr = fmt.Sprintf("%s:%d", domain, port)
	fmt.Println("[API Router] Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, keycentral.apiRouter))
}
