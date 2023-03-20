package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/neumanf/mally-cli/services"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into the Mally website",
	Long:  `If you do not have an account, create one here: https://mally.neumanf.com/signup`,
	Run: func(cmd *cobra.Command, args []string) {
		email, password := getUserCredentials()

		token := login(email, password)

		services.StoreToken(token)

		fmt.Println("Login was successfull. You can now use all features.")
	},
}

func getUserCredentials() (string, string) {
	var email, password string

	fmt.Println("[DISCLAIMER] Your login credentials are not stored.")
	fmt.Print("Email: ")
	_, err := fmt.Scan(&email)

	if err != nil {
		log.Fatal("Could not read email input")
	}

	fmt.Print("Password: ")
	_, err = fmt.Scan(&password)

	if err != nil {
		log.Fatal("Could not read password input")
	}

	return email, password
}

func login(email string, password string) string {
	values := map[string]string{"username": email, "password": password}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://api.mally.neumanf.com/api/auth/login", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 401 {
		log.Fatal("Invalid credentials")
	}

	var res map[string]string

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		log.Fatal("Could not decode API response")
	}

	return res["accessToken"]
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
