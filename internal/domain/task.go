package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/v8"
)

func GetDatasTasks(es *elasticsearch.Client, indexName string) []model.TaskToFeed {
	var taskResponse []model.TaskToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{}, // Match all documents
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	readerTask := bytes.NewReader(buf.Bytes())

	// To deserialize the response, we need a struct that matches the data in the JSON response.
	type ResponseES struct {
		Hits struct {
			Hits []struct {
				Source model.TaskToFeed `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	searchResponse, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(readerTask),
		es.Search.WithSize(1000),
	)
	if err != nil {
		log.Fatalf("Error searching Elasticsearch: %v", err)
	}

	if searchResponse.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
			log.Printf("error parsing the response body: %s", err)
			return nil
		} else {
			log.Printf("[%s] error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			return nil
		}
	}

	// Deserialize the response into a map.
	var r ResponseES
	if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil
	}

	for _, hit := range r.Hits.Hits {
		taskResponse = append(taskResponse, hit.Source)
	}
	defer searchResponse.Body.Close()

	return taskResponse
}

func GetDatasOrgas(es *elasticsearch.Client, indexName string) []model.OrgaToFeed {
	var orgaResponse []model.OrgaToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{}, // Match all documents
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	readerOrga := bytes.NewReader(buf.Bytes())

	// To deserialize the response, we need a struct that matches the data in the JSON response.
	type ResponseES struct {
		Hits struct {
			Hits []struct {
				Source model.OrgaToFeed `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	searchResponse, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(readerOrga),
		es.Search.WithSize(1000),
	)
	if err != nil {
		log.Fatalf("Error searching Elasticsearch: %v", err)
	}

	if searchResponse.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
			log.Printf("error parsing the response body: %s", err)
			return nil
		} else {
			log.Printf("[%s] error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			return nil
		}
	}

	// Deserialize the response into a map.
	var r ResponseES
	if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil
	}

	for _, hit := range r.Hits.Hits {
		orgaResponse = append(orgaResponse, hit.Source)
	}
	defer searchResponse.Body.Close()

	return orgaResponse
}

func GetDatasShifts(es *elasticsearch.Client, indexName string) []model.ShiftToFeed {
	var shiftResponse []model.ShiftToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{}, // Match all documents
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	readerShift := bytes.NewReader(buf.Bytes())

	// To deserialize the response, we need a struct that matches the data in the JSON response.
	type ResponseES struct {
		Hits struct {
			Hits []struct {
				Source model.ShiftToFeed `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	searchResponse, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(readerShift),
		es.Search.WithSize(1000),
	)
	if err != nil {
		log.Fatalf("Error searching Elasticsearch: %v", err)
	}

	if searchResponse.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
			log.Printf("error parsing the response body: %s", err)
			return nil
		} else {
			log.Printf("[%s] error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			return nil
		}
	}

	// Deserialize the response into a map.
	var r ResponseES
	if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil
	}

	for _, hit := range r.Hits.Hits {
		shiftResponse = append(shiftResponse, hit.Source)
	}
	defer searchResponse.Body.Close()

	return shiftResponse
}

func GetDatasUsers(es *elasticsearch.Client, indexName string) []model.UserToFeed {
	var userResponse []model.UserToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{}, // Match all documents
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	readerUser := bytes.NewReader(buf.Bytes())

	// To deserialize the response, we need a struct that matches the data in the JSON response.
	type ResponseES struct {
		Hits struct {
			Hits []struct {
				Source model.UserToFeed `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	searchResponse, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(readerUser),
		es.Search.WithSize(1000),
	)
	if err != nil {
		log.Fatalf("Error searching Elasticsearch: %v", err)
	}

	if searchResponse.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
			log.Printf("error parsing the response body: %s", err)
			return nil
		} else {
			log.Printf("[%s] error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			return nil
		}
	}

	// Deserialize the response into a map.
	var r ResponseES
	if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil
	}

	for _, hit := range r.Hits.Hits {
		userResponse = append(userResponse, hit.Source)
	}
	defer searchResponse.Body.Close()

	return userResponse
}

func GetDatasSlots(es *elasticsearch.Client, indexName string) []model.SlotToFeed {
	var slotResponse []model.SlotToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{}, // Match all documents
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return nil
	}
	readerSlot := bytes.NewReader(buf.Bytes())

	// To deserialize the response, we need a struct that matches the data in the JSON response.
	type ResponseES struct {
		Hits struct {
			Hits []struct {
				Source model.SlotToFeed `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	searchResponse, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(readerSlot),
		es.Search.WithSize(1000),
	)
	if err != nil {
		log.Fatalf("Error searching Elasticsearch: %v", err)
	}

	if searchResponse.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
			log.Printf("error parsing the response body: %s", err)
			return nil
		} else {
			log.Printf("[%s] error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			return nil
		}
	}

	// Deserialize the response into a map.
	var r ResponseES
	if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return nil
	}

	for _, hit := range r.Hits.Hits {
		slotResponse = append(slotResponse, hit.Source)
	}
	defer searchResponse.Body.Close()

	return slotResponse
}

func FilterTasksByStatus(shifts []model.ShiftToFeed) []model.ShiftToFeed {
	var filteredTasks []model.ShiftToFeed
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

func BuildResponse(tasks []model.TaskToFeed, orga []model.OrgaToFeed, shift []model.ShiftToFeed, user []model.UserToFeed, slots []model.SlotToFeed) []model.TaskResponse {
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

func UpdateTask(assigneeId string, taskId string, es *elasticsearch.Client) error {

	//var taskResponse model.TaskToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"id": taskId,
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return err
	}
	readerTask := bytes.NewReader(buf.Bytes())

	// To deserialize the response, we need a struct that matches the data in the JSON response.
	type ResponseES struct {
		Hits struct {
			Hits []struct {
				ID     string           `json:"_id"`
				Source model.TaskToFeed `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	searchResponse, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("tasks"),
		es.Search.WithBody(readerTask),
	)
	if err != nil {
		log.Fatalf("Error searching Elasticsearch: %v", err)
	}

	if searchResponse.IsError() {
		var error map[string]interface{}
		if err := json.NewDecoder(searchResponse.Body).Decode(&error); err != nil {
			log.Printf("error parsing the response body: %s", err)
			return err
		} else {
			log.Printf("[%s] error searching index: %s", searchResponse.Status(), error["error"].(map[string]interface{})["reason"])
			return err
		}
	}

	// Deserialize the response into a map.
	var r ResponseES
	if err := json.NewDecoder(searchResponse.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return err
	}

	for _, hit := range r.Hits.Hits {
		taskToUpdate := hit.Source // Get the source document

		// Update assigneeId in the source document
		taskToUpdate.AssigneeID = assigneeId // Assuming first element in task has the updated ID

		// Construct the update request
		update := map[string]interface{}{
			"doc": map[string]interface{}{
				"assigneeId": taskToUpdate.AssigneeID,
			},
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(update); err != nil {
			log.Printf("Error encoding query: %s", err)
			return err
		}
		readerUpdate := bytes.NewReader(buf.Bytes())
		resp, err := es.Update(
			"tasks",
			hit.ID,
			readerUpdate,
			es.Update.WithRefresh("true"),
		)

		if err != nil {
			log.Printf("Error updating document %s: %v", hit.ID, err)
		}
		if resp.IsError() {
			log.Printf("Error updating document %s: %s", hit.ID, resp.Status())
		} else {
			log.Printf("Document %s updated.", hit.ID)
		}
		defer resp.Body.Close()
	}
	defer searchResponse.Body.Close()

	return nil
}
