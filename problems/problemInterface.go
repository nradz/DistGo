package problems

type data interface{}

type ProblemUpdate struct{
	Alg string
	Data data
}

type Problem interface{
	Init(chan ProblemUpdate) ProblemUpdate
	NewResult([]string)
	Loop()
}


