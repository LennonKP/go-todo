package main

import (
	"log"
	// "net/http"

    "github.com/gin-gonic/gin"

	"github.com/LennonKP/go-todo/src/handlers"
)

func main() {
    r := gin.Default()
    r.GET("/tarefas", handlers.GetTarefas)
    r.POST("/tarefas", handlers.AddTarefa)
    r.GET("/tarefas/:id", handlers.GetTarefaById)
    // r.PUT("/tarefas/:id", handlers.AlterTarefa)
    // r.DELETE("/tarefas/:id", handlers.DeleteTarefa)

	log.Println("Servidor iniciado na porta 8080")
	r.Run()
}
