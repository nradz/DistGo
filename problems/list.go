package problems

var problemMap = make(map[string]*Problem)

//Getproblem returns a problem pointer by his name.
func GetProblem(prob string) *Problem{

	saved, ok := problemMap[prob]

	if !ok{
		return nil
	}

	return saved
	
}

//AddProblem insert a problem in the list of available problems.
func AddProblem(name string, prob Problem){
	problemMap[name] = &prob
}