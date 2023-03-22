package cmd

import (
	"fmt"
	"github.com/neumanf/mally-cli/services"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"log"
	"syscall"
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
	var email string

	fmt.Println("[DISCLAIMER] Your login credentials are not stored.")
	fmt.Print("Email: ")
	_, err := fmt.Scan(&email)

	if err != nil {
		log.Fatal("Could not read email input")
	}

	fmt.Print("Password: ")
	password, err := term.ReadPassword(syscall.Stdin)

	if err != nil {
		log.Fatal("Could not read password input")
	}

	return email, string(password)
}

type loginResponse struct {
	AccessToken string `json:"accessToken"`
	Email       string `json:"email"`
	Sub         uint   `json:"sub"`
}

func login(email string, password string) string {
	data := map[string]string{"username": email, "password": password}

	res := services.PostRequest[map[string]string, loginResponse]("/auth/login", data)

	return res.AccessToken
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
