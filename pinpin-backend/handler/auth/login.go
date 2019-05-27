package auth

import (
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	"net/http"
)

const (
	login_url string= "https://account.ccnu.edu.cn/cas/login"
	//login_url string= "https://baidu.com"
)
func try_login(requesturl, username, password string) error {
	//client
	c := &http.Client{}
	req1, _ := http.NewRequest("GET", requesturl, nil)
	req1.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36")
	res, _ := c.Do(req1)
	cookie := res.Header
	fmt.Println(cookie)
	//doc, _ := goquery.NewDocumentFromReader(res.Body)
	////fmt.Println(reflect.TypeOf(doc))
	////fmt.Println(doc.Html())
	//x := doc.Find("div[id=login]").Find("section[class*=btn]")
	//lt, _ := x.Find("input[name=lt]").Attr("value")
	//execution, _ := x.Find("input[name=execution]").Attr("value")
	//
	//req2, _ := http.NewRequest("POST", requesturl, nil)
	//req2.PostForm.Set("username", username)
	//req2.PostForm.Add("password", password)
	//req2.PostForm.Add("lt", lt)
	//req2.PostForm.Add("execution", execution)
	//req2.PostForm.Add("_eventld", "submit")
	//req2.PostForm.Add("submit", "登录")
	//req2.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36")
	//req2.Header.Add("Cookie", cookie)

}

