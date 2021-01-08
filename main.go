package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	accessKey := os.Getenv("SPACES_KEY")
	secKey := os.Getenv("SPACES_SECRET")
	endpoint := "fra1.digitaloceanspaces.com"

	ssl := true

	// client, err := minio.New(endpoint, accessKey, secKey, ssl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/home")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
