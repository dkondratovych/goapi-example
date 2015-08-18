package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages/task"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/version"
)

type AddTaskRequest struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Priority    *int       `json:"priority"`
	CompletedAt *time.Time `json:"completedAt"`
	IsDeleted   bool       `json:"isDeleted"`
	IsCompleted bool       `json:"isCompleted"`
}

func (th *TaskHandler) Add(c *gin.Context) {
	ta := &AddTaskRequest{}
	if err := c.BindJSON(ta); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId:      54, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage:      "An error occured whipe processing your request.",
		})
		return
	}

	st := &task.Task{
		Title:       ta.Title,
		Description: ta.Description,
		Priority:    ta.Priority,
		CompletedAt: ta.CompletedAt,
		IsDeleted:   ta.IsDeleted,
		IsCompleted: ta.IsCompleted,
	}

	err := th.Storage.Task().Add(st)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId:      60, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage:      "An error occured while processing your request.",
		})
		return
	}

	c.JSON(http.StatusCreated,
		responses.CommonResponse{
			Data: map[string]string{"url": fmt.Sprintf("/api/%s/tasks/%d", version.Version, st.Id)}})

}
