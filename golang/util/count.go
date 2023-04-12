package main

import (
    "bufio"
    "fmt"
    "os"
)

func getcount(file_path string) {
    buffer_size := 1024 * 1024  // 1MB buffer
    line_count := 0

    file, err := os.Open(file_path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := bufio.NewReaderSize(file, buffer_size)
    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        line_count++
    }

    fmt.Printf("%s has %d lines\n", file_path, line_count)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run count.go <file_path>")
        return
    }
    file_path := os.Args[1]
    getcount(file_path)
}

