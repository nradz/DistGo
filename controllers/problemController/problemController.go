package problemController


import(
	"problems"
)


type problemController interface{
	Init()
	NewRequest(uint32, uint32) (string, data, uint32, error)
	NewResult(uint32, []string) error
	Close()
}

func New(prob *problems.Problem) *problemController{
	switch prob.Type(){
		case "simple":
			return simple.New(prob)
		default:
			return nil
	}
}