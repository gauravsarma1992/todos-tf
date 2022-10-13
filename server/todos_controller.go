package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (apiServer *ApiServer) TodoShow(c *gin.Context) {
	var (
		todoItem TodoItem
		todoId   int
		err      error
	)

	reqTodoId := c.Param("id")
	if todoId, err = strconv.Atoi(reqTodoId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Todo Item not found",
		})
		return
	}

	if result := apiServer.db.First(&todoItem, todoId); result.Error != nil || todoItem.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Todo Item not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"todo": todoItem,
	})
	return
}

func (apiServer *ApiServer) TodoCreate(c *gin.Context) {
	var (
		todoItem TodoItem
		err      error
	)
	todoItem = TodoItem{}
	if err = c.ShouldBindJSON(&todoItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := apiServer.db.Create(&todoItem); result.Error != nil || todoItem.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Todo Item not created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"todo": todoItem,
	})
	return
}
