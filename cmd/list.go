/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available users ",
	Long:  `The list command all available users.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/getall"
		user := User{Username: username, Password: password}
		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodGet, SERVER+PORT+endpoint, buf)
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
			fmt.Println(resp)
			return
		}

		var users = []User{}
		SliceFromJSON(&users, resp.Body)
		data, err := PrettyJSON(users)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(data)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
