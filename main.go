package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	//如果有乱码
	//抓取页面的Body
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}
