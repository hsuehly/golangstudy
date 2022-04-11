package gofile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Files() {
	// 打开文件
	//fileName := "/Users/hsuehly/code/goCode/src/gofile/a.txt"
	// readFile(fileName)
	//readFileBuff(fileName)
	//writeFile(fileName)
	//appendFile(fileName)
	//readAndWriteFile(fileName)
	//writeB(fileName)
	// 拷贝文件
	//copyFile(fileName)
	statString()
}

// 读取文件
func readFile(fileName string) {
	//file 叫file指针，对象，文件句柄
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file %v \n", file)
	fmt.Println(file.Name(), "filename")
	// 当函数退出时要及时关闭文件 关闭文件
	defer file.Close() // 及时关闭，否则会有内存泄露
	// 读取文件，代缓冲区
	reader := bufio.NewReader(file)
	// 循环读取文件的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.eof 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取完成。。")

}

// 带缓存读取文件
func readFileBuff(fileName string) {

	//读取文件 ioutil.readfile 一次性将文件读取完成，适合文件较小的
	// 返回的byte切片
	content, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("error", error)
	}
	// 转为string
	fmt.Printf("%v", string(content))

}

// 写文件
func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hello,hsueh\r\n") // \r \n 表示换行

	}
	// 因为 writer是带缓存的，因此在调用writerString 的方法时
	// 内容先写入到缓存的，所以需要调用Flush 方法，将缓存数据写入到文件中 否则文件中会没有数据
	writer.Flush()

}

// 追加文件内容
func appendFile(fileName string) {

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("你好啊\r\n") // \r \n 表示换行

	}
	// 因为 writer是带缓存的，因此在调用writerString 的方法时
	// 内容先写入到缓存的，所以需要调用Flush 方法，将缓存数据写入到文件中 否则文件中会没有数据
	writer.Flush()
}

// 读写文件
func readAndWriteFile(fileName string) {

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("你好啊,北京\r\n") // \r \n 表示换行

	}
	// 因为 writer是带缓存的，因此在调用writerString 的方法时
	// 内容先写入到缓存的，所以需要调用Flush 方法，将缓存数据写入到文件中 否则文件中会没有数据
	writer.Flush()

}

// 读取a 中的文件，并写入b中
func writeB(fileName string) {
	filepath := "/Users/hsuehly/code/goCode/src/gofile/b.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = ioutil.WriteFile(filepath, data, 0777)
	if err != nil {
		fmt.Println("err", err)
	}
}

// 拷贝文件
func copy(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	// 打开dstdilename
	dstFile, err := os.OpenFile(dstFileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// 通过dstfile 获取到writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
func copyFile(fileName string) {
	copyFile := "/Users/hsuehly/code/goCode/src/gofile/a.txt"
	_, err := copy(fileName, copyFile)
	if err != nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝失败")
	}
}

type CharCount struct {
	char  int // 英文个数
	num   int // 数字
	space int // 空格
	other int // 其他
}

// 统计不同的字符个数
func statString() {
	fileName := "/Users/hsuehly/code/goCode/src/gofile/d.txt"
	var count CharCount
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()
	// 创建一个reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.char++
			case v == ' ' || v == '\t':
				count.space++
			case v >= '0' && v <= '9':
				count.num++
			default:
				count.other++

			}
		}
	}
	fmt.Printf("字符串的个数为%v,数字的个数为%v,空格的个数为%v,其他为%v", count.char, count.num, count.space, count.other)
}
