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
		switch doc.Name {
		case "task":
			var tasks []model.Task
			err := json.Unmarshal(doc.Data, &tasks)
			if err != nil {
				log.Fatalf("Error unmarshalling task: %v", err)
			}
			FeedTask(es, tasks)
		case "orga":
			var orga []model.Orga
			err := json.Unmarshal(doc.Data, &orga)
			if err != nil {
				log.Fatalf("Error unmarshalling orga: %v", err)
			}
			FeedOrga(es, orga)
		case "shift":
			var shift []model.Shift
			err := json.Unmarshal(doc.Data, &shift)
			if err != nil {
				log.Fatalf("Error unmarshalling shift: %v", err)
			}
			FeedShift(es, shift)
		case "user":
			var user []model.User
			err := json.Unmarshal(doc.Data, &user)
			if err != nil {
				log.Fatalf("Error unmarshalling user: %v", err)
			}
			FeedUser(es, user)
		case "slots":
			var slots []model.Slot
			err := json.Unmarshal(doc.Data, &slots)
			if err != nil {
				log.Fatalf("Error unmarshalling slot: %v", err)
			}
			FeedSlot(es, slots)
		}

	}
	c.JSON(200, gin.H{
		"message": "Feed done",
	})
}

func FeedTask(es *elasticsearch.Client, tasks []model.Task) {
	// if index already exists do nothing
	exists, _ := es.Indices.Exists(
		[]string{"tasks"},
	)
	if exists.StatusCode == 200 {
		return
	}
	for _, task := range tasks {
		taskToFeed := task.ToTaskToFeed()
		jsonData, err := json.Marshal(taskToFeed)
		if err != nil {
			log.Fatalf("Error marshalling task: %v", err)
		}

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

func FeedShift(es *elasticsearch.Client, shifts []model.Shift) {
	if exists, _ := es.Indices.Exists(
		[]string{"shifts"},
	); exists.StatusCode == 200 {
		return
	}
	for _, shift := range shifts {
		shiftToFeed := shift.ToShiftToFeed()
		jsonData, err := json.Marshal(shiftToFeed)
		if err != nil {
			log.Fatalf("Error marshalling shift: %v", err)
		}

		req := esapi.IndexRequest{
			Index:   "shifts",
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
			log.Printf("Error indexing shift ID %s: %s", shift.ID, res.String())
		} else {
			log.Printf("Shift ID %s indexed.", shift.ID)
		}
	}
}

func FeedUser(es *elasticsearch.Client, users []model.User) {
	if exists, _ := es.Indices.Exists(
		[]string{"users"},
	); exists.StatusCode == 200 {
		return
	}
	for _, user := range users {
		userToFeed := user.ToUserToFeed()
		jsonData, err := json.Marshal(userToFeed)
		if err != nil {
			log.Fatalf("Error marshalling user: %v", err)
		}

		req := esapi.IndexRequest{
			Index:   "users",
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
			log.Printf("Error indexing user ID %s: %s", user.ID, res.String())
		} else {
			log.Printf("User ID %s indexed.", user.ID)
		}
	}
}

func FeedSlot(es *elasticsearch.Client, slots []model.Slot) {
	if exists, _ := es.Indices.Exists(
		[]string{"slots"},
	); exists.StatusCode == 200 {
		return
	}
	for _, slot := range slots {
		slotToFeed := slot.ToSlotToFeed()
		jsonData, err := json.Marshal(slotToFeed)
		if err != nil {
			log.Fatalf("Error marshalling slot: %v", err)
		}

		req := esapi.IndexRequest{
			Index:   "slots",
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
			log.Printf("Error indexing slot ID %s: %s", slot.ID, res.String())
		} else {
			log.Printf("Slot ID %s indexed.", slot.ID)
		}
	}
}

func FeedOrga(es *elasticsearch.Client, orga []model.Orga) {
	if exists, _ := es.Indices.Exists(
		[]string{"orgas"},
	); exists.StatusCode == 200 {
		return
	}
	for _, org := range orga {
		orgaToFeed := org.ToOrgaToFeed()
		jsonData, err := json.Marshal(orgaToFeed)
		if err != nil {
			log.Fatalf("Error marshalling orga: %v", err)
		}

		req := esapi.IndexRequest{
			Index:   "orgas",
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
			log.Printf("Error indexing orga ID %s: %s", org.ID, res.String())
		} else {
			log.Printf("Orga ID %s indexed.", org.ID)
		}
	}
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
	tasks, orgas, shifts, users, slots := domain.GetDatasTasks(es, "tasks"), domain.GetDatasOrgas(es, "orgas"), domain.GetDatasShifts(es, "shifts"), domain.GetDatasUsers(es, "users"), domain.GetDatasSlots(es, "slots")

	if query.FilterStatus {
		shifts = domain.FilterTasksByStatus(shifts)
	}
	response := domain.BuildResponse(tasks, orgas, shifts, users, slots)

	c.JSON(200, response)
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
