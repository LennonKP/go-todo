package handlers

import (
	// "encoding/json"

	"fmt"
	"net/http"
	"strconv"

	"github.com/LennonKP/go-todo/src/models"
	"github.com/gin-gonic/gin"
)

var tarefas = []models.Tarefa{
	{ID: 1, Descricao: "Descrição 1", Concluida: false},
}

//	func GetAllTarefas(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(tarefas)
//	}
func GetTarefas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tarefas)
}

func AddTarefa(c *gin.Context) {
	var newTarefa models.Tarefa

	if err := c.BindJSON(&newTarefa); err != nil {
		return
	}

	tarefas = append(tarefas, newTarefa)
	c.IndentedJSON(http.StatusCreated, newTarefa)
}

func GetTarefaById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID", err)
		return
	}

	for _, tarefa := range tarefas {
		if tarefa.ID == id {
			c.IndentedJSON(http.StatusOK, tarefa)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
}

func AlterTarefa(c *gin.Context) {
	c.JSON(200, tarefas)
}
func DeleteTarefa(c *gin.Context) {
	c.JSON(200, tarefas)
}
