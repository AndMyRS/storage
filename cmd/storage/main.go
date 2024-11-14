package main

import (
	"fmt"
	"log"

	"github.com/ASlavchik/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file was uploaded", file)

	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("restored file", restoredFile)

}
