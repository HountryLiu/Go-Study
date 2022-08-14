package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//创建一个结构体
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 创建一个原始的数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	// chessMap[4][6] = 1
	row := 0
	col := 0
	for _, i := range chessMap {
		row++
		for _, _ = range i {
			col++
			// fmt.Printf("%d\t", j)
		}
		// fmt.Println()
	}

	// 创建稀疏数组的直节点
	var valnode []ValNode
	// 标准的稀疏数首行是 代表着 行和列 以及二维数组的默认值
	node := ValNode{
		row: row,
		col: col / row,
		val: 0,
	}
	valnode = append(valnode, node)
	for a, i := range chessMap {

		for b, j := range i {
			if j != 0 {
				node := ValNode{
					row: a,
					col: b,
					val: j,
				}
				valnode = append(valnode, node)
			}
		}
	}
	f, err := os.OpenFile("valnode.txt", os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(f)
	for _, val := range valnode {
		str := fmt.Sprintf("%d %d %d\n", val.row, val.col, val.val)
		writer.WriteString(str)
		writer.Flush()
	}
	// 稀疏数组变成二位数组
	var str1 []string
	// 遍历文件的每一行
	f, err = os.Open("valnode.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		// fmt.Println(str)
		if err == nil {
			str1 = append(str1, str)
		} else {
			break
		}
	}
	// 将稀疏数组存在字符串切片中
	// 创建 ValNode 切片
	var valnode1 []ValNode
	var str2 []string
	for _, i := range str1 {
		str2 = strings.Split(i, " ")
		row1, _ := strconv.Atoi(str2[0])
		col1, _ := strconv.Atoi(str2[1])
		s2 := strings.Replace(str2[2], "\n", "", -1)
		val1, _ := strconv.Atoi(s2)
		// fmt.Println(val1)
		node1 := ValNode{
			row: row1,
			col: col1,
			val: val1,
		}
		valnode1 = append(valnode1, node1)
	}
	// 将稀疏数组还原成二维数组
	row = valnode1[0].row
	col = valnode1[0].col
	var Sparse = make([][]int, row)
	for i, _ := range Sparse {
		Sparse[i] = make([]int, col)
	}
	for i, str3 := range valnode1 {
		if i != 0 {
			Sparse[str3.row][str3.col] = str3.val
		}
	}
	// 已经进行了稀疏数组替换成二维数组
	for _, x := range Sparse {
		for _, y := range x {
			fmt.Printf("%d\t", y)
		}
		fmt.Println()
	}

}
