package request

import (
	"fmt"
	"bytes"
	"errors"
	"net/http"
	"encoding/json"
)

func GetListTaskLists(webClient *http.Client) (*ListTaskLists, error) {
	req, requestErr := http.NewRequest(
					http.MethodGet,
					"https://graph.microsoft.com/beta/me/todo/lists",
					nil)

	if requestErr != nil {
		return &ListTaskLists{}, requestErr
	}

	response, responseErr := webClient.Do(req)
	if responseErr != nil {
		return &ListTaskLists{}, responseErr
	}

	if response.Status != "200 OK" {
		saveStatusErr := SaveStatus(response)
		if saveStatusErr != nil {
			return nil, saveStatusErr
		}

		return nil, errors.New("Wrong response while getting the ListTaskLists. Check the logs")
	}

	defer response.Body.Close()

	var listOfLists ListTaskLists
	decodingErr := json.NewDecoder(response.Body).Decode(&listOfLists)
	if decodingErr != nil {
		return &ListTaskLists{}, decodingErr
	}

	return &listOfLists, nil
}

func GetTaskList(webClient *http.Client, name string) (*TaskList, error) {
	list, listTaskListsErr := GetListTaskLists(webClient)
	if listTaskListsErr != nil {
		return &TaskList{}, listTaskListsErr
	}

	var currentNumber int
	currentNumber = -1

	for i := 0; i < len(list.TaskLists); i++ {
		if list.TaskLists[i].DisplayName == name {
			currentNumber = i
			break
		}
	}

	if(currentNumber == -1) {
		fmt.Println("No such list of tasks")
		return nil, errors.New("The list doesn't exist")
	} else {
		return &list.TaskLists[currentNumber], nil
	}
}

func GetListTasks(webClient *http.Client, listId string) (*ListTasks, error) {
	req, requestErr := http.NewRequest(
				http.MethodGet,
				"https://graph.microsoft.com/beta/me/todo/lists/" + listId + "/tasks",
				nil)

	if requestErr != nil {
		return nil, requestErr
	}

	response, responseErr := webClient.Do(req)
	if responseErr != nil {
		return nil, responseErr
	}

	if response.Status != "200 OK" {
		saveStatusErr := SaveStatus(response)
		if saveStatusErr != nil {
			return nil, saveStatusErr
		}

		return nil, errors.New("Wrong response while getting ListTasks. Check the logs")
	}

	defer response.Body.Close()

	var list ListTasks
	decodingErr := json.NewDecoder(response.Body).Decode(&list)
	if decodingErr != nil {
		return nil, decodingErr
	}

	return &list, nil
}

func GetTask(webClient *http.Client, listId, taskId string) (*Task, error) {
	req, requestErr := http.NewRequest(
		"GET",
		"https://graph.microsoft.com/beta/me/todo/lists/" + listId + "/tasks/" + taskId,
		nil)

	if requestErr != nil {
		return nil, requestErr
	}

	response, responseErr := webClient.Do(req)
	if responseErr != nil {
		return nil, responseErr
	}

	if response.Status != "200 OK" {
		saveStatusErr := SaveStatus(response)
		if saveStatusErr != nil {
			return nil, saveStatusErr
		}

		return nil, errors.New("Wrong response while getting task. Check the logs")
	}

	defer response.Body.Close()

	var task Task
	decodingErr := json.NewDecoder(response.Body).Decode(&task)
	if decodingErr != nil {
		return nil, decodingErr
	}

	return &task, nil
}

func DeleteTaskList(webClient *http.Client, listId string) error {

	requestUrl := "https://graph.microsoft.com/beta/me/todo/lists/" + listId
	delReq, reqErr 	:= http.NewRequest("DELETE", requestUrl, nil)
	if reqErr != nil {
		return reqErr
	}

	response, delErr := webClient.Do(delReq)

	if delErr != nil {
		return errors.New("Can't delete the list of tasks")
	}
			
	if response.Status != "204 No Content" {
		saveStatusErr := SaveStatus(response)
		if saveStatusErr != nil {
			return saveStatusErr
		}

		return errors.New("Wrong response while deleting a TaskList. Check the logs")
	}

	return nil
}

func CreateTaskList(webClient *http.Client, name string) error {
	data := map[string]string{"displayName":name}

	requestJson, mrshErr := json.Marshal(data)
	if mrshErr != nil {
		return mrshErr
	}

	bodyReader := bytes.NewReader(requestJson)

    req, postErr := http.NewRequest(http.MethodPost,
                                 "https://graph.microsoft.com/beta/me/todo/lists",
                                  bodyReader)
    if postErr != nil {
    	return postErr
    }

    req.Header.Add("Content-Type", "application/json")
    response, reqErr := webClient.Do(req)

    defer response.Body.Close()

    if reqErr != nil {
    	return reqErr
    }

    if response.Status != "201 Created" {
		saveStatusErr := SaveStatus(response)
		if saveStatusErr != nil {
			return saveStatusErr
		}

		return errors.New("Wrong response while creating a new TaskList. Check logs")
    }

    return nil
}

func CreateTask(webClient *http.Client, listId, taskName string) error {
	taskData := map[string]string{"title":taskName}

	requestJson, mrshErr := json.Marshal(taskData)
	if mrshErr != nil {
		return mrshErr
	}

	bodyReader := bytes.NewReader(requestJson)

    req, postErr := http.NewRequest(
    		http.MethodPost,
            "https://graph.microsoft.com/beta/me/todo/lists/" + listId + "/tasks",
            bodyReader)

    if postErr != nil {
    	return postErr
    }

    req.Header.Add("Content-Type", "application/json")
    response, reqErr := webClient.Do(req)

    defer response.Body.Close()

    if reqErr != nil {
    	return reqErr
    }

    if response.Status != "201 Created" {
		saveStatusErr := SaveStatus(response)
		if saveStatusErr != nil {
			return saveStatusErr
		}
    }

    return nil
}
