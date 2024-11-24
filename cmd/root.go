/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	SERVER   string
	PORT     string
	data     string
	username string
	password string
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

const (
	empty = ""
	tab   = "\t"
)

func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func SliceFromJSON(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func PrettyJSON(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}

	return buffer.String(), nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-client",
	Short: "A REST API client",
	Long:  `A Client for a RESTful server.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "username", "The username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "admin", "The password")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "{}", "JSON Record")
	rootCmd.PersistentFlags().StringVarP(&SERVER, "server", "s", "http://localhost", "RESTful server hostname")
	rootCmd.PersistentFlags().StringVarP(&PORT, "port", "P", ":1234", "Port of RESTful Server")
}
