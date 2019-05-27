package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
)

const (
	login_url string= "https://account.ccnu.edu.cn/cas/login"
	//login_url string= "https://baidu.com"
)
func get_html(requesturl, username, password string) error {
	//client
	c := &http.Client{}
	req1, _ := http.NewRequest("GET", requesturl, nil)
	req1.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36")
	res, _ := c.Do(req1)

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	//fmt.Println(reflect.TypeOf(doc))
	//fmt.Println(doc.Html())
	x := doc.Find("div[id=login]").Find("section[class*=btn]")
	lt, _ := x.Find("input[name=lt]").Attr("value")
	execution, _ := x.Find("input[name=execution]").Attr("value")
	postParam := url.Values{
		"username":	{username},
		"password":	{password},
		"lt":		{lt},
		"execution":{execution},
		"_eventld":	{"submit"},
		"submit":	{"登录"},
	}

	res1, err1 := http.PostForm(requesturl, postParam)
	fmt.Println(res1.Header)
	fmt.Println(res1.Cookies())
	fmt.Println(err1)
	return err1
}

func main() {
	get_html(login_url,"2018212576","Yu@14796825550")
}
