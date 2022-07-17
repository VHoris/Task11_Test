package main

import (
	     "task11/do"
		 "fmt"
       )

func main(){
   
    str,err:=do.Do("b",5,true)
	if err==nil{ 
	   fmt.Printf("%v",str)
	}
}