package main

import (
	req "github.com/Microsoft-To-Do-API/src"
)

func main() {
	webClient := req.GetDefaultClient()
	req.CreateTask(webClient, "working list", "one more fucking task")
}

