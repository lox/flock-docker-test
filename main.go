package main

import (
	"log"
	"time"

	"github.com/gofrs/flock"
)

const (
	lockFile = "go-lock.lock"
)

func main() {
	fileLock := flock.New(lockFile)

	log.Printf("Trying lock %s", lockFile)
	locked, err := fileLock.TryLock()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Got %v from lock", locked)

	if locked {
		log.Printf("Working...")
		time.Sleep(10 * time.Second)

		log.Printf("Unlocking %s", lockFile)
		if err := fileLock.Unlock(); err != nil {
			log.Fatal(err)
		}
	}
}
