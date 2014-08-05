//problems package contains the implemented problems and tools to do it.
//
//NOTE: All the problems have to implement the function 'init'. This function must 
//call AddProblem.
package problems

//The information that the problem sends to the clients.
type Data interface{}

//It is the struct where the problem sends updates to the problemController.
type ProblemUpdate struct{
	Alg string //The algorithm that the clients have to execute.
	Data Data //Data to the clients.
	Number uint32 //Update Number. It will increase with new updates.
}

//The interface that the problems have to implement it.
type Problem interface{
	Start(chan ProblemUpdate) ProblemUpdate //The first call to the problem.
	//This have to configure it.
	NewResult([]string, uint32) //The function will be called when a new result is received.
	Loop() //If the problem needs to execute locally a algorithm, use this function.
	// In other case, write a simple 'return'.
	Type() string //It returns the type of the problem.
}


