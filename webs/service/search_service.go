package service

import (
    "os"
    "os/exec"
    "fmt"
    "strings"
)

func SearchTheWeb(command string, url string, args []string) {
    if len(args) < 1 {
        fmt.Println("You need to pass a argument to complete the search")
        os.Exit(1)
    }
    searchString := getSearchString(command, args)
    exec.Command("firefox", "--new-tab", url + searchString).Start()
}

func getSearchString(command string, args []string) string {
    var searchString string
    switch command {
        case "dictionary":
            searchString = strings.Join(args, " ")
            searchString = strings.Replace(searchString, " ", "-", strings.Count(searchString, " "))
        default: 
            searchString = strings.Join(args, " ")
    }

    return searchString
}
