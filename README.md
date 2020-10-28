# Simple Library for Microsoft To-Do API

This little liblary contains tools for management Microsoft To Do tasks. It uses Microsoft REST API and Microsoft graph for these aims. Look through [this link](https://docs.microsoft.com/en-us/graph/api/resources/todo-overview?view=graph-rest-beta) to get more information. 

### The library provides next methods :

- [GetListTaskLists](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L16) - To get **ListTaskLists** structure. It contains brief information about all lists of tasks.
- [GetTaskListShort](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L40) - To get **todoTaskList** structure. It contains basic data about current list of tasks like Id etc.
- [GetListTasks](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L61) - To get **List tasks** structure. It contains all information about tasks containing int the list.
- [GetTask](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L97) - To get **Task** structure. It contains data linked with a task.
- [DeleteTaskList](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L152) - To delete current **todoTaskList**.
- [CreateTaskList](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L180) - To create a new **todoTaskList**.
- [CreateTask](https://github.com/bigtelescope/Microsoft-To-Do-API/blob/e2a728f93d1e06092617d1c4f36ea7233072ac88/requests.go#L204) - To create a new **Task**.