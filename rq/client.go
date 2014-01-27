package rq

import(
	"math/rand"
)


type Client struct{
	id int
}



func NewClient(){

	id = rand.Uint32()
}