package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gustavocosta/controle-trafego-aereo/internal/service"
)

type VooController struct {
    service *service.VooService
}

func NewVooController(service *service.VooService) *VooController {
    return &VooController{service: service}
}

func (c *VooController) ListarVoos(ctx *gin.Context) {
    voos, err := c.service.ListarTodos()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar voos"})
        return
    }
    ctx.JSON(http.StatusOK, voos)
}