package problems


func GetProblem(prob string) Problem{

	switch prob{
	case "pruebaProblem":
		return pruebaProblem{}
	default:
		println("The problem doesn't exists.")
	}

	return nil
	
}