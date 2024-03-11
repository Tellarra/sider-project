package ui

import (
	"encoding/json"
	"log"

	"github.com/adrichard/siderproject/infrastructure"
	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context, es *elasticsearch.Client) {
	var query model.Query
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error parsing params",
		})
		return
	}
	tasks, orgas, shifts, users, slots := domain.GetDatasTasks(es, "tasks"), domain.GetDatasOrgas(es, "orgas"), domain.GetDatasShifts(es, "shifts"), domain.GetDatasUsers(es, "users"), domain.GetDatasSlots(es, "slots")

	if query.FilterStatus {
		shifts = domain.FilterTasksByStatus(shifts)
	}
	response := domain.BuildResponse(tasks, orgas, shifts, users, slots)

	c.JSON(200, response)
}

func Feed(c *gin.Context, es *elasticsearch.Client) {
	docs, err := domain.GetAllFilesToRightTypes()
	if err != nil {
		log.Fatal("Could not get files ", err)
	}
	c.JSON(200, gin.H{
		"message": "Starting feed",
	})
	for _, doc := range docs {
		switch doc.Name {
		case "task":
			var tasks []model.Task
			err := json.Unmarshal(doc.Data, &tasks)
			if err != nil {
				log.Fatalf("Error unmarshalling task: %v", err)
			}
			infrastructure.FeedTask(es, tasks)
		case "orga":
			var orga []model.Orga
			err := json.Unmarshal(doc.Data, &orga)
			if err != nil {
				log.Fatalf("Error unmarshalling orga: %v", err)
			}
			infrastructure.FeedOrga(es, orga)
		case "shift":
			var shift []model.Shift
			err := json.Unmarshal(doc.Data, &shift)
			if err != nil {
				log.Fatalf("Error unmarshalling shift: %v", err)
			}
			infrastructure.FeedShift(es, shift)
		case "user":
			var user []model.User
			err := json.Unmarshal(doc.Data, &user)
			if err != nil {
				log.Fatalf("Error unmarshalling user: %v", err)
			}
			infrastructure.FeedUser(es, user)
		case "slots":
			var slots []model.Slot
			err := json.Unmarshal(doc.Data, &slots)
			if err != nil {
				log.Fatalf("Error unmarshalling slot: %v", err)
			}
			infrastructure.FeedSlot(es, slots)
		}
	}
	c.JSON(200, gin.H{
		"message": "Feed done",
	})
}

func UpdateAssigneeID(c *gin.Context, es *elasticsearch.Client) {
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

	errUpdate := domain.UpdateTask(assigneeId.AssigneeID, taskId, es)
	if errUpdate != nil {
		c.JSON(400, gin.H{
			"message": "Error updating task",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":    "Task updated",
		"assigneeId": assigneeId.AssigneeID,
		"taskId":     taskId,
	})
}
