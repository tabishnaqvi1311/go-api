package main

import (
	"log"
	"context"

	"./config"
	"github/gin-gonic/gin"
)

func main(){
	client, err := config.connectMongo()

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	router.Run("localhost:8000")
}