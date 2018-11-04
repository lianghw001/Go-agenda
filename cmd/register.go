// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"regexp"

	"github.com/cyulei/agenda/datarw"
	"github.com/cyulei/agenda/entity"

	"github.com/spf13/cobra"
)

//var cfgFile string
var registerName, registerPassword string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new User",
	Long: `register:Users are registered through username passwords and phone and email.
	For example:
	register a new user,with name:User1,password:12345678
	agenda register -n=User1 -p=12345678 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		register(registerName, registerPassword)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&registerName, "name", "n", "", "user's name")
	registerCmd.Flags().StringVarP(&registerPassword, "password", "p", "", "user's password")

}

func register(name string, password string) {
	logInit()
	defer logFile.Close()
	logSave("cmd: register called", "[Info]")

	if isValidName(name) && isValidPassword(password) {
		users := datarw.GetUsers()
		if entity.HasUser(name, users) {
			logSave("The username has been registered", "[Warning]")
			logSave("Register fail", "[Warning]")
			return
		}
		var email, phone string
		fmt.Println("please input your email:(xxx@xx.xx)")
		fmt.Scanln(&email)
		for !isValidEmail(email) {
			fmt.Println("please input your email, and input the correct format:(xxx@xx.xx)")
			fmt.Scanln(&email)
		}
		fmt.Println("please input your phone:(for example 11011912010)")
		fmt.Scanln(&phone)
		for !isValidPhone(phone) {
			fmt.Println("please input your phone,and input the correct format:(for example 11011912010)")
			fmt.Scanln(&phone)
		}
		newuser := entity.User{Name: name, Password: password, Email: email, Phone: phone}
		users = append(users, newuser)
		datarw.SaveUsers(users)
		logSave("Registration complete", "[Info]")
		return
	}
	logSave("Register fail", "[Warning]")

}

func isValidName(n string) bool {
	b := []byte(n)
	val, _ := regexp.Match(".+", b)
	if !val {
		logSave("flag -n ,name is invaild", "[Warning]")
	}
	return val
}
func isValidPassword(p string) bool {
	b := []byte(p)
	val, _ := regexp.Match(".+", b)
	if len(p) < 8 {
		fmt.Println("The password must be longer than 8 digits")
		val = false
	}
	if !val {
		logSave("flag -p ,password is invaild", "[Warning]")
	}
	return val
}
func isValidEmail(e string) bool {
	b := []byte(e)
	val, _ := regexp.Match("\\w*@\\w*\\.w*", b)

	if !val {
		logSave("email is invaild", "[Warning]")
	}
	return val
}
func isValidPhone(p string) bool {
	b := []byte(p)

	val, _ := regexp.Match("\\d{11}", b)

	if !val {
		logSave("phone is invaild", "[Warning]")
	}
	return val
}
