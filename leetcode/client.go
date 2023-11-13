package leetcode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"qq_bot/errors"
	"qq_bot/utils"
	"strings"
	"time"
)

type leetcodeCredential struct {
	cookies map[string]string
}

func newLeetcodeCredential(cookie string) *leetcodeCredential {
	cookies := utils.ParseCookies(cookie)
	return &leetcodeCredential{
		cookies: cookies,
	}
}

func (o *leetcodeCredential) csrfToken() string {
	return o.cookies["csrftoken"]
}

type LeetCodeClient struct {
	client     *http.Client
	credential *leetcodeCredential
	endpoint   string
	headers    http.Header
}

func NewLeetCodeClient(cookie string) IService {
	var endpoint = ""
	var b = false
	if endpoint, b = os.LookupEnv("LEETCODE_GRAPHQL_ENDPOINT"); !b {
		endpoint = "https://leetcode.cn/graphql/"
	}
	var credential = newLeetcodeCredential(cookie)
	var headers = http.Header{
		"Content-Type": {"application/json; charset=UTF-8"},
		"Cookie":       {cookie},
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"},
	}

	csrfToken := credential.csrfToken()
	if len(csrfToken) > 0 {
		headers.Add("X-Csrf-Token", csrfToken)
	}

	return &LeetCodeClient{
		client: &http.Client{
			Timeout: time.Second * 60,
		},
		credential: credential,
		endpoint:   endpoint,
		headers:    headers,
	}
}

func (o *LeetCodeClient) makeRequest(endpoint, method, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = o.headers

	resp, err := o.client.Do(req)
	return resp, err
}

func (o *LeetCodeClient) makeGraphQLRequest(endpoint, method, graphQL string) (interface{}, error) {
	resp, err := o.makeRequest(endpoint, method, graphQL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res interface{}
	if err = json.Unmarshal(bs, &res); err != nil {
		return nil, err
	}
	return res, nil

}

func (o *LeetCodeClient) GetUsername() (string, error) {
	endpoint := o.endpoint + "noj-go/"
	graphQL := &GraphQL{
		OperationName: "globalData",
		Query:         "\n    query globalData {\n  userStatus {\n    isSignedIn\n    isPremium\n    username\n    realName\n    avatar\n    userSlug\n    isAdmin\n    checkedInToday\n    useTranslation\n    premiumExpiredAt\n    isTranslator\n    isSuperuser\n    isPhoneVerified\n    isVerified\n  }\n  jobsMyCompany {\n    nameSlug\n  }\n}\n    ",
		Variables:     map[string]interface{}{}}
	key := "data.userStatus.username"
	r, err := o.getValue(endpoint, http.MethodPost, key, graphQL)
	if err != nil {
		return "", err
	}

	return r.(string), nil
}

func (o *LeetCodeClient) GetProblemSet() ([]map[string]interface{}, error) {
	graphQl := &GraphQL{
		OperationName: "problemsetQuestionList",
		Query:         "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    hasMore\n    total\n    questions {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId\n      isFavor\n      paidOnly\n      solutionNum\n      status\n      title\n      titleCn\n      titleSlug\n      topicTags {\n        name\n        nameTranslated\n        id\n        slug\n      }\n      extra {\n        hasVideoSolution\n        topCompanyTags {\n          imgUrl\n          slug\n          numSubscribed\n        }\n      }\n    }\n  }\n}\n    ",
		Variables: map[string]interface{}{
			"categorySlug": "",
			"skip":         0,
			"limit":        100,
			"filters":      map[string]interface{}{},
		},
	}
	key := "data.problemsetQuestionList.questions"
	v, err := o.getValue(o.endpoint, http.MethodPost, key, graphQl)
	if err != nil {
		return nil, err
	}

	v1 := v.([]interface{})
	res := []map[string]interface{}{}
	for _, item := range v1 {
		res = append(res, item.(map[string]interface{}))
	}

	return res, nil
}

func (o *LeetCodeClient) GetDailyProblem() (map[string]interface{}, error) {
	graphQl := &GraphQL{
		OperationName: "questionOfToday",
		Query:         "\n    query questionOfToday {\n  todayRecord {\n    date\n    userStatus\n    question {\n      questionId\n      frontendQuestionId: questionFrontendId\n      difficulty\n      title\n      titleCn: translatedTitle\n      titleSlug\n      paidOnly: isPaidOnly\n      freqBar\n      isFavor\n      acRate\n      status\n      solutionNum\n      hasVideoSolution\n      topicTags {\n        name\n        nameTranslated: translatedName\n        id\n      }\n      extra {\n        topCompanyTags {\n          imgUrl\n          slug\n          numSubscribed\n        }\n      }\n    }\n    lastSubmission {\n      id\n    }\n  }\n}\n    ",
		Variables:     map[string]interface{}{},
	}

	key := "data.todayRecord[0]"
	value, err := o.getValue(o.endpoint, http.MethodPost, key, graphQl)
	if err != nil {
		return nil, err
	}
	return value.(map[string]interface{}), nil
}

func (o *LeetCodeClient) GetQuestionByTitleSlug(titleSlug string) (map[string]interface{}, error) {
	graphQl := &GraphQL{
		OperationName: "questionTranslations",
		Query:         "\n    query questionTranslations($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    translatedTitle\n    translatedContent\n  }\n}\n    ",
		Variables: map[string]interface{}{
			"titleSlug": titleSlug,
		},
	}
	key := "data.question"
	value, err := o.getValue(o.endpoint, http.MethodPost, key, graphQl)
	if err != nil {
		return nil, err
	}
	return value.(map[string]interface{}), nil
}

func (o *LeetCodeClient) getValue(endpoint, method, key string, payload interface{}) (interface{}, error) {

	graphql := ""
	// todo: 简单的check一下类型，如果要判断是否为结构体类型，则需要用反射判断
	switch payload.(type) {
	case string:
		graphql = payload.(string)
	default:
		bs, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.NewError(errors.InternalCodeError, err)
		}
		graphql = string(bs)
	}

	res, err := o.makeGraphQLRequest(endpoint, method, graphql)
	if err != nil {
		return "", err
	}
	err = handleErrorMsg(res)
	if err != nil {
		return "", err
	}
	r, err := utils.GetDictValueByKey(res, key)
	if err != nil {
		return "", errors.NewError(errors.InternalCodeError, err)
	}
	if r == nil {
		return "", errors.NewError(errors.InternalCodeError, fmt.Errorf("key error"))
	}
	return r, nil
}

func handleErrorMsg(d interface{}) error {
	r, err := utils.GetDictValueByKey(d, "errors[0].message")
	if err != nil {
		return errors.NewError(errors.InternalCodeError, err)
	}
	if r == nil {
		return nil
	}

	return errors.NewError(errors.BillResponseError, fmt.Errorf(r.(string)))
}
