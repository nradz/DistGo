package problems


func GetProblem(prob string) Problem{

	switch prob{
	case "pruebaProblem":
		return &pruebaProblem{}
	//case "TSPGeneticProblem":
	//	return &simpleGeneticProblem{}
	default:
		println("The problem doesn't exists.")
	}

	return nil
	
}