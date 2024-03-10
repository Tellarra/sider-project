package ui

import (
	"encoding/json"

	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/model"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context, files []model.DocumentToIndex) {
	var query model.Query
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error parsing params",
		})
		return
	}
	tasks, orga, shift, user, slots := unmarshalAll(files)
	if query.FilterStatus {
		shift = domain.FilterTasksByStatus(shift)
	}
	response := domain.BuildResponse(tasks, orga, shift, user, slots)

	c.JSON(200, response)
}

func UpdateAssigneeID(c *gin.Context, files []model.DocumentToIndex) {
	taskId := c.Param("id")
	type AssigneeID struct {
		AssigneeID string `json:"assigneeId"`
	}
	var assigneeId AssigneeID
	err := c.BindJSON(&assigneeId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error parsing body",
		})
		return
	}

	// update task
	// get task
	// update assignee
	// save task
	// return response

	tasks, _, _, _, _ := unmarshalAll(files)

	errUpdate := domain.UpdateTask(assigneeId.AssigneeID, taskId, tasks)
	if errUpdate != nil {
		c.JSON(400, gin.H{
			"message": "Error updating task",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Task updated"})
}

func unmarshalAll(files []model.DocumentToIndex) ([]model.Task, []model.Orga, []model.Shift, []model.User, []model.Slot) {
	var tasks []model.Task
	var orga []model.Orga
	var shift []model.Shift
	var user []model.User
	var slots []model.Slot

	for _, file := range files {
		switch file.Name {
		case "task":
			json.Unmarshal(file.Data, &tasks)
		case "orga":
			json.Unmarshal(file.Data, &orga)
		case "shift":
			json.Unmarshal(file.Data, &shift)
		case "user":
			json.Unmarshal(file.Data, &user)
		case "slots":
			json.Unmarshal(file.Data, &slots)
		}
	}
	return tasks, orga, shift, user, slots
}
