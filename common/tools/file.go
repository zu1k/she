package tools

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func OpenCSC(filepath string) (csvReader *csv.Reader, err error) {
	cntb, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	csvReader = csv.NewReader(strings.NewReader(string(cntb)))
	return
}

func LineCounter(filepath string) (int, error) {
	r, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	var readSize int
	var count int

	buf := make([]byte, 1024)

	for {
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || readSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
	}
	if readSize > 0 && count == 0 || count > 0 {
		count++
	}
	if err == io.EOF {
		return count, nil
	}

	return count, err
}

func Path2Name(filePath string) string {
	fileName := path.Base(filePath)
	if runtime.GOOS == "windows" {
		filepaths := strings.Split(filePath, "\\")
		fileName = filepaths[len(filepaths)-1]
	}
	return fileName
}

func Path2Path(filePath string) string {
	dirpath := ""

	if runtime.GOOS == "windows" {
		filepaths := strings.Split(filePath, ":\\")
		dirpath = filepaths[1]
	} else {
		dirpath := path.Dir(filePath)
		if dirpath == "." {
			return ""
		}
	}
	return dirpath
}
