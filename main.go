package main

import (
    "fmt"
    "log"
    "os/exec"
    "os"
    "strings"
    // "reflect"
)

func main() {
    parseConfig()
    
    cmd := exec.Command("/bin/sh", "-c", config.Command)
    
    var file *os.File
    if config.Logfile != "" {
        createFile(config.Logfile)
        file, _ = os.OpenFile(config.Logfile, os.O_APPEND|os.O_WRONLY, 0644)
        defer file.Close()
    }

    var out = outstream{file}
    cmd.Stdout = out
    
    var outerr = outerr{file}
    cmd.Stderr = outerr
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(cmd.Wait())
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func createFile(filename string) {
    // detect if file exists
    var _, err = os.Stat(filename)

    // create file if not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(filename)
        check(err)
        defer file.Close()
    }
}

func writeFile(str string, file *os.File) {
    // write some text to file
    file.WriteString(str)

    // save changes
    file.Sync()
}

type outstream struct{
    file *os.File
}

func (out outstream) Write(p []byte) (int, error) {
    var outputText string = string(p)
    fmt.Println(outputText)

    // specific to bunyan
    if sliceContains(config.IncludeErrorKeywords, outputText) == false {
        // performAction(val, errorText)
    }

    writeFile(outputText, out.file)

    // fmt.Println(reflect.TypeOf(out.file))

    return len(p), nil
}

type outerr struct{
    file *os.File
}

func (out outerr) Write(p []byte) (int, error) {
    var errorText string = string(p)
    fmt.Println(errorText)

    writeFile(errorText, out.file)

    if sliceContains(config.ExcludeErrorKeywords, errorText) == false {
        for _, val := range config.Actions {
            performAction(val, errorText)
        }
    }

    return len(p), nil
}

func performAction(actionType string, logText string){
    if actionType == "twilio" {
        sendText(logText)
    }
    if actionType == "gmail" {
        // TODO need to implement gmail action
        fmt.Println("gmail!!!")
    }
}

func sliceContains(keywords []string, searchText string) bool {
    for _, keyword := range keywords {
        if strings.Contains(searchText, keyword) {
            return true
        }
    }
    return false
}