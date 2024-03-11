package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

// You can modify this if you want to try to feed more data
const FeedSize = 3000

func FeedTask(es *elasticsearch.Client, tasks []model.Task) {
	// if index already exists do nothing
	exists, _ := es.Indices.Exists(
		[]string{"tasks"},
	)
	if exists != nil && exists.StatusCode == 200 {
		return
	}
	for i, task := range tasks {
		if i == FeedSize {
			break
		}
		go func(task model.Task) {
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
		}(task)
	}
}

func FeedShift(es *elasticsearch.Client, shifts []model.Shift) {
	exists, _ := es.Indices.Exists(
		[]string{"shifts"},
	)
	if exists != nil && exists.StatusCode == 200 {
		return
	}

	for i, shift := range shifts {
		if i == FeedSize {
			break
		}

		go func(shift model.Shift) {
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

			// Execute the indexing request
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Printf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("Error indexing shift ID %s: %s", shift.ID, res.String())
			} else {
				log.Printf("Shift ID %s indexed.", shift.ID)
			}
		}(shift)
	}
}

func FeedUser(es *elasticsearch.Client, users []model.User) {
	exists, _ := es.Indices.Exists(
		[]string{"users"},
	)

	if exists != nil && exists.StatusCode == 200 {
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
			log.Printf("Error getting response: %s", err)
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
	exists, _ := es.Indices.Exists(
		[]string{"slots"},
	)

	if exists != nil && exists.StatusCode == 200 {
		return
	}
	for i, slot := range slots {
		if i == FeedSize {
			break
		}
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
			log.Printf("Error getting response: %s", err)
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
	exists, _ := es.Indices.Exists(
		[]string{"orgas"},
	)

	if exists != nil && exists.StatusCode == 200 {
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
			log.Printf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error indexing orga ID %s: %s", org.ID, res.String())
		} else {
			log.Printf("Orga ID %s indexed.", org.ID)
		}
	}
}
