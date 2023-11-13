package leetcode

type UserService interface {
	GetUsername() (string, error)
}

type ProblemService interface {
	GetProblemSet() ([]map[string]interface{}, error)
	// GetDailyProblem 如果需要获取题目id、题目名称，则在返回值的map中获取key：question.frontendQuestionId 和 question.titleCn
	GetDailyProblem() (map[string]interface{}, error)
	// GetQuestionByTitleSlug titleSlug: 类似 two-sum 显示在url中的字符串
	GetQuestionByTitleSlug(titleSlug string) (map[string]interface{}, error)
}

type IService interface {
	UserService
	ProblemService
}

type GraphQL struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}
