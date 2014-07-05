package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/MichaelTJones/walk"
)

// tour.go: Take a tour of one or more file hierarchies

func main() {
	t0 := time.Now()
	var files, bytes int64
	var lock sync.Mutex
	visitor := func(path string, info os.FileInfo, err error) error {
		if err == nil {
			lock.Lock()
			files++
			bytes += info.Size()
			lock.Unlock()
		}
		return nil
	}
	for _, root := range os.Args[1:] {
		walk.Walk(root, visitor)
	}
	𝛥t := float64(time.Since(t0)) / 1e9

	fmt.Printf("walked %d files containing %d bytes in %.3f seconds\n", files, bytes, 𝛥t)
}
