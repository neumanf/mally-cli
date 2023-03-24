package cmd

import (
	"errors"
	"fmt"
	"github.com/neumanf/mally-cli/services"
	"github.com/spf13/cobra"
	"log"
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

type shortURLResponse struct {
	Slug string `json:"slug"`
}

func createShortUrl(url string) string {
	data := map[string]string{"url": url}

	token, err := services.GetToken()

	if err != nil {
		log.Fatal(err)
	}

	res := services.PostRequest[map[string]string, shortURLResponse]("/url-shortener", data, &token)

	return res.Slug
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
