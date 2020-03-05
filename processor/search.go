package processor

import "github.com/zu1k/she/source"

func Search(key string) (resultText string) {
	results := source.Search(key)
	for _, result := range results {
		resultText += result.Text + "\r\n"
	}
	return
}
