package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages/task"
)

func (th *TaskHandler) Find(c *gin.Context) {
	sid := c.Param("id")

	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId:      57, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage:      "Invalid task id.",
		})
		return
	}

	ts, err := th.Storage.Task().FindById(id)

	if err == task.ErrTaskNotFound {
		c.JSON(http.StatusNotFound, responses.ResponseError{
			ErrorCodeId:      22, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage:      "Task you are looking for does not exists.",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId:      55, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage:      "An error occured while processing your request.",
		})
		return
	}

	c.JSON(http.StatusOK, responses.CommonResponse{Data: ts})
}
