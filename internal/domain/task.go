package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"slices"
	"time"

	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/v8"
)

func GetDatas(es *elasticsearch.Client, indexName []string) ([]model.TaskToFeed, []model.OrgaToFeed, []model.ShiftToFeed, []model.UserToFeed, []model.SlotToFeed) {
	var taskResponse []model.TaskToFeed

	// Define the search query (replace with your specific query)
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{}, // Match all documents
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	r := bytes.NewReader(buf.Bytes())

	for _, index := range indexName {
		searchResponse, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex(index),
			es.Search.WithBody(r), // Use the byte slice reader
		)
		if err != nil {
			log.Fatalf("Error searching Elasticsearch: %v", err)
		}

		// Handle the search response
		if searchResponse.IsError() {
			var error map[string]interface{}
			if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
				log.Fatalf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and error information
				log.Fatalf("[%s] Error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			}
		}

		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		log.Print(r)

		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			var task model.TaskToFeed
			source := hit.(map[string]interface{})["_source"]
			data, err := json.Marshal(source)
			if err != nil {
				log.Fatalf("Error marshalling task: %v", err)
			}
			err = json.Unmarshal(data, &task)
			if err != nil {
				log.Fatalf("Error unmarshalling task: %v", err)
			}
			taskResponse = append(taskResponse, task)
		}

		defer searchResponse.Body.Close()
	}

	return taskResponse, nil, nil, nil, nil
}

func FilterTasksByStatus(shifts []model.Shift) []model.Shift {
	var filteredTasks []model.Shift
	for _, shift := range shifts {
		if shift.Status == "cancelled" {
			continue
		}
		// end date is before the time.Now()
		// means it's in the past
		if shift.Time.EndDate.String() < time.Now().String() {
			shift.Status = "done"
		} else if shift.Time.StartDate.String() < time.Now().String() {
			shift.Status = "ongoing"
		} else {
			shift.Status = "upcoming"
		}
		filteredTasks = append(filteredTasks, shift)
	}
	return filteredTasks
}

func BuildResponse(tasks []model.Task, orga []model.Orga, shift []model.Shift, user []model.User, slots []model.Slot) []model.TaskResponse {
	var response []model.TaskResponse
	for _, task := range tasks {
		orga := model.GetOrga(task.OrganisationID, orga)
		shift := model.GetShift(task.ID, task.ShiftIDs, shift)
		response = append(response, model.TaskResponse{
			Id:   task.ID,
			Name: task.Alias,
			Orga: model.OrgaResponse{
				Name:       orga.Name,
				Address:    orga.Address,
				PictureUrl: orga.PictureUrl,
			},
			Shift: model.ShiftResponse{
				Id:        shift.Id,
				StartDate: shift.StartDate,
				EndDate:   shift.EndDate,
				Slots: model.SlotsResponse{
					Filled: shift.Slots.Filled,
					Total:  shift.Slots.Total,
				},
				Status: shift.Status,
			},
		})
	}
	return response
}

func UpdateTask(assigneeId string, taskId string, tasks []model.Task) error {
	index := slices.IndexFunc(tasks, func(t model.Task) bool {
		return t.ID == taskId
	})
	if index != -1 {
		tasks[index].AssigneeID = assigneeId
		return nil
	}
	return errors.New("task not found")
}
