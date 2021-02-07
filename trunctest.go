package main 

import ("fmt")


func main(){
	var numToConvert  float64
	fmt.Scan(&numToConvert)
	fmt.Printf("int part %v float part %v ", parseFloatToInt(numToConvert))
	fmt.Println(" ")
}

func parseFloatToInt(floatNum float64){
	floatToNum = int(floatNum)
	x := floatNum - floatToNum
	return x, floatToNum
}
