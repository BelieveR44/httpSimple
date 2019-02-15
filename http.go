package httpSimple

import (
	"bytes"
	//	"fmt"
	"io/ioutil"

	//	"fmt"
	"compress/gzip"
	"log"
	"time"

	"github.com/kirinlabs/HttpRequest"
)

func SGet(url string, query map[string]interface{}, headers, cookies map[string]string, timeout int) ([]byte, error) {
	//var resp *HttpRequest.Response
	req := HttpRequest.NewRequest()
	// 设置超时时间，不设置时，默认30s
	req.SetTimeout(time.Duration(timeout) * time.Second)
	// 设置Headers
	req.SetHeaders(headers)
	// 设置Cookies
	req.SetCookies(cookies)
	// GET 默认调用方法
	resp, err := req.Get(url, nil)
	// GET 传参调用方法
	// 第2个参数默认为nil，也可以传参map[string]interface{}
	// 第2个参数不为nil时，会把传入的map以query传参的形式重新构造新url
	// 新的URL: http://127.0.0.1:8000?name=flyfreely&id=1&title=csdn
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	if resp.StatusCode() == 200 {
		body, err := resp.Body()
		if err != nil {
			log.Println(err)
			return []byte(""), err
		}
		//fmt.Println(resp.Response().Header.Get("Content-Encoding"))
		if resp.Response().Header.Get("Content-Encoding") == "gzip" {
			//var unbody string
			r, _ := gzip.NewReader(bytes.NewReader(body))
			defer r.Close()
			return ioutil.ReadAll(r)

		}
		return body, nil
	} else {
		return []byte(""), nil
	}
}
func SPost(url string, data map[string]interface{}, query map[string]interface{}, headers, cookies map[string]string, timeout int) ([]byte, error) {
	req := HttpRequest.NewRequest()
	// 设置超时时间，不设置时，默认30s
	req.SetTimeout(time.Duration(timeout) * time.Second)
	// 设置Headers
	req.SetHeaders(headers)
	// 设置Cookies
	req.SetCookies(cookies)
	// Post 默认调用方法
	resp, err := req.Post(url, data)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	if resp.StatusCode() == 200 {
		body, err := resp.Body()
		if err != nil {
			log.Println(err)
			return []byte(""), err
		}
		if resp.Response().Header.Get("Content-Encoding") == "gzip" {
			//var unbody string
			r, _ := gzip.NewReader(bytes.NewReader(body))
			defer r.Close()
			return ioutil.ReadAll(r)

		}
		return body, nil
	} else {
		return []byte(""), nil
	}

}
