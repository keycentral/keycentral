package common

import (
	"fmt"

	"github.com/keycentral/keycentral/lib/service"

	flag "github.com/ogier/pflag"
)

// Run KeyCentral server
func Run() {
	fmt.Println("Keycentral!")

	var dbg bool

	flag.BoolVarP(&dbg, "debug", "d", false, "Enable debug mode")
	flag.Parse()

	if dbg {
		fmt.Println("Debug mode")
	}

	keycentral := service.NewKeyCentral()
	keycentral.Start()
}
