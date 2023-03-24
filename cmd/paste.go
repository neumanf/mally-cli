package cmd

import (
	"errors"
	"fmt"
	"github.com/neumanf/mally-cli/services"
	"log"
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

type pasteResponse struct {
	Slug string `json:"slug"`
}

func createPasteFromSnippet(syntax string, snippet string) string {
	data := map[string]string{"syntax": syntax, "content": snippet}

	token, err := services.GetToken()

	if err != nil {
		log.Fatal(err)
	}

	res := services.PostRequest[map[string]string, pasteResponse]("/pastebin", data, &token)

	return res.Slug
}

func init() {
	rootCmd.AddCommand(pasteCmd)

	pasteCmd.Flags().StringVarP(&File, "file", "f", "", "A path to a file to paste")
}
