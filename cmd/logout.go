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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout for user",
	Long:  `Logging out a user from the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/logout"
		user := User{Username: username, Password: password}

		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, SERVER+PORT+endpoint, buf)
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
		} else {
			fmt.Println("User", user.Username, "logged out!")
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
