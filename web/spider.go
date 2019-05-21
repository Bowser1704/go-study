package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	var url string
	fmt.Println("input url:")
	fmt.Scanln(&url)
	fmt.Println(url2html(url))
}

func url2html(url string) string {
	//url --> html
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error",err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("http read error",err)
		return ""
	}
	return string(body)
}
