package handlers

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages/task"
)

type UpdateTaskRequest struct {
	TaskId int `json:"taskId"`
	Task struct {
		 Title       string `json:"title"`
		 Description *string `json:"description"`
		 Priority    *int64 `json:"priority"`
		 CompletedAt *time.Time `json:"completedAt"`
		 IsDeleted   bool `json:"isDeleted"`
		 IsCompleted bool `json:"isCompleted"`
	 } `json:"task"`
}

func(th *TaskHandler) Update(c *gin.Context) {
	ta := &UpdateTaskRequest{}
	if err := c.BindJSON(ta); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId: 53, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "An error occured while processing your request.",
		})
		return
	}

	taskUpdates := &task.Task{
		Title: ta.Task.Title,
		Description: ta.Task.Description,
		Priority: ta.Task.Priority,
		CompletedAt: ta.Task.CompletedAt,
		IsDeleted: ta.Task.IsDeleted,
		IsCompleted: ta.Task.IsCompleted,
	}

	err := th.Storage.Task().UpdateById(ta.TaskId, taskUpdates)

	if err != nil {
		if err == task.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, responses.ResponseError{
				ErrorCodeId: 22, // some fictional code
				DeveloperMessage: err.Error(),
				UserMessage: "Task you are looking for does not exists.",
			})
			return
		}

		c.JSON(http.StatusInternalServerError,  responses.ResponseError{
			ErrorCodeId: 60, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "An error occured while processing your request.",
		})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}