package main

import (
	"checkIn/checkIn"
	"log"
)

func main() {
	c,err:=checkIn.SetUp()
	if err!=nil{
		log.Println(err)
		return 
	}
	err=c.SignIn()
	if err!=nil{
		log.Println(err)
		return 
	}
}