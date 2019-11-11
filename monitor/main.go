package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    fmt.Println("Starting main...")

    // read file
    log_file_list, err := bufio.NewScanner("server.txt")
    if err != nil {
        fmt.Println("Error reading in server.txt")
        os.Exit(0)
    }

    // initialize file list array
    

    for log_file_list.Scanner() {
        var tmp string = log_file_list.Text()
        // remove empty lines
        if tmp.trim() != "" {

        }
    }

}
