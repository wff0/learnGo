package testDemo

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestColly(t *testing.T) {
	// NewCollector(options ...func(*Collector)) *Collector
	// 声明初始化NewCollector对象时可以指定Agent，连接递归深度，URL过滤以及domain限制等
	c := colly.NewCollector(
		//colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"))

	// 发出请求时附的回调
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", "baidu.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Referer", "http://www.baidu.com")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")

		fmt.Println("Visiting", r.URL)
	})

	// 对响应的HTML元素处理
	c.OnHTML("title", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		fmt.Println("title:", e.Text)
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		e.ForEach(".hotnews a", func(i int, el *colly.HTMLElement) {
			band := el.Attr("href")
			title := el.Text
			fmt.Printf("新闻 %d : %s - %s\n", i, title, band)
			// e.Request.Visit(band)
		})
	})

	// 发现并访问下一个连接
	//c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", r.StatusCode)
		// 设置context
		// fmt.Println(r.Ctx.Get("url"))
	})

	// 对visit的线程数做限制，visit可以同时运行多个
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})

	c.Visit("http://news.baidu.com")
}

func TestGoquery(t *testing.T) {
	urlstr := "https://news.baidu.com"
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	c := colly.NewCollector()
	// 超时设定
	c.SetRequestTimeout(100 * time.Second)
	// 指定Agent信息
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", u.Host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", u.Host)
		r.Headers.Set("Referer", urlstr)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("title:", e.Text)
	})

	c.OnResponse(func(resp *colly.Response) {
		fmt.Println("response received", resp.StatusCode)

		// goquery直接读取resp.Body的内容
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))

		// 读取url再传给goquery，访问url读取内容，此处不建议使用
		// htmlDoc, err := goquery.NewDocument(resp.Request.URL.String())

		if err != nil {
			log.Fatal(err)
		}

		// 找到抓取项 <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		htmlDoc.Find(".hotnews a").Each(func(i int, s *goquery.Selection) {
			band, _ := s.Attr("href")
			title := s.Text()
			fmt.Printf("热点新闻 %d: %s - %s\n", i, title, band)
			c.Visit(band)
		})

	})

	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
	})

	err = c.Visit(urlstr)
}
