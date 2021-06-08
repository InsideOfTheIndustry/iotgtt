package filemd5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var a = 1

func ByReadall() string {
	a++
	file, err := os.Open("./asd.txt")
	if err != nil {
		return ""
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(all))
}

func ByPieceRead() string {
	a++
	file, err := os.Open("./asd.txt")
	if err != nil {
		return ""
	}
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}
