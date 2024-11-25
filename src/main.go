package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    filePath := "../TestData/"

    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatalf("Failed to read file: %s", err)
    }

    fileContent := string(data)

    fmt.Println(fileContent)
}