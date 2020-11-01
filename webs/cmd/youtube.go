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
    "webs/service"
)

// youtubeCmd represents the youtube command
var youtubeCmd = &cobra.Command{
	Use:   "youtube",
    Aliases: []string {"yt"},
	Short: "Will make a youtube search",
	Long: `
        I will search on youtube the passed string
    `,
	Run: func(cmd *cobra.Command, args []string) {
        const url string = "https://youtube.com/results?search_query="
        service.SearchTheWeb("youtube", url, args)
	},
}

func init() {
	rootCmd.AddCommand(youtubeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// youtubeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// youtubeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
