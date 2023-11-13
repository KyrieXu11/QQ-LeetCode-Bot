package leetcode

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

var client IService

func init() {
	bs, err := ioutil.ReadFile("../leetcode_credential.credential")
	if err != nil {
		log.Fatal(err.Error())
	}
	cookies := string(bs)
	client = NewLeetCodeClient(cookies)
}

func TestLeetCodeClient_GetUsername(t *testing.T) {
	username, err := client.GetUsername()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(username)
}

func TestLeetCodeClient_GetDailyProblem(t *testing.T) {
	res, err := client.GetDailyProblem()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%+v", res)
}

func TestLeetCodeClient_GetQuestionByTitleSlug(t *testing.T) {
	slug := "couples-holding-hands"
	ques, err := client.GetQuestionByTitleSlug(slug)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%+v", ques)
}
