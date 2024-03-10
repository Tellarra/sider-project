package ui

import (
	"encoding/json"

	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/model"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context, files []model.DocumentToIndex) {
	err := c.ShouldBindQuery(&model.Query{})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error parsing params",
		})
		return
	}

	var tasks []model.Task
	var orga []model.Orga
	var shift []model.Shift
	var user []model.User
	var slots []model.Slot

	for _, file := range files {
		switch file.Name {
		case "task":
			err := json.Unmarshal(file.Data, &tasks)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error unmarshalling file",
				})
				return
			}
		case "orga":
			err := json.Unmarshal(file.Data, &orga)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error unmarshalling file",
				})
				return
			}
		case "shift":
			err := json.Unmarshal(file.Data, &shift)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error unmarshalling file",
				})
				return
			}
		case "user":
			err := json.Unmarshal(file.Data, &user)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error unmarshalling file",
				})
				return
			}
		case "slots":
			err := json.Unmarshal(file.Data, &slots)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error unmarshalling file",
				})
				return
			}
		}
	}

	response := domain.BuildResponse(tasks, orga, shift, user, slots)

	c.JSON(200, response)
}
