package conf

import(
	"fmt"
	"testing"
	)

func TestLoadConf(t *testing.T){
	LoadConf()
	fmt.Println("Ip:", ip," Port:", port, " Cors:", cors, " Problem:", problem)
}