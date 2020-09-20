package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//定义2个接口 一个读接口  一个写接口
type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

//定义2个类
type ReadFromFile struct {
	path string
}
type WriteToInfluxDB struct {
	influxDBDsn string
}

//实现2个接口
func (r *ReadFromFile) Read(rc chan []byte) {
	//读取模块
	//打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	f.Seek(0, 2)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}

}

func (w *WriteToInfluxDB) Write(wc chan string) {
	//写入模块
	for v := range wc {
		fmt.Println(v)
	}
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Reader
	write Writer
}

func (l *LogProcess) Process() {
	//解析模块

	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

//执行入口
func main() {
	r := &ReadFromFile{
		path: "access.log",
	}
	w := &WriteToInfluxDB{
		influxDBDsn: "username&password",
	}
	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		write: w,
		read:  r,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)
	time.Sleep(1 * time.Second)
}
