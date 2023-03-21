package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/neumanf/mally-cli/services"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var shortenCmd = &cobra.Command{
	Use:   "shorten <url>",
	Short: "Shortens a URL",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || args[0] == "" {
			return errors.New("URL is required")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		slug := createShortUrl(args[0])

		fmt.Println("URL shortened successfully.\nShort URL: ", "https://mally.neumanf.com/s/"+slug)
	},
}

func createShortUrl(url string) string {
	values := map[string]string{"url": url}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	token := services.GetToken()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.mally.neumanf.com/api/url-shortener", bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "accessToken="+token)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]string

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		log.Fatal("Could not decode API response", err)
	}

	if resp.StatusCode != 201 {
		if resp.StatusCode == 401 {
			log.Fatal("You are not authorized, please use 'mally-cli login' first to login and then try again.")
		}

		log.Fatal("Status code: ", resp.StatusCode, ", message: ", res["message"])
	}

	return res["slug"]
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
