/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the database",
	Long: `Search the database for a user, identified by a User ID.
	The command returns the full record of the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/username"
		user := User{Username: username, Password: password}

		var u2 User
		err := json.Unmarshal([]byte(data), &u2)
		if err != nil {
			fmt.Println("Unmarshal:", err)
			return
		}

		buf := new(bytes.Buffer)
		err = user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		URL := SERVER + PORT + endpoint + "/" + fmt.Sprint(u2.ID)
		req, err := http.NewRequest(http.MethodGet, URL, buf)
		if err != nil {
			fmt.Println("Search - Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("Do:", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp)
			return
		}

		var Returned = User{}
		SliceFromJSON(&Returned, resp.Body)
		data, err := PrettyJSON(Returned)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(data)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
