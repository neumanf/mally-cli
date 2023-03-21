package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/neumanf/mally-cli/services"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var File string

var pasteCmd = &cobra.Command{
	Use:   "paste <syntax> [snippet]",
	Short: "Creates a pastes from a code or text snippet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || args[0] == "" {
			return errors.New("A syntax must be provided")
		}

		if File == "" && len(args) < 2 {
			return errors.New("A text snippet or a file must be provided")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var slug string

		if File != "" {
			slug = createPasteFromFile(args[0])
		} else {
			slug = createPasteFromSnippet(args[0], args[1])
		}

		fmt.Println("Paste created successfully.\nPaste URL: ", "https://mally.neumanf.com/p/"+slug)
	},
}

func createPasteFromFile(syntax string) string {
	text, err := os.ReadFile(File)

	if err != nil {
		log.Fatal(err)
	}

	snippet := string(text)

	return createPasteFromSnippet(syntax, snippet)
}

func createPasteFromSnippet(syntax string, snippet string) string {
	values := map[string]string{"syntax": syntax, "content": snippet}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	token := services.GetToken()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.mally.neumanf.com/api/pastebin", bytes.NewBuffer(jsonData))

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
	rootCmd.AddCommand(pasteCmd)

	pasteCmd.Flags().StringVarP(&File, "file", "f", "", "A path to a file to paste")
}
