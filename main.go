package main

import (
	"log"
	"publisher/infraestructure/routes"
	"github.com/gin-gonic/gin"
)


func main() {

router := gin.Default()

	routes.SetUpRoutes(router)

	port :=":8081"
	log.Println("Servidor corriendo en el puerto", port)
	log.Fatal(router.Run(port))


}
