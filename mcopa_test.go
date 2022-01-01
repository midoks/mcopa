package mcopa

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func ReadFile(file string) (string, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0600)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return string(b), err
}

func TestParse(t *testing.T) {
	appDir, _ := os.Getwd()
	testData, _ := ReadFile(fmt.Sprintf("%s/testdata/attachment.eml", appDir))
	// fmt.Println(appDir, testData)

	bufferedBody := bufio.NewReader(strings.NewReader(testData))

	email, err := Parse(bufferedBody)
	fmt.Println(email, err)
}
