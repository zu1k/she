package plaintext

import (
	"bufio"
	"bytes"
	"os"
	"sync"

	"github.com/zu1k/she/common"
	"github.com/zu1k/she/log"
	"github.com/zu1k/she/source"
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
func (p *plaintext) Search(key interface{}, resChan chan common.Result, wg *sync.WaitGroup) {
	str := key.(string)
	log.Infoln("Search plain text, key = %s", str)
	//开始搜索
	cmp := []byte(str)
	f, err := os.Open(p.filePath)
	if err != nil {
		wg.Done()
		return
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		info := input.Bytes()
		if bytes.Contains(info, cmp) {
			result := common.Result{
				Score: 1,
				Hit:   str,
				Text:  string(info),
			}
			resChan <- result
		}
	}
	wg.Done()
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
