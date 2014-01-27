package rq

import(
	"fmt"
	"net/http"
)

//Give a problem to a client
func LoadAnt(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, `{'func':"windows.alert('Funciona!');"}`)
}