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

// dictionaryCmd represents the dictionary command
var dictionaryCmd = &cobra.Command{
	Use:   "dictionary",
    Aliases: []string{"dic"},
	Short: "Will search a word in the dicionary",
	Long: `
        will get the word passes as argument and search in the
        website https://www.dicio.com.br/. PS: has to be a
        portuguese word.
    `,
	Run: func(cmd *cobra.Command, args []string) {
        const url string = "https://www.dicio.com.br/"
        service.SearchTheWeb("dictionary", url, args)
	},
}

/*
func SearchonTheDicionary(args []string) {
    const url = "https://www.dicio.com.br/"
    if len(args) < 1 {
        fmt.Println("You need to give something to search for")
        os.Exit(1)
    }
    exec.Command("firefox", url + strings.ToLower(args[0])).Start()
    fmt.Println("searching web....")
}
*/

func init() {
	rootCmd.AddCommand(dictionaryCmd)
}
