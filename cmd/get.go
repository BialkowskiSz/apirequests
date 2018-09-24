// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"

	sm "github.com/BialkowskiSz/go-sportmonks"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use: "get -t <api_token> -e <endpoint>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("Unspecified argument(s)")
		}
		return nil
	},
	Short: "Make a GET request to Sportmonks",
	Run: func(cmd *cobra.Command, args []string) {
		sm.SetAPIToken(token)
		d, err := sm.Get(endpoint, includes, page, allPages)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(d))
	},
}

var token string
var endpoint string
var includes string
var page int
var allPages bool

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&token, "token", "t", "", "API access token for request")
	getCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "API Endpoint")
	getCmd.Flags().StringVarP(&includes, "includes", "i", sm.NoIncludes, "API request includes")
	getCmd.Flags().IntVarP(&page, "page", "p", sm.FirstOrAllPages, "Specify results page to get")
	getCmd.Flags().BoolVarP(&allPages, "all-pages", "a", sm.AllPages, "Get all request pages")

	getCmd.MarkFlagRequired("api-token")
	getCmd.MarkFlagRequired("endpoint")
}
