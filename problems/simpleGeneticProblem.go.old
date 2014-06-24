package problems


import(
	"os"
	"bufio"
	"sync"
	"fmt"
	"sort"
	"strconv"
	"encoding/xml"
	"encoding/json"
	"github.com/nradz/DistGo/channels"
	)

var(
	problemChan = channels.ProblemControlChannel()
	)


//XML structs
type edge struct{
	Cost float64 `xml:"cost,attr"`
	City int `xml:",chardata"`
}

type city struct{
	Edges []edge `xml:"edge"`
}

type TSPxml struct{
	XMLName xml.Name `xml:"travellingSalesmanProblemInstance"`
	Name string `xml:"name"`
	Desc string `xml:"description"`
	Cities []city `xml:"graph>vertex"`
}
//

//problem struct
type simpleGeneticProblem struct{
	sync.Mutex
	graph [][]int
	size int
	resultList results
	bestCromosoma []int
	bestResult int

}

//
type results []result

type result struct{
	cromosoma []int
	valor int
}


func (prob *simpleGeneticProblem) Init() channels.ProblemUpdate{

	file, err := os.Open("problems/genetic_problems/ch130.xml")

	if err != nil{
		fmt.Println(err)
	}

	defer file.Close()

	auxTSP := &TSPxml{}

	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(auxTSP); err != nil{
		fmt.Println(err)
		return channels.ProblemUpdate{}
	}

	prob.size = len(auxTSP.Cities)
	prob.resultList = make(results,0,10)
	prob.graph = make([][]int,prob.size)

	for i := range prob.graph{
		aux := make([]int,prob.size)
		for _,ed := range auxTSP.Cities[i].Edges {
			aux[ed.City] = int(ed.Cost)
		}
		prob.graph[i] = aux
	}
	
	alg, _ := prob.alg()
	return channels.ProblemUpdate{alg, nil}

}

func (prob *simpleGeneticProblem) NewResult(data []string){	

	fmt.Println("Nuevo resultado")

	if len(data) != prob.size{
		fmt.Println("Error->Incorrect data length: ", len(data)," ", prob.size)
		return
	}
	cromosoma := make([]int, prob.size)

	for i,v := range data{
		aux, err := strconv.ParseInt(v, 10, 0)

		if err != nil{
			fmt.Println("Error->string no puede convertirse en int: ", v )
			return
		}

		cromosoma[i] = int(aux)
	}
	//chequea no repetidos
	auxCheck := make([]int, prob.size)
	for _,v := range cromosoma{
		auxCheck[v] = 0
	}

	if len(auxCheck) != prob.size{
		fmt.Println("Error->Elementos repetidos")
	}

	prob.Lock()
	defer prob.Unlock()
	fit := prob.fitness(cromosoma)

	res := result{cromosoma,fit}
		prob.resultList = append(prob.resultList, res)
	
	if len(prob.resultList) == 10{
		
		sort.Sort(prob.resultList)
		
		bestResults := prob.resultList[:3]
		
		//Guardar el mejor para consulta
		prob.bestResult = bestResults[0].valor
		prob.bestCromosoma = bestResults[0].cromosoma
		
		bestCromosomas := make([][]int,3)
		
		for i, v := range bestResults{
			bestCromosomas[i] = v.cromosoma
		}
			
		update := channels.ProblemUpdate{Data:bestCromosomas}
		
		problemChan.SendUpdate(update)
		
		prob.resultList = make(results,0,10)

		fmt.Println(prob.bestResult)
		
	}

	return	

}

func (prob *simpleGeneticProblem) Loop(){
	return
}

func (prob simpleGeneticProblem) fitness(cromosoma []int) int{

	var valor int = 0

	for i,v := range cromosoma{
		if i == (prob.size-1){
			break
		}else{
			valor += prob.graph[v][cromosoma[i+1]]
		}
	}

	return valor
}


func (prob simpleGeneticProblem) alg() (string, error){
	file, err := os.Open("problems/genetic_problems/algorithm_without_graph.js")
 	if err != nil {
		return "", err
 	}
 	defer file.Close()


 	var alg string
 	scanner := bufio.NewScanner(file)
 	for scanner.Scan() {
   		alg += scanner.Text()
  	}

  	graphString, err := json.Marshal(prob.graph)

  	if err != nil{
  		fmt.Println("Error->No se puede marshalear prob.graph")
  	}

  	lastPart := "graph = "+ string(graphString) + ";" 

  	completeAlg := lastPart

 	return completeAlg, scanner.Err()
}


//sort functions

func (res results) Len() int{
	return len(res)
}

func (res results) Less(i, j int) bool{
	return res[i].valor > res[j].valor
}

func (res results) Swap(i, j int){
	res[i], res[j] = res[j], res[i]
}