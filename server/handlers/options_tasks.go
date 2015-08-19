package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TasksOptions map[string]Method

var taskParameters = Parameters{
	Parameter{
		"title": ParameterOption{
			"type":        "string",
			"description": "Task title",
			"required":    false,
		},
	},
	Parameter{
		"description": ParameterOption{
			"type":        "string",
			"description": "Task decription",
			"required":    false,
		},
	},
	Parameter{
		"priority": ParameterOption{
			"type":        "int",
			"description": "Task priority",
			"required":    false,
		},
	},
	Parameter{
		"completedAt": ParameterOption{
			"type":        "datetime",
			"description": "Task completion date. Should be formatted according RFC3339.",
			"required":    false,
		},
	},
	Parameter{
		"isDeleted": ParameterOption{
			"type":        "boolean",
			"description": "Indicate if task is deleted.",
			"required":    false,
		},
	},
	Parameter{
		"isCompleted": ParameterOption{
			"type":        "boolean",
			"description": "Indicate if task is completed.",
			"required":    false,
		},
	},
}

func (o *OptionsHandler) Tasks(c *gin.Context) {
	c.Header("Allow", "POST, PUT, DELETE")

	c.JSON(http.StatusCreated, TasksOptions{
		"POST": Method{
			Description: "Creates new task.",
			Parameters:  taskParameters,
			Example: map[string]interface{}{
				"Title":       "Test task",
				"Description": "Test description",
				"Priority":    1,
				"CompletedAt": time.Now().Format(time.RFC3339),
				"IsDeleted":   false,
				"IsCompleted": true,
			},
		},
		"PUT": Method{
			Description: "Updates task fields by task ID",
			Parameters: Parameters{
				Parameter{
					"taskId": ParameterOption{
						"type":        "int",
						"description": "Task id to update.",
						"required":    true,
					},
					"task": taskParameters,
				},
			},
			Example: map[string]interface{}{
				"taskId": 1,
				"task": map[string]interface{}{
					"Title":       "Test task",
					"Description": "Test description",
					"Priority":    1,
					"CompletedAt": time.Now().Format(time.RFC3339),
					"IsDeleted":   false,
					"IsCompleted": true,
				},
			},
		},
		"DELETE": Method{
			Description: "Mar task as deleted by task ID",
			Parameters: Parameters{
				Parameter{
					"id": ParameterOption{
						"type":        "int",
						"description": "Task ID to delete",
						"required":    true,
					},
				},
			},
			Example: map[string]int{
				"id": 1,
			},
		},
	})
}

func (o *OptionsHandler) TaskById(c *gin.Context) {
	c.Header("Allow", "GET")
}
