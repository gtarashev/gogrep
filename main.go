package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func search(filename, needle string, name bool) (int, error) {
    file, err := os.Open(filename)
    defer file.Close()
    if err != nil {
        return 0, err
    }

    occurances := 0
    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        if strings.Contains(line, needle) {
            if occurances == 0 && name {
                fmt.Println(filename)
            }
            occurances += 1
            fmt.Println(line)
        }
    }

    return occurances, nil
}

func main() {
    args := os.Args[1:]
    if len(args) < 2 {
        fmt.Println("Not enough arguments supplied")
        os.Exit(1)
    }
    needle := args[0]
    files := args[1:]

    name := false
    for _, filename := range files {
        if len(files) != 1 {
            name = true
        }

        occurances, err := search(filename, needle, name)
        if err != nil {
            fmt.Println("Error:", err)
            fmt.Println()
        }
        if occurances > 0 {
            fmt.Println()
        }
    }
}
