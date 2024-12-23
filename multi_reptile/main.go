package main

// 图床地址
// https://depositphotos.com/cn/photos/%E5%A4%A7%E8%87%AA%E7%84%B6.html?offset=0
// https://depositphotos.com/cn/photos/%E5%A4%A7%E8%87%AA%E7%84%B6.html?offset=100

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// 下载图片
func DownloadFile(url string, fileName string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	fileName = "img/" + fileName
	// 写出数据
	err = os.WriteFile(fileName, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 截取url名字
func GetFileNameFromUrl(url string) (fileName string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	fileName = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixMilli()))
	fileName = timePrefix + "_" + fileName
	return
}

// 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		fileName := GetFileNameFromUrl(url)
		ok := DownloadFile(url, fileName)
		if ok {
			fmt.Printf("%s 下载成功\n", fileName)
		} else {
			fmt.Printf("%s 下载失败\n", fileName)
		}
	}
	waitGroup.Done()
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

// 获取当前页图片链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string) {
	urls := getImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _, urlI := range urls {
		chanImageUrls <- urlI
	}

	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	waitGroup.Done()
}

// // 任务统计协程
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 10 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

func main() {
	// 1.初始化通道
	chanImageUrls = make(chan string, 10000)
	chanTask = make(chan string, 10)
	// 2.爬虫协程
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://depositphotos.com/cn/photos/%E5%A4%A7%E8%87%AA%E7%84%B6.html?offset=" + strconv.Itoa(i*100))
	}
	// 3.任务统计协程，统计10个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i := 0; i < 4; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}
