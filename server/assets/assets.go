package assets

import (
	"io/ioutil"
	"os"
)

func Asset(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
