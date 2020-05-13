package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseAreaList(t *testing.T) {
	contents, err := ioutil.ReadFile("shanghailianjia.html")
	if err != nil {
		panic(err)
	}

	//TODO: finish test
	result := ParseAreaList(contents)

	const resultSize = 17
	expectedUrls := []string{
		"https://sh.lianjia.com/ershoufang/pudong/",
		"https://sh.lianjia.com/ershoufang/minhang/",
		"https://sh.lianjia.com/ershoufang/baoshan/",
	}
	expectedAreas := []string{
		"Area 浦东", "Area 闵行", "Area 宝山",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
	for i, area := range expectedAreas {
		if result.Items[i] != area {
			t.Errorf("expected area #%d: %s, but was %s", i, area, result.Items[i])
		}
	}
}
