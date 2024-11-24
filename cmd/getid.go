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

// getidCmd represents the getid command
var getidCmd = &cobra.Command{
	Use:   "getid",
	Short: "Return User ID, given a username",
	Long:  `This command return the User ID of a user, given their username`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/getid"
		user := User{Username: username, Password: password}

		var u2 User
		err := json.Unmarshal([]byte(data), &u2)
		if err != nil {
			fmt.Println("Unmarshal:", err)
			return
		}

		if u2.Username == "" {
			fmt.Println("Emtpy username!")
			return
		}

		buf := new(bytes.Buffer)
		err = user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		URL := SERVER + PORT + endpoint + "/" + u2.Username
		req, err := http.NewRequest(http.MethodGet, URL, buf)
		if err != nil {
			fmt.Println("GetId - Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "aplication/json")

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

		fmt.Println("User", Returned.Username, "has ID:", Returned.ID)
	},
}

func init() {
	rootCmd.AddCommand(getidCmd)
}
