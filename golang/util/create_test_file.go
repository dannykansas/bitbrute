package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "strconv"
    "time"
)

func main() {
    // Number of lines to write
    numLines := 1000000000

    // Open file in write mode
    file, err := os.Create("output_file.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Use buffered writer to write lines in batches
    start := time.Now()
    writer := bufio.NewWriter(file)
    batchSize := 1000
    for i := 0; i < numLines; i += batchSize {
        lines := make([]string, batchSize)
        for j := 0; j < batchSize; j++ {
            lines[j] = strconv.FormatFloat(rand.Float64(), 'f', 6, 64) + "\n"
        }
        _, err := writer.WriteString(strings.Join(lines, ""))
        if err != nil {
            panic(err)
        }
    }
    // Flush buffer to disk
    err = writer.Flush()
    if err != nil {
        panic(err)
    }
    elapsed := time.Since(start)

    fmt.Printf("Total time taken: %v\n", elapsed)
}

