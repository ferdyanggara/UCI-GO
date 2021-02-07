package main

import ("fmt")

func main(){
	var convertNum float64
	fmt.Printf("Please input a floating number:")
	fmt.Scan(&convertNum)
	fmt.Printf("%v \n", int(convertNum))	
}


