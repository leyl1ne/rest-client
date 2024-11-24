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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into the system",
	Long:  `The login command logins an existing user the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/login"
		user := User{Username: username, Password: password}

		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("Login - Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 + time.Second,
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
			fmt.Println("User", user.Username, "logged in!")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

}
