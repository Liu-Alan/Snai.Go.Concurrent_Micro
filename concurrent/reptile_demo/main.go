package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// 这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
	// w代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 爬QQ邮箱
func GetQQEmail() {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/51480004")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	// 遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	// 1.去网站拿数据
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

// 爬邮箱
func GetEmail(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬IdCard
func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬链接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬图片
func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

func main() {
	// 爬QQ邮箱
	//GetQQEmail()

	// 爬邮箱
	GetEmail("https://tieba.baidu.com/p/51480004")
	// 爬链接
	GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	// 爬手机号
	GetPhone("https://www.zhaohaowang.com/")
	// 爬身份证号
	GetIdCard("https://zuoyinyun.com/tool/id-card")
	// 爬图片
	GetImg("https://www.cnblogs.com/alan-lin/p/18572306")
}
