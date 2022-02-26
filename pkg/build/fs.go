package build

import (
	// "crypto/md5"
	// "crypto/sha256"
	// "fmt"
	"io"
	"log"
)

func logAndClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println("error closing:", err)
	}
}

// func shaDir(dir string) error {
// 	return filepath.Walk()
// }
