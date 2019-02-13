package main

import (
	"HttpRequest"
	"fmt"
	"log"
)

func main() {

	req := HttpRequest.NewRequest()

	// 设置超时时间，不设置时，默认30s
	req.SetTimeout(5)

	// 设置Headers
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded", //这也是HttpRequest包的默认设置
	})

	// 设置Cookies
	req.SetCookies(map[string]string{
		"sessionid": "LSIE89SFLKGHHASLC9EETFBVNOPOXNM",
	})

	postData := map[string]interface{}{
		"id":    1,
		"title": "csdn",
	}

	// GET 默认调用方法
	resp, err := req.Get("http://127.0.0.1:8000?name=flyfreely", nil)

	// GET 传参调用方法
	// 第2个参数默认为nil，也可以传参map[string]interface{}
	// 第2个参数不为nil时，会把传入的map以query传参的形式重新构造新url
	// 新的URL: http://127.0.0.1:8000?name=flyfreely&id=1&title=csdn

	//resp, err := req.Get("http://127.0.0.1:8000?name=flyfreely", postData)

	// POST 调用方法

	//resp, err := req.Post("http://127.0.0.1:8000", postData)

	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode() == 200 {
		body, err := resp.Body()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(body))
	}

	//或者打印Json
	if resp.StatusCode() == 200 {
		body, err := resp.Json()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(body)
	}
}
