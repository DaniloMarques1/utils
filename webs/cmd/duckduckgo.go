/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
    "fmt"
    "os"
    "os/exec"
    "strings"

)

// duckduckgoCmd represents the duckduckgo command
var duckduckgoCmd = &cobra.Command{
	Use:   "duckduckgo",
    Aliases: []string{"duc"},
	Short: "It will make a search of the argument passed to duckduckgo",
	Long: `
      It will grab everything passed as argument and search using the
      duckduckgo search engine
    `,
	Run: func(cmd *cobra.Command, args []string) {
        SearchOnTheDuckDuckGo(args)
	},
}

func SearchOnTheDuckDuckGo(args []string) {
    const url string = "https://duckduckgo.com/?q="
    if len(args) < 1 {
        fmt.Println("You need to give something to search for")
        os.Exit(1)
    }
    fmt.Println("Searching the web")
    searchString := strings.Join(args, " ")
    exec.Command("firefox", url + searchString).Start()
}

func init() {
	rootCmd.AddCommand(duckduckgoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// duckduckgoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// duckduckgoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
