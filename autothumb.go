package main
/**
 Auto thumbnail based on file system notification(fsnotify) and imagemagic with golang
  */
import (
	"log"
	"os/exec"
	"strings"
	"flag"
	"github.com/howeyc/fsnotify"
	"time"
)

func main() {
	log.Println("start autothumb")
	folder := flag.String("folder", "/var/www/images", "the folder to monitor")
	target := flag.String("target", "/var/www/images/thumbnail", "target thumbs path")
	size := flag.String("size", "341x267", "thumb image size")
	wait := flag.Int64("wait", 10, "Time to wait the image written done before generating thumb image")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsCreate() {
					time.Sleep(time.Duration(*wait)*time.Second)
					f := ev.Name
					i := strings.LastIndexAny(f, ".")
					log.Println(f)
					if i != -1 {
						format := f[i + 1:]
						//log.Println(format, f)
						err := exec.Command("mogrify", format, "jpg", "-path", *target, "-thumbnail", *size, f).Run()
						if err != nil {
							log.Println(err)
						}
					}
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
		done<-true
	}()

	err = watcher.Watch(*folder)
	if err != nil {
		log.Fatal(err)
	}

	select {
	case <-done:
		watcher.Close()
		log.Println("done,exit autothumb")
	}

}
