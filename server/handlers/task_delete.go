package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages/task"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
)

type DeleteTaskRequest struct {
	Id int `json:"id" binding:"required"`
}

func(th *TaskHandler) Delete(c *gin.Context) {
	td := &DeleteTaskRequest{}

	if err := c.BindJSON(td); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId: 54, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "An error occured whipe processing your request.",
		})
		return
	}

	err := th.Storage.Task().DeleteById(td.Id)

	if err == task.ErrTaskNotFound {
		c.JSON(http.StatusNotFound, responses.ResponseError{
			ErrorCodeId: 22, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "Task you are looking for does not exists.",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError,  responses.ResponseError{
			ErrorCodeId: 55, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "An error occured while processing your request.",
		})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}