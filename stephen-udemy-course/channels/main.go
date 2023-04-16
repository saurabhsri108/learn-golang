package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.youtube.com",
		"http://www.tmall.com",
		"http://www.baidu.com",
		"http://www.qq.com",
		"http://www.sohu.com",
		"http://www.facebook.com",
		"http://www.taobao.com",
		"http://www.360.cn",
		"http://www.jd.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.wikipedia.org",
		"http://www.weibo.com",
		"http://www.sina.com",
		"http://www.reddit.com",
		"http://www.netflix.com",
		"http://www.microsoft.com",
		"http://www.zoom.us",
		"http://www.office.com",
		"http://www.live.com",
		"http://www.instagram.com",
		"http://www.alipay.com",
		"http://www.csdn.net",
		"http://www.bing.com",
		"http://www.microsoftonline.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// wait for the channel to written some link and then pass the link to checkLink again
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	defer resp.Body.Close()
	fmt.Println(link, "is up")
	c <- link
}
