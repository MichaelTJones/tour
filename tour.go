package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/MichaelTJones/walk"
)

// tour.go: Take a tour of one or more file hierarchies

func main() {
	fmt.Printf("filepath.Walk(%v)\n", os.Args[1:])
	for i := 0; i < 3; i++ {
		walker(filepath.Walk)
	}
	fmt.Println()
	fmt.Printf("walk.Walk(%v)\n", os.Args[1:])
	for i := 0; i < 3; i++ {
		walker(walk.Walk)
	}
}

func walker(w func(string, walk.WalkFunc) error) {
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
		err := w(root, visitor)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
	𝛥t := float64(time.Since(t0)) / 1e9
	fmt.Printf("walked %d files containing %d bytes in %.3f seconds\n", files, bytes, 𝛥t)
}
