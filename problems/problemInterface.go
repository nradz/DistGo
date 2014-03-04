package problems

import(
	"github.com/nradz/DistGo/channels"
	)

type Problem interface{
	Init() channels.ProblemUpdate
	NewResult([]string)
	Loop()
}