package request

import (
	oauth "golang.org/x/oauth2"
)

type TokenResponse struct {
	TokenValue 	*oauth.Token
	TokenErr 	error
}

type TaskList struct {
	DataType		string 	`json:"@odata.etag"`
	Id  			string 	`json:"id,omitempty"`
	DisplayName 	string 	`json:"displayName"`
	IsOwner  		bool   	`json:"isOwner"`
	IsShared  		bool 	`json:"isShared"`
	WellknownListName string `json:"wellknownListName"`
}

type ListTaskLists struct {
	DataContext string 		`json:"@odata.context"`
	TaskLists 	[]TaskList 	`json:"value"`
}

type ListTasks struct {
	Context 	string `json:"@odata.context"`
	ListOfTasks []Task `json:"value"`
}

type TaskBody struct {
	Content 	string `json:"content"`
	ContentType string `json:"contentType"`
}

type DateStruct struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type Task struct {
    Context 		string   	`json:"@odata.context"`
    Tag 			string   	`json:"@odata.etag"`
    Importance 		string   	`json:"importance"`
    IsReminder 		bool 	 	`json:"isReminderOn"`
    Status 			string 	 	`json:"status"`
    Title 			string 	 	`json:"title"`
    CreationTime 	string 	 	`json:"createdDateTime"`
    LastModified 	string 	 	`json:"lastModifiedDateTime"`
    Id 				string 	 	`json:"id"`
    Body  			TaskBody 	`json:"body"`
    DueDateTime 	DateStruct 	`json:"dueDateTime"`
}

func (response TokenResponse) Token() (*oauth.Token, error) {
	return response.TokenValue, response.TokenErr
}
