package main

import "go-ecommerce-be-api/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"). If you want to change the port, you can change it here. r.Run(":8082")
}
