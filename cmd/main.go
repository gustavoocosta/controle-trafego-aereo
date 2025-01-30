package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gustavocosta/controle-trafego-aereo/config"
    "github.com/gustavocosta/controle-trafego-aereo/internal/controller"
    "github.com/gustavocosta/controle-trafego-aereo/internal/repository"
    "github.com/gustavocosta/controle-trafego-aereo/internal/service"
)

func main() {
    config.InitDB()

    repo := repository.NewVooRepository(config.DB)
    service := service.NewVooService(repo)
    controller := controller.NewVooController(service)

    router := gin.Default()

    router.GET("/voos", controller.ListarVoos)

    router.Run(":8080")
}