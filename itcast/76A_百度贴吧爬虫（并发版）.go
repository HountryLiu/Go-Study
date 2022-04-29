package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

/**
 * 实际上爬虫一共就四个主要步骤
 * 1）明确目标（要知道你准备在哪个范围或者网站去搜索）
 * 2）爬（将所有的网站的内容全部爬下来）
 * 3）取（去掉对我们没用的数据）
 * 4）处理数据（按照我们想要的方式存储和利用）
 */

// http://tieba.baidu.com/f?kw=%E8%A1%A1%E6%B0%B4%E5%AD%A6%E9%99%A2&ie=utf-8&pn=0
// http://tieba.baidu.com/f?kw=%E8%A1%A1%E6%B0%B4%E5%AD%A6%E9%99%A2&ie=utf-8&pn=50
// http://tieba.baidu.com/f?kw=%E8%A1%A1%E6%B0%B4%E5%AD%A6%E9%99%A2&ie=utf-8&pn=100
func main() {
	var start, end int;
	fmt.Printf("请输入起始页（>= 1）：");
	fmt.Scan(&start);
	fmt.Printf("请输入终止页（>= 起始页）：");
	fmt.Scan(&end);

	//
	SpiderPages(start, end);
}

/**
 * 知道为什么要channel吗
 * 当主程提前执行完任务之后退出，协程连任务都没执行完，就关了
 */
func SpiderPages(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面", start, end);
	page := make(chan int);

	//
	for i := start; i <= end; i++ {
		go SpiderPage(i, page);
	}

	//
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page); // 这就是为了收工作的
	}
}

func SpiderPage(i int, page chan<- int) {
	baseUrl := "http://tieba.baidu.com/f?kw=%E8%A1%A1%E6%B0%B4%E5%AD%A6%E9%99%A2&ie=utf-8&pn=";
	url := baseUrl + strconv.Itoa((i-1)*50);
	fmt.Println("url =", url);

	// 2）爬（将所有网站的内容抓取下来
	result, httpGetErr := HttpGet(url);
	if httpGetErr != nil {
		fmt.Println("HttpGetErr =", httpGetErr);
	}

	// 3）把内容写进文件
	fileName := "./iofiles/tieba.hsnc-" + strconv.Itoa(i) + ".html";
	f, fCreateErr := os.Create(fileName);
	if fCreateErr != nil {
		fmt.Println("fCreateErr =", fCreateErr);
		return;
	}
	f.WriteString(result);
	f.Close();

	// 爬完了写一个
	page <- i;
}

func HttpGet(url string) (result string, err error) {
	resp, httpGetErr := http.Get(url);
	if httpGetErr != nil {
		return "", httpGetErr;
	}

	//
	defer resp.Body.Close();

	//
	buf := make([]byte, 1024*4);
	for {
		n, respReadErr := resp.Body.Read(buf);
		if n == 0 { // 读取结束或者出问题
			fmt.Println("respReadErr =", respReadErr);
			break;
		}

		//
		result += string(buf[:n]);
	}

	return result, nil;
}
