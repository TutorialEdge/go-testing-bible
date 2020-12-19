package markdown

import (
	"fmt"
	"os"
)

// ConvertMarkdownToHTML - takes in a markdown file and converts it to HTML
func ConvertMarkdownToHTML(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	fmt.Println(f)

	return "", nil
}
