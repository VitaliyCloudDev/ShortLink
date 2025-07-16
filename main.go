package main

import router "shortlink/server"

func main() {
	router := router.SetupRouter()
	router.Run("localhost:8080")
}
