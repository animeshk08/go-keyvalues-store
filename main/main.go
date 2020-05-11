package main

import (
	"fmt"
	kv "keyval/keyval"
)

//Example file to show usage of keyval package

func main() {

	// Open db
	store, err:= kv.Open("/path/to/store.db")

	if err!=nil{
		fmt.Println("Error is: #v", err)
	}

	var sessionVal = "sesh-123"

	// store a object
	store.Put("sess-341356", sessionVal)

	// Retrieve object
	store.Get("sess-341356", &sessionVal)

	// delete key
	store.Delete("sess-341356")

	// Close db
	store.Close()
}
