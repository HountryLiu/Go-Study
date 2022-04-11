package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 网页规律：
// https://www.pengfue.com/xiaohua_1.html
// https://www.pengfue.com/xiaohua_2.html
// https://www.pengfue.com/xiaohua_3.html

// 主页面规律：
// <div class="list-item bg1 b1 boxshadow" id="1857269">

// <h1 class="dp-b"><a href="https://www.pengfue.com/content_1857269_1.html" target="_blank">标题</a></h1>
func main() {
	var start, end int;
	fmt.Printf("请输入起始页（>= 1）：");
	fmt.Scan(&start);
	fmt.Printf("请输入终止页（>= 起始页）：");
	fmt.Scan(&end);

	//
	SpiderPages(start, end);
}

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
	baseUrl := "https://www.pengfue.com/xiaohua_";
	// 1）明确目标（要知道你准备在哪个范围或者网站去搜索）
	url := baseUrl + strconv.Itoa(i) + ".html";
	fmt.Printf("正在爬取第%d个网页：%s\n", i, url);

	// 定义暂存数据，待会要写入文件的
	fileTitles := make([]string, 0);
	fileContents := make([]string, 0);

	// 2）爬（将所有网站的内容抓取下来
	result, httpGetErr := HttpGet(url); //fmt.Println(result);
	if httpGetErr != nil {
		fmt.Println("HttpGetErr =", httpGetErr);
	}

	// 解释表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))" target="_blank">`);
	if re == nil {
		fmt.Println("regexp.MustCompile err =");
		return;
	}

	// 取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1); // 返回一个二维数组
	//fmt.Println(joyIds);

	//
	for _, jValue := range joyUrls {
		joyUrl := jValue[1]; // 每个笑话的详情网址
		title, content, oneJoyErr := SpiderOneJoy(joyUrl);
		if oneJoyErr != nil {
			fmt.Println("oneJoyErr err =", oneJoyErr);
			continue;
		}

		//// 结果展示
		//fmt.Printf("title = %s\n", title);
		//fmt.Printf("content = %s\n", content);

		// 写入暂存数组
		fileTitles = append(fileTitles, title);
		fileContents = append(fileContents, content);
	}

	// 3）把内容写进文件
	saveJoys(i, fileTitles, fileContents);

	//
	page <- i;
}

func SpiderOneJoy(url string) (title, content string, err error) {
	result, httpGetErr := HttpGet(url);
	if httpGetErr != nil {
		err = httpGetErr;
		return;
	}

	// 标题<h1>标题</h1> 只取1个
	re := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`);
	if re == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err");
		return;
	}

	// 取标题
	tempTitle := re.FindAllStringSubmatch(result, 1);
	for _, data := range tempTitle {
		title = trimStr(data[1]);
		break;
	}

	// 内容<div class="content-txt pt10">内容<a id="prev" 只取1个
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`);
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err");
		return;
	}

	// 取内容
	tempContent := re2.FindAllStringSubmatch(result, 1);
	for _, data := range tempContent {
		content = trimStr(data[1]);
		break;
	}

	//
	return;
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
	result = "";
	for {
		n, respReadErr := resp.Body.Read(buf);
		if n == 0 { // 读取结束或者出问题
			if respReadErr == io.EOF {
				fmt.Println("GET读取结束");
			} else {
				fmt.Println("respReadErr =", respReadErr);
			}
			break;
		}

		//
		result += string(buf[:n]);
	}

	return result, nil;
}

func trimStr(str string) string {
	str = strings.Replace(str, "\t", "", -1);
	str = strings.Replace(str, "\n", "", -1);
	str = strings.Replace(str, "\r", "", -1);
	str = strings.Replace(str, "<br />", "", -1);
	str = strings.Replace(str, "&nbsp;", "", -1);

	//
	return str;
}

func saveJoys(page int, fileTitles, fileContents []string) {
	f, fileErr := os.Create("./iofiles/joy-" + strconv.Itoa(page) + ".txt");
	if fileErr != nil {
		fmt.Println("or.Create err =", fileErr);
		return;
	}
	defer f.Close();

	//
	length := len(fileTitles);
	for i := 0; i < length; i++ {
		// 正式写入
		_, fileWriteErr := f.WriteString(fileTitles[i] + "\r\n" + fileContents[i] + "\r\n\r\n");
		if fileWriteErr != nil {
			fmt.Println("WriteString err =", fileWriteErr);
			continue;
		}
	}
}
