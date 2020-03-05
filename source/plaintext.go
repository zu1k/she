package source

import (
	"bufio"
	"bytes"
	"os"

	"github.com/zu1k/she/log"
)

type plaintext struct {
	filePath string
}

func init() {
	register("plaintext", newPlain)
}

// GetName return plaintext name
func (p *plaintext) GetName() string {
	return "PlainText"
}

// Search return result slice from source plaintext
func (p *plaintext) Search(key interface{}) (result []Result) {
	str := key.(string)
	log.Infoln("Search 12306, key = %s", str)
	return searchFileContainsStr(p.filePath, str)
}

func newPlain(info interface{}) Source {
	path := info.(string)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	return &plaintext{filePath: path}
}

func searchFileContainsStr(path, str string) (results []Result) {
	results = make([]Result, 0)
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
			results = append(results, Result{
				Score: 1,
				Hit:   str,
				Text:  string(info),
			})
		}
	}
	return
}
