/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"Agenda/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Use register to register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")
		err := false
		if username == "" || password == "" || email == "" || telephone == "" {
			fmt.Println("Error:")
			err = true
		}
		if username=="" {
			fmt.Println("	Please enter your username!")
		}
		if password=="" {
			fmt.Println("	Please enter your password!")
		}
		if email=="" {
			fmt.Println("	Please enter your email!")
		}
		if telephone=="" {
			fmt.Println("	Please enter your telephone!")
		}
		if err{
			fmt.Println("Use -h/--help for details.")
			return
		}
		if success := entity.UserRegister(username, password, email, telephone); success{
			fmt.Println("Register success! Please remember your username and password!")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.
	registerCmd.Flags().StringP("username", "u", "", "your username")
	registerCmd.Flags().StringP("password", "p", "", "your password")
	registerCmd.Flags().StringP("email", "e", "", "your email address")
	registerCmd.Flags().StringP("telephone", "t", "", "your telephone number")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
