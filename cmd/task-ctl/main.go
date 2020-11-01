package main

import (
	"fmt"
	req "github.com/Microsoft-To-Do-API/src"
)

func main() {
	webClient := req.GetDefaultClient()
	err := req.DeleteTaskList(webClient, "AAMkADg3NGRjZTk2LTY2MmItNGY1MS1hNTczLWEwMTU2ZDlhZjIxNwAuAAAAAAAylY3yThb0TZ016ZEL6RCkAQAKFTyMOPoWSLF7YUzGp3WzAAAYeRFRAAA=")
	fmt.Println("err = ", err)
	//fmt.Println(list)
}

