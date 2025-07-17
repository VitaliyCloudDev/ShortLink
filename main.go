package main

import server "shortlink/server"

func main() {
	router := server.SetupRouter()
	router.Run("localhost:8080")
}
