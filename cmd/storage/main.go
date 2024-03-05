package main

import (
	"fmt"
	"log"

	"storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	f, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(f.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello world", restoredFile)
}
