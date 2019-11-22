package main

import (
    "fmt"
    "os"
    "net/http"
    "io"
    "strings"
)

func main() {
    for _,url := range os.Args[1:] {
        //1.8
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch error %v\n",err)
            os.Exit(1)
        }
        //1.9
        fmt.Fprintf(os.Stdout,resp.Status)
        //后面的参数要实现io.Reader接口，前面的是io.Writer
        if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
            fmt.Fprintf(os.Stderr, "fetch error %v\n",err)
            os.Exit(1)
        }
        defer resp.Body.Close()
    }
}
