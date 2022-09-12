package main

import (
	"github.com/redbeestudios/go-seed/cmd"
)

func main() {
	deps := cmd.InitDependencies()
	router := cmd.InitRoutes(deps)
	router.Get("/pokemon/{id}")
	cmd.StartServer(router)
}
