package cmd

func main() {
	deps := InitDependencies()
	router := InitRoutes(deps)

	StartServer(router)
}
