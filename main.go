package main

import (
	"fmt"
	"healthtrack/db"
	"healthtrack/server"
	"os"
)

func main() {
	db.Connect()
	//server.SetUpRoute(db.DB)
	r := server.SetUpRoute(db.DB)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8090"
	}
	r.Run(":" + port)
	fmt.Printf("Server Runing at port % s", port)
	r.Run(port)
}
