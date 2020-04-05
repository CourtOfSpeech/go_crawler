package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

//Fetch 根据传入的url地址，获取对应的内容返回
func Fetch(URL string) ([]byte, error) {
	//打开页面,直接用http.Get容易遇到403
	//resp, err := http.Get(url)
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	//resp 是一种资源，用完后需要关闭
	defer resp.Body.Close()
	//判断返回的信息中头部中的StatusCode==200
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error: staus code", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	//如果有乱码, 这个需要第三方库：golang.org/x/text 的支持
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//如果要自动识别抓取页面的字符编码
	readerBody := bufio.NewReader(resp.Body)
	e := determineEncoding(readerBody)
	utf8Reader := transform.NewReader(readerBody, e.NewDecoder())
	//如果抓取的页面编码格式不为utf-8,
	//all, err := ioutil.ReadAll(utf8Reader)
	//抓取页面的Body,这种方式没有处理字符编码，有可能会乱码
	//all, err := ioutil.ReadAll(resp.Body)

	return ioutil.ReadAll(utf8Reader)
}

//封装一个方法，用来返回抓取页面的字符编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//将数据转化为1024的[]byte
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	//func charset.DetermineEncoding(content []byte, contentType string)
	//(e encoding.Encoding, name string, certain bool)
	//e 即字符编码，name，名称，certain 返回的字符编码是否正确
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
