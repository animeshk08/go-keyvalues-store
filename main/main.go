package main

import (
	"fmt"
	kv "keyval/keyval"
)

//Example file to show usage of keyval package

func main() {

	// Open db
	store, err:= kv.Open("./store.db")

	if err!=nil{
		fmt.Println("Error is: #v", err)
	}

	var sessionVal = "sesh-123"

	// store a object
	store.Put("sess-341356", sessionVal)
	fmt.Println("Value inserted")
	var val string
	// Retrieve object
	store.Get("sess-341356", &val)
	fmt.Println("Value retrieve Value.")
	fmt.Println("Value is", val)

	// delete key
	store.Delete("sess-341356")
	fmt.Println("Value deleted")
	var delVal string

	err =store.Get("sess-341356", &delVal)
	fmt.Println(err)

	// Close db
	store.Close()
}
