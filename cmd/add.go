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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Long:  `Add a new user to the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoit := "/add"
		u1 := User{Username: username, Password: password}

		// converting a data string to a custom structure
		var u2 User
		err := json.Unmarshal([]byte(data), &u2)
		if err != nil {
			fmt.Println("Unmarshal:", err)
			return
		}

		users := []User{}
		users = append(users, u1)
		users = append(users, u2)

		buf := new(bytes.Buffer)
		err = SliceToJSON(users, buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, SERVER+PORT+endpoit, buf)
		if err != nil {
			fmt.Println("GetAll - Error in req: ", err)
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
			fmt.Println("Status code:", resp.Status)
		} else {
			fmt.Println("User", u2.Username, "added.")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
