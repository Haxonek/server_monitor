package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

// wait time between checking log file and posting to server
const WAIT_SEC int = 7200

func getLogFiles(file string) []string {
    // read file
    f, err := os.Open(file)
    if err != nil {
        fmt.Println("Error reading in server.txt")
        os.Exit(0)
    }
    log_file_list := bufio.NewScanner(f)

    // initialize file list array
    var files int = 0
    var log_files []string = make([]string, files) // allocates no memory

    for log_file_list.Scan() {
        var cur_line string = log_file_list.Text()
        // remove empty lines
        if strings.Trim(cur_line, " ") != "" {
            // increase memory allocation and add file to list
            files++
            tmp := log_files
            log_files = make([]string, files)
            copy(log_files, tmp)
            log_files[files-1] = cur_line

        }
    }

    return log_files
}

func getRecentLine(fileName string) string {
    var lastLine, tmp string;

    // Open opens as read-only
    f, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error reading in log file ", fileName)
        os.Exit(0)
    }
    logFile := bufio.NewScanner(f)

    // read only the first line
    for logFile.Scan() && lastLine == "" {
        tmp = strings.Trim(logFile.Text()," ") // reads first line, trims spaces
        if tmp != "" {
            lastLine = tmp
        }
    }

    return lastLine
}

func postToS3(logFile string) bool {
    // docs https://aws.amazon.com/sdk-for-go/
    // post last ten lines to S3 bucket to be read by client
}

func watchClosely(logFile string, watching *map[string]bool) {
    // here I'll want to watch the log file ever ~5 minutes, if it updates,
    // then the bot has resolved the issue itself and we can take down the error
    // file on S3. In my first iteration this will likely overwrite issue if two
    // go down, then one comes up and the other stays down

    // NOTE! mutex the watching has table!

}

func main() {
    fmt.Println("Starting main...")

    // read in log file URL's
    logFiles := getLogFiles(os.Args[1])

    var lastLogLine map[string]string = make(map[string]string)
    var watching map[string]bool = make(map[string]bool)

    // get the first line of each file; initialize
    for _, v := range logFiles {
        // read last line of the log file
        lastLogLine[v] = getRecentLine(v);
    }

    // infinite loop
    for true {

        // wait to allow bots to repost content
        time.Sleep(WAIT_SEC * time.Seconds())

        for _, v := range logFiles {
            // check to see if the most recent logs are the same
            // if they are, that means it's likely not updating properly
            if lastLogLine[v] == getRecentLine(v) && !watching[v] {
                postToS3(v)

                watching[v] = true
                go watchClosely(v, &watching)
            }
        }
    }

}
