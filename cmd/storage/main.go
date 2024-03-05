package storage

import (
	"fmt"

	"storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello world", st)
}
