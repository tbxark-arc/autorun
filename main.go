package main

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"path"
)


func main() {
	configPath := flag.String("c", "autorun.config", "Profile path")
	targetPath := flag.String("t", ".", "The target path")
	var cmdHolder *exec.Cmd
	flag.Parse()
	config := loadConfig(*configPath)
	done := make(chan bool)
	buildWatcher(*targetPath, done, func(event fsnotify.Event) {
		Try(func() {
			log.Println("[file] ", event.Op, event.Name)
			for _, e := range config.Exclude {
				if ok, _ := path.Match(e, event.Name); ok {
					log.Println("[file ignore] ", event.Op, event.Name)
					return
				}
			}
			for _, i := range config.Include {
				if ok, _ := path.Match(i, event.Name); ok {
					log.Println("[file Match] ", event.Op, event.Name)
					if cmdHolder != nil {
						cmdHolder.Process.Kill()
					}
					for _, cmd := range config.Build {
						build := exec.Command(cmd.Name, cmd.Args...)
						build.Stdout = os.Stdout
						build.Stderr = os.Stderr
						if err := build.Run(); err != nil {
							log.Println("[error] ", err)
						}
					}
					cmdHolder = exec.Command(config.Run.Name, config.Run.Args...)
					cmdHolder.Stdout = os.Stdout
					cmdHolder.Stderr = os.Stderr
					go func() {
						if err := cmdHolder.Run(); err != nil {
							log.Println("[error] ", err)
						}
					}()
					break
				}
			}
		}).Catch(func(i interface{}) {
			log.Println("[error] ", i)
		})
	}, func(err error) {
		log.Println("[error] ", err)
	})
}
