package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	AppVersion = "0.0.1"
)

var (
	configPath = flag.String("c", "autorun.config", "Profile path")
	targetPath = flag.String("t", ".", "The target path")
	version    = flag.Bool("v", false, "autorun version")
	showLog    = flag.Bool("i", false, "show log")
)

var (
	mainCmd *exec.Cmd
	config  *AutoRunConfig
	done    = make(chan bool)
)

func main() {

	flag.Parse()
	if *version {
		fmt.Print(AppVersion)
		os.Exit(0)
	}
	config = loadConfig(*configPath)

	buildWatcher(*targetPath, done, func(event fsnotify.Event) {
		Try(func() {
			printLog("[file] ", event.Op, event.Name)
			for _, e := range config.Exclude {
				if ok, _ := path.Match(e, event.Name); ok {
					printLog("[file ignore] ", event.Op, event.Name)
					return
				}
			}
			for _, i := range config.Include {
				if ok, _ := path.Match(i, event.Name); ok {
					printLog("[file Match] ", event.Op, event.Name)
					restart()
					return
				}
			}
		}).Catch(func(i interface{}) {
			printLog("[error] ", i)
		})
	}, func(err error) {
		printLog("[error] ", err)
	})
}

func printLog(v ...interface{}) {
	if *showLog {
		log.Println(v...)
	}
}

func runCmd(command Command) *exec.Cmd {
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("[error] ", err)
	}
	return cmd
}

func restart() {
	if mainCmd != nil {
		if err := mainCmd.Process.Kill(); err != nil {
			log.Println("[error] ", err)
		}
	}
	for _, cmd := range config.Build {
		runCmd(cmd)
	}
	go runCmd(config.Run)
}
