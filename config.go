package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "strings"
)

type AutoRunConfig struct {
    Build   []Command
    Run     Command
    Include []string
    Exclude []string
}

type Config struct {
    Build   []Command `json:"build"`
    Run     Command   `json:"run"`
    Include Template  `json:"include"`
    Exclude Template  `json:"exclude"`
}

type Command struct {
    Name string          `json:"name"`
    Args []string       `json:"args"`
}

type Template struct {
    Import  []string  `json:"import"`
    Pattern []string  `json:"pattern"`
}

func loadTemplate(template Template) []string  {
    res := template.Pattern
    for _, i := range template.Import {
        file, err := ioutil.ReadFile(i)
        if err != nil {
            log.Fatal(err)
        }
        res = append(res, strings.Split(string(file), "\n")...)
    }
    return res
}

func loadConfig(path string) AutoRunConfig  {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    config := Config{}
    if jErr := json.Unmarshal(file, &config); jErr != nil {
        log.Fatal(jErr)
    }
    include := loadTemplate(config.Include)
    exclude := loadTemplate(config.Exclude)
    return AutoRunConfig{
        config.Build,
        config.Run,
        include,
        exclude,
    }
}