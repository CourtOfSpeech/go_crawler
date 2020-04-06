package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

//限制一下发送请求的时间,不然请求太快，会卡住
var rateLimiter = time.Tick(2 * time.Second)

//Fetch 根据传入的url地址，获取对应的内容返回
func Fetch(URL string) ([]byte, error) {
	<-rateLimiter
	//打开页面,直接用http.Get容易遇到403
	//resp, err := http.Get(url)
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	// if strings.Contains(URL, "https://album.zhenai.com/u") {
	// 	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// 	request.Header.Set("accept-encoding", "gzip, deflate, br")
	// 	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	// 	request.Header.Set("referer", URL)
	// 	request.Header.Set("sec-fetch-dest", "document")
	// 	request.Header.Set("sec-fetch-mode", "navigate")
	// 	request.Header.Set("ssec-fetch-site", "same-origin")
	// 	request.Header.Set("upgrade-insecure-requests", "1")

	// 	cookie := &http.Cookie{Name: "FSSBBIl1UgzbN7N443S",
	// 		Value: "uH8MP_P2KNTXzjKPmsOLdKKCqZO1b2lJlmLOmH6HM7S.rn2LRxJPeM5qh5HjxWcA"}
	// 	request.AddCookie(cookie)
	// 	cookie = &http.Cookie{Name: "sid",
	// 		Value: "93183818-9b99-4d2e-a183-53d4aff3a5a5"}
	// 	request.AddCookie(cookie)
	// 	cookie = &http.Cookie{Name: "Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2",
	// 		Value: "1586098837"}
	// 	request.AddCookie(cookie)
	// 	cookie = &http.Cookie{Name: "Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2",
	// 		Value: "1586098878"}
	// 	request.AddCookie(cookie)
	// 	cookie = &http.Cookie{Name: "FSSBBIl1UgzbN7N443T",
	// 		Value: "4WQDvNoOwoCE0XmOq7kpWDm1huXwrx6DNqutv0nkrcG649fS8KS2dUgu.GW3YjbpyHBspvmahFvlFfretOF1xg2GPmbakCgWxezERC8VRpnDBX_oZkOqnxEcgb7kcaXJiqDzdf_M4jglYUyyukZKsf4VPyFogYA7MLwFmywEPxyYOQpKLymhCr7kfujMCiVw_uh_uM3xuL_4ZvjN6eufsiqUaxBU6t8pjHcFPhlpWpsZSa6csrILrMkD.PSe.iyvCPicJx6sRdo.QqOt41UVzVo3iIW94EYeERVAfpQjEnZIv5mallEazU.AldzArGD6r_C2tCksVtfmQ0FwVFAVB7yY2cSbYor3BIzJbXKEPT8ol.fDjxGvxI1v1no6COG5d8LQ"}
	// 	request.AddCookie(cookie)
	// }

	//fmt.Println(request)
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
		//panic(err)
		log.Printf("determineEncoding error %v", err)
	}
	//func charset.DetermineEncoding(content []byte, contentType string)
	//(e encoding.Encoding, name string, certain bool)
	//e 即字符编码，name，名称，certain 返回的字符编码是否正确
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
