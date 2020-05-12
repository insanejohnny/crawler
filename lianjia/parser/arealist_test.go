package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseAreaList(t *testing.T) {
	contents, err := ioutil.ReadFile("shanghailianjia.html")
	if err != nil {
		panic(err)
	}

	//TODO: finish test
	// result := ParseAreaList(contents)
	fmt.Println(string(contents))
}
