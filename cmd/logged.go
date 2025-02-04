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

// loggedCmd represents the logged command
var loggedCmd = &cobra.Command{
	Use:   "logged",
	Short: "List add logged in users",
	Long:  `This command shows all logged in users`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/logged"
		user := User{Username: username, Password: password}

		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodGet, SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("Logged - Error in req: ", err)
			return
		}
		req.Header.Set("Contetn-Type", "application/json")

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
	rootCmd.AddCommand(loggedCmd)

}
