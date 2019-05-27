package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	login_url string= "https://account.ccnu.edu.cn/cas/login"
	//login_url string= "https://baidu.com"
)
func get_html(requesturl, username, password string) error {
	jar, _ := cookiejar.New(nil)
	//client
	c := &http.Client{}
	c.Jar = jar

	res1,_ := c.Get(login_url)
	doc, _ := goquery.NewDocumentFromReader(res1.Body)
	defer res1.Body.Close()
	x := doc.Find("div[id=login]").Find("section[class*=btn]")
	lt, _ := x.Find("input[name=lt]").Attr("value")
	execution, _ := x.Find("input[name=execution]").Attr("value")
	fmt.Println(lt,execution)
	form := url.Values{
		"username":	{username},
		"password":	{password},
		"lt":		{lt},
		"execution":{execution},
		"_eventld":	{"submit"},
		"submit":	{"登录"},
	}
	res2, err := c.PostForm(login_url, form)
	doc1, _ := goquery.NewDocumentFromReader(res2.Body)
	defer res2.Body.Close()
	fmt.Println(doc1.Html())
	fmt.Println(res2.Header)
	fmt.Println(res2.StatusCode)
	fmt.Println(err)
	return err
}

func main() {
	get_html(login_url,"","")
}
