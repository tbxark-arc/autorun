package main

import (
    "github.com/fsnotify/fsnotify"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func buildWatcher(path string, done chan bool, eventHandler func(event fsnotify.Event), errorHandler func(err error))  {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                eventHandler(event)
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                errorHandler(err)
            }
        }
    }()

    watcher.Add(path)
    wErr := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() && strings.Index(path, "/.") < 0 {
            if aErr := watcher.Add(path); aErr != nil {
                return aErr
            }
        }
        return nil
    })
    if wErr != nil {
        log.Fatal(err)
    }
    <-done
}
