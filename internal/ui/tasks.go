package ui

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/adrichard/siderproject/internal/domain"
	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context, es *elasticsearch.Client) {
	docs, err := domain.GetAllFilesToRightTypes()
	if err != nil {
		log.Fatal("Could not get files ", err)
	}

	for _, doc := range docs {
		if doc.Name != "task" {
			continue
		}

		var tasks []model.Task
		err := json.Unmarshal(doc.Data, &tasks)
		if err != nil {
			log.Fatalf("Error unmarshalling task: %v", err)
		}

		for _, task := range tasks {
			// Convertit l'objet task en JSON pour l'indexation
			taskToFeed := task.ToTaskToFeed()
			jsonData, err := json.Marshal(taskToFeed)
			if err != nil {
				log.Fatalf("Error marshalling task: %v", err)
			}

			// Créez une requête d'indexation pour chaque document
			req := esapi.IndexRequest{
				Index:   "tasks",
				Body:    bytes.NewReader(jsonData),
				Refresh: "true",
			}

			// Exécutez la requête d'indexation
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("Error indexing task ID %s: %s", task.ID, res.String())
			} else {
				log.Printf("Task ID %s indexed.", task.ID)
			}
		}
	}
	c.JSON(200, gin.H{
		"message": "Feed done",
	})
}

func GetTasks(c *gin.Context, es *elasticsearch.Client) {
	var query model.Query
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error parsing params",
		})
		return
	}
	tasks, orga, shift, user, slots := unmarshalAll(nil)
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
