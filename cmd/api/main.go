package main

import (
	_ "github.com/antonioalfa22/go-rest-template/docs"
	"github.com/antonioalfa22/go-rest-template/internal/api"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
