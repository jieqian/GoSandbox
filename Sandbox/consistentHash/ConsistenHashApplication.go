package main

import (
	//rings "GoSandbox/Sandbox/consistentHash/rings"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var myMap map[string]int

func init(){
	println("initial testing .............")
}

func main() {
	//println(RandStringBytes(rand.Intn(10)))
	//println(RandStringBytes(rand.Intn(10)))
	//println(RandStringBytes(rand.Intn(10)))

	myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	println(myMap["two"])
	println(len(myMap))

	//var numberOfReplicas int32 = 3
	//nodes := []string{}
	//nodes = append(nodes,"node1","node2","node3")
	//nodeRing := rings.NewConsistentHashNodesRing(numberOfReplicas,nodes)
	//nodeRing.SetNodes([]string{"node4"})
	//rv,_ := nodeRing.GetNode("xx")
	//println(rv)
	//nodeRing.RemoveNodes([]string{"node3"})
	//rv2,_ := nodeRing.GetNode("xx")
	//println(rv2)
}
