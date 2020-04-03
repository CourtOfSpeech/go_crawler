package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	//打开页面
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	//resp 是一种资源，用完后需要关闭
	defer resp.Body.Close()
	//判断返回的信息中头部中的StatusCode==200
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: staus code", resp.StatusCode)
		return
	}
	//如果有乱码, 这个需要第三方库：golang.org/x/text 的支持
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//如果要自动识别抓取页面的字符编码
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//如果抓取的页面编码格式不为utf-8,
	all, err := ioutil.ReadAll(utf8Reader)
	//抓取页面的Body,这种方式没有处理字符编码，有可能会乱码
	//all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}

//封装一个方法，用来返回抓取页面的字符编码格式
func determineEncoding(r io.Reader) encoding.Encoding {
	//将数据转化为1024的[]byte
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	//func charset.DetermineEncoding(content []byte, contentType string)
	//(e encoding.Encoding, name string, certain bool)
	//e 即字符编码，name，名称，certain 返回的字符编码是否正确
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
