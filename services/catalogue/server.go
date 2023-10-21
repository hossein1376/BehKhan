package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/hossein1376/BehKhan/catalogue/internal/handlers"
	"github.com/hossein1376/BehKhan/catalogue/proto/cataloguePB"
)

func doGrpc() {
	lis, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	cataloguePB.RegisterBookServiceServer(s, &handlers.Server{})

	log.Println("starting grpc server...")

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func doHttp() {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("starting HTTP server...")

	err := r.Run(":8002")
	if err != nil {
		log.Fatal(err)
	}
}
