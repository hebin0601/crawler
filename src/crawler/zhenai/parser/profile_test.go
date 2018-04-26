package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_date.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents)
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}

	profile := result.Items[0].(model.Porfile)
	expected := model.Porfile{
		Name:     "开欣",
		Gender:   "女",
		Age:      24,
		Weight:   0,
		Height:   164,
		Marriage: "未婚",
	}

	if profile != expected {
		t.Errorf("expected %v;but was %v", expected, profile)
	}
}
