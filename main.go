package main

import "fmt"
import "github.com/prnvkv/grpcTest/test"

func main(){
	fmt.Println("Get dep")
	test.TestSort()
	test.TestFiltering()
}

