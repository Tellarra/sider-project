package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/adrichard/siderproject/internal/model"
	"github.com/elastic/go-elasticsearch/esapi"
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
	// get task

	tasks := GetDatasTasks(es, "tasks")
	taskToUpdate := []model.TaskToFeed{}
	for _, task := range tasks {
		if task.ID == taskId {
			task.AssigneeID = assigneeId
			taskToUpdate = append(taskToUpdate, task)
		}
	}
	if len(taskToUpdate) > 0 {
		UpdateInEs(taskToUpdate, es)
		return nil
	}
	/* index := slices.IndexFunc(tasks, func(t model.TaskToFeed) bool {
		return t.ID == taskId
	})

	if index != -1 {
		tasks[index].AssigneeID = assigneeId
		UpdateInEs(tasks[index], es)
	} */
	return errors.New("task not found")
}

func UpdateInEs(task []model.TaskToFeed, es *elasticsearch.Client) error {
	//var taskResponse model.TaskToFeed
	// general query
	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"id": task[0].ID,
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
			ID   string `json:"_id"`
			Hits []struct {
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
		taskToUpdate.AssigneeID = task[0].AssigneeID // Assuming first element in task has the updated ID

		// Construct the update request
		update := map[string]interface{}{
			"doc": taskToUpdate, // Include updated document in the request
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(update); err != nil {
			log.Printf("Error encoding query: %s", err)
			return err
		}
		readerUpdate := bytes.NewReader(buf.Bytes())

		updateReq := esapi.UpdateRequest{
			Index:      "tasks",
			DocumentID: r.Hits.ID, // Use the ID from the search result
			Body:       readerUpdate,
			Refresh:    "true",
		}

		// Execute the update request
		_, err := es.Update(updateReq.Index, updateReq.DocumentID, readerUpdate)
		if err != nil {
			log.Printf("Error updating document %s: %v", r.Hits.ID, err)
			// Handle individual update errors (optional)
		}
	}
	defer searchResponse.Body.Close()

	return nil
	/* // Définir l'ID du document à modifier
	docID := "123"

	// Définir les modifications à apporter
	update := map[string]interface{}{
		"nom": "Nouveau nom",
		"age": 30,
	}

	// Créer la requête de mise à jour
	req := esapi.UpdateRequest{
		Index:   index,
		Body:    strings.NewReader(json.Marshal(update)),
		Refresh: "true",
	}
	return nil */
}
