package request

import (
	"fmt"
	"bytes"
	"errors"
	"context"
	"net/http"
	"encoding/json"
	oauth "golang.org/x/oauth2"
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
		///
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

func GetListTasks(webClient *http.Client, name string) (*ListTasks, error) {
	taskListShort, shortErr := GetTaskList(webClient, name)
	if(shortErr != nil) {
		return nil, shortErr
	}

	id := taskListShort.Id
	
	req, requestErr := http.NewRequest(
				http.MethodGet,
				"https://graph.microsoft.com/beta/me/todo/lists/" + id + "/tasks",
				nil)

	if requestErr != nil {
		return nil, requestErr
	}

	response, responseErr := webClient.Do(req)
	if responseErr != nil {
		return nil, responseErr
	} else if response.Status != "200 OK" {
		/////
	}

	defer response.Body.Close()

	var list ListTasks
	decodingErr := json.NewDecoder(response.Body).Decode(&list)
	if decodingErr != nil {
		return nil, decodingErr
	}

	return &list, nil
}

func GetTask(webClient *http.Client, listName, taskName string) (*Task, error) {
	listTaskShort, shortErr	:= GetTaskList(webClient, listName)
	if shortErr != nil {
		return nil, errors.New("No the list with the name")
	}

	listTask, listErr := GetListTasks(webClient, listName)
	if listErr != nil {
		return nil, errors.New("No the task with the name")
	}

	var taskId, listId string
	taskId = "default"

	for i := 0; i < len(listTask.ListOfTasks); i++ {
		if listTask.ListOfTasks[i].Title == taskName {
			taskId = listTask.ListOfTasks[i].Id
			listId = listTaskShort.Id
			break
		}
	}

	if(taskId == "default") {
		return nil, errors.New("The list doesn't exist")
	}

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
		/////
	}

	defer response.Body.Close()

	var task Task
	decodingErr := json.NewDecoder(response.Body).Decode(&task)
	if decodingErr != nil {
		return nil, decodingErr
	}

	return &task, nil
}

func DeleteTaskList(webClient *http.Client, name string) error {
	list, deleteErr := GetListTaskLists(webClient)
	if deleteErr != nil {
		return deleteErr
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
		return errors.New("No such list of tasks")
	} else {
		requestUrl := "https://graph.microsoft.com/beta/me/todo/lists/" + list.TaskLists[currentNumber].Id
		delReq, reqErr 	:= http.NewRequest("DELETE", requestUrl, nil)
		if deleteErr != nil {
			return reqErr
		}

		_, delErr := webClient.Do(delReq)

		if delErr != nil {
			return errors.New("Can't delete the list of tasks")
		}
		//check the status here !!!
		//here is a response above
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
    resp, reqErr := webClient.Do(req)

    defer resp.Body.Close()

    if reqErr != nil {
    	return reqErr
    }

    if resp.Status != "201 Created" {
    	//////
    }

    return nil
}

func CreateTask(webClient *http.Client, listName, taskName string) error {
	listTaskShort, shortErr	:= GetTaskList(webClient, listName)
	if shortErr != nil {
		return shortErr
	}

	listId 	 := listTaskShort.Id
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
    	/////
    }

    return nil
}

func GetDefaultClient() *http.Client {
	token, tokenErr := GetToken()
	if tokenErr != nil {
		return nil
	}

	tempToken := TokenResponse {
		TokenValue	: token,
		TokenErr 	: nil,
	}

	clientContext := context.Background()
	webClient  	  := oauth.NewClient(clientContext, tempToken)

	return webClient
}

func GetToken() (*oauth.Token, error) {
	authEndpoint := oauth.Endpoint {
		AuthURL : "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		TokenURL : "https://login.microsoftonline.com/common/oauth2/v2.0/token",
	}

	authConfig := oauth.Config {
		ClientID 	: "1f0be847-d5ad-4872-aa14-7e584d7cc940",
		ClientSecret 	: "mSq-Ser-o66Bt-ASLQ3RkQ2i~~OTB.n02a",
		Endpoint 	: authEndpoint,
		RedirectURL	: "https://login.microsoftonline.com/common/oauth2/nativeclient",
		Scopes 		: []string {
					"offline_access",
					"Tasks.ReadWrite",
		},
	}

	resp := authConfig.AuthCodeURL("state", oauth.AccessTypeOffline)

	fmt.Println("go to the next link : ")
	fmt.Println(resp)
	fmt.Println()
	fmt.Println()

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		fmt.Println("Can`t get code!")
		return nil, err
	}

	ctx := context.Background()
	httpClient := &http.Client{}
	ctx = context.WithValue(ctx, oauth.HTTPClient, httpClient)

	token, err := authConfig.Exchange(ctx, code)

	return token, err
}
