package utils

import (
    "bufio"
    "fmt"
    "os"
    "time"
)


func Mzlog(s string) {
    filePath := "./mz.log"
    file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("MMM 文件打开失败", err)
    }
    //及时关闭file句柄
    defer file.Close()
    //写入文件时，使用带缓存的 *Writer
    write := bufio.NewWriter(file)
    write.WriteString(fmt.Sprintf("[%v] %v \n", time.Now().Unix(), s))
    //Flush将缓存的文件真正写入到文件中
    write.Flush()
}