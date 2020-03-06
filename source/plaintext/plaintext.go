package plaintext

import (
	"bufio"
	"bytes"
	"os"

	"github.com/zu1k/she/source"

	"github.com/zu1k/she/log"
)

type plaintext struct {
	filePath string
}

func init() {
	source.Register("plaintext", newPlain)
}

// GetName return plaintext name
func (p *plaintext) GetName() string {
	return "PlainText"
}

// Search return result slice from source plaintext
func (p *plaintext) Search(key interface{}) (result []source.Result) {
	str := key.(string)
	log.Infoln("Search plain text, key = %s", str)
	return searchFileContainsStr(p.filePath, str)
}

func newPlain(info interface{}) source.Source {
	path := info.(string)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	return &plaintext{filePath: path}
}

func searchFileContainsStr(path, str string) (results []source.Result) {
	results = make([]source.Result, 0)
	cmp := []byte(str)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		info := input.Bytes()
		if bytes.Contains(info, cmp) {
			results = append(results, source.Result{
				Score: 1,
				Hit:   str,
				Text:  string(info),
			})
		}
	}
	return
}
