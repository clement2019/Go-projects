package main

import (
	"fmt"

	httplayer "github.com/clement2019/first-Go/Handler"
)



func main(){

	message :="You are welcomed to hello world"
	fmt.Println(message)

	//Now call the The serviceprovider function in httplayer package

	mH := httplayer.Serverprovider()
	mH.Myserver() 
}