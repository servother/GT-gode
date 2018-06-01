package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"io"
)

func HelpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error2!: ", err)
	}
	fmt.Println(string(body))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func createFile(filename, wstring string, f *os.File, err1 error) {
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	CheckErr(err1)
	n, err1 := io.WriteString(f, wstring) //写入文件(字符串)
	CheckErr(err1)
	fmt.Printf("写入 %d 个字节n", n)
}


