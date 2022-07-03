package hotreload

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Watcher ...
func Watcher() {
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
				if needReload(event.Op) {
					fn, ok := callbackTable.Load(event.Name)
					if ok {
						fn.(CallbackFunc)(event.Name)
					}
				} else {
					log.Printf("other event ignore: %s\n", event)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	addWatchDir(watcher, watchDirectory)

	<-done
}

func addWatchDir(watcher *fsnotify.Watcher, dir string) error {
	fmt.Println("dir", dir)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		return watcher.Add(path)
	})
}

// 增删改都需要reload
func needReload(op fsnotify.Op) bool {
	if op&fsnotify.Create == fsnotify.Create {
		return true
	}

	if op&fsnotify.Remove == fsnotify.Remove {
		return true
	}

	if op&fsnotify.Write == fsnotify.Write {
		return true
	}

	if op&fsnotify.Chmod == fsnotify.Chmod {
		return true
	}

	if op&fsnotify.Rename == fsnotify.Rename {
		return true
	}

	return false
}
