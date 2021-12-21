package testDemo

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func TestHttp(t *testing.T) {
	http.HandleFunc("/", myfunc)
	http.ListenAndServe(":8080", nil)

	// 更多http.Server的字段可以根据情况初始化
	//server := http.Server{
	//	Addr:         ":8080",
	//	ReadTimeout:  0,
	//	WriteTimeout: 0,
	//}
	//http.HandleFunc("/", myfunc)
	//server.ListenAndServe()

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", myfunc)
	//http.ListenAndServe(":8080", mux)
}

func TestRequest(t *testing.T) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{Name: "userId", Value: strconv.Itoa(12345)}
	request.AddCookie(cookie)

	// 设置request的Header，具体可参考http协议
	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK, utf-8;q=0.7, *;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	request.Header.Set("Accept-Language", "zh-CN, zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	if response.StatusCode == http.StatusOK {
		body, err := gzip.NewReader(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		defer body.Close()
		r, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err)
		}
		// 打印出http Server返回的http Response信息
		fmt.Println(string(r))
	}
}

func TestGet(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(r))
}

func TestPost(t *testing.T) {
	// application/x-www-form-urlencoded：为POST的contentType
	// strings.NewReader("mobile=xxxxxxxxxx&isRemberPwd=1") 理解为传递的参数
	resp, err := http.Post("http://localhost:8080/login.do",
		"application/x-www-form-urlencoded", strings.NewReader("mobile=xxxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestPostForm(t *testing.T) {
	postParam := url.Values{
		"mobile":      {"xxxxxx"},
		"isRemberPwd": {"1"},
	}
	// 数据的键值会经过URL编码后作为请求的body传递
	resp, err := http.PostForm("http://localhost:8080/login.do", postParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func TestCustomerHandler(t *testing.T) {
	mux := http.NewServeMux()

	th := &timeHandler{time.RFC1123}
	mux.Handle("/", th)

	log.Println("Listening...")
	http.ListenAndServe(":3030", mux)
}

func timeHandlerV2(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func TestFuncHandle(t *testing.T) {
	mux := http.NewServeMux()

	// Convert the timeHandler function to a HandlerFunc type
	th := http.HandlerFunc(timeHandlerV2)
	// And add it to the ServeMux
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
          Welcome
        </p>
        </body>
</html>`
	fmt.Fprintln(w, html)
}

func TestServer(t *testing.T) {
	http.HandleFunc("/", index)

	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	server.ListenAndServe()
}
