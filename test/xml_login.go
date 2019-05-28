package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Form struct {
	lt	string	`xml:"charset,attr"`
}

func main() {
	html, _ := http.Get("https://account.ccnu.edu.cn/cas/login")
	data, _ := ioutil.ReadAll(html.Body)
	form := Form{}
	_ = xml.Unmarshal(data, form)
	fmt.Println(form)
}