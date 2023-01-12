package main

import (
	"fmt"
)

func main() {
	kvStore, err := GetKVStore()
	if err != nil {
		panic(err.Error())
	}

	kvStore.Set([]byte("temp"), []byte("hello"))
	b := kvStore.Get([]byte("temp"))
	fmt.Printf("value = %v \n", string(b))
	commitId := kvStore.Commit()
	fmt.Printf("commit id = %v \n", commitId)
}
