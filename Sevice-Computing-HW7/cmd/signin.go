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

// signinCmd represents the signin command
var signinCmd = &cobra.Command{
	Use:   "signin",
	Short: "Use signin to sign in.",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		err := false
		if username == "" || password == "" {
			fmt.Println("Error:")
			err = true
		}
		if username=="" {
			fmt.Println("	Please enter your username!")
		}
		if password=="" {
			fmt.Println("	Please enter your password!")
		}
		if err{
			fmt.Println("Use -h/--help for details.")
			return
		}
		if success := entity.UserSignIn(username, password); success{
			fmt.Println("SignIn success! Welcome", username)
		}

	},
}

func init() {
	rootCmd.AddCommand(signinCmd)

	// Here you will define your flags and configuration settings.
	signinCmd.Flags().StringP("username", "u", "", "your username")
	signinCmd.Flags().StringP("password", "p", "", "your password")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
