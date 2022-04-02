/*
 * @Author: Alexleslie
 * @Date: 2022-03-28 14:12:33
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-04-01 14:45:25
 * @FilePath: \src\下载器\main.go
 * @Description:
 */

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Pool struct {
	active   chan int
	currSize chan int
	wg       *sync.WaitGroup
}

type Link struct {
	contentLength int
	contentType   string
	filename      string
}

type F struct {
	name string
	data []byte
}

var goNums int = 6
var distance int = 1 << 24

func main() {
	var url string
	//url = "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farchive%2F2b003bd13a3e6a9b720fd5d27c367913ce966e0b.jpg&refer=http%3A%2F%2Fi0.hdslb.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1651040048&t=ef708f8b03e440ded61a9ac78f67d95c"
	//url = "http://idl.hbdlib.cn/book/00000000000000/pdfbook/018/017/156277.pdf"
	//url = "https://issuepcdn.baidupcs.com/issue/netdisk/yunguanjia/BaiduNetdisk_7.14.1.6.exe"
	//url = "https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe"
	fmt.Scan(&url)

	f, err := control(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	toFile(f.data, f.name)

}

func control(url string) (*F, error) {
	pool := &Pool{
		make(chan int, goNums),
		make(chan int, 1),
		&sync.WaitGroup{},
	}
	pool.currSize <- 1

	log.Println("Start connecting the remote server --->")
	link, err := inistalize(url)
	if err != nil {
		return nil, err
	}

	data := make([]byte, link.contentLength)
	bounds := getBounds(link.contentLength)

	log.Printf("Start downloading file: %s \n", link.filename)
	log.Printf("The numbers of goroutine: %d \n", goNums)

	Now := time.Now()
	go func() {
		pool.Add(1)
		runtime.LockOSThread()
		displayProcess(pool, link.contentLength)
	}()
	for len(bounds) > 0 {
		pool.Add(1)
		interval := <-bounds
		go goDownload(pool, data, url, interval[0], interval[1], bounds)
	}

	pool.Wait()

	spend := time.Now().Sub(Now)
	speed := link.contentLength / int(spend.Seconds())
	log.Printf("Avg Speed : %v KB/s, total times : %v", speed/1024, spend.String())

	return &F{data: data, name: link.filename}, nil
}

func goDownload(pool *Pool, data []byte, url string, start, end int, bounds chan []int) {
	defer pool.Done()
	temp, err := download(url, start, end)
	if err != nil {
		bounds <- []int{start, end}
		return
	}
	copy(data[start:end+1], temp)

	t := <-pool.currSize
	t = t + end - start
	pool.currSize <- t
}

/**
 * @description: 下载文件的函数
 * @param {string} url
 * @param {*} start
 * @param {int} end
 * @return {*}
 */
func download(url string, start, end int) ([]byte, error) {
	resq, _ := http.NewRequest("GET", url, nil)
	fileRange := fmt.Sprintf("bytes=%d-%d", start, end)
	browser := "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"

	resq.Header.Add("Range", fileRange)
	resq.Header.Add("user-agent", browser)

	client := &http.Client{}
	resp, err := client.Do(resq)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func displayProcess(p *Pool, size int) {
	defer p.Done()
	for {
		temp := <-p.currSize
		p.currSize <- temp
		percent := int((float64(temp) / float64(size)) * 100)
		graph := ">"
		for i := 0; i < percent/2; i++ {
			graph += ">"
		}

		fmt.Printf("\r[%-51s]%3d%%  %8d/%d", graph, percent, temp, size)
		if temp >= size {
			break
		}
		time.Sleep(time.Second / 5)
	}
	fmt.Println()
}

func toFile(data []byte, myfile string) {
	fd, _ := os.Create(myfile)
	fd.Write(data)
	fmt.Printf("download file successfully :%s", myfile)

}

func inistalize(url string) (*Link, error) {
	client := &http.Client{}
	hp, err := client.Head(url)
	if err != nil {
		return nil, err
	}

	respH := hp.Header
	length := respH["Content-Length"][0]
	size, _ := strconv.Atoi(length)
	fileType := respH["Content-Type"][0]

	link := &Link{size, fileType, getFilename(fileType)}
	return link, nil
}

func getFilename(fileType string) string {
	mime.AddExtensionType(".exe", "application/x-msdownload")
	mime.AddExtensionType(".exe", "binary/octet-stream")
	extend, _ := mime.ExtensionsByType(fileType)
	rand.Seed(time.Now().UnixMicro())
	myfile := strconv.Itoa(rand.Int()) + extend[len(extend)-1]
	return myfile
}

func getBounds(size int) chan []int {
	capity := size / distance
	bounds := make(chan []int, capity+1)
	for i := 0; i < size; {
		if i+distance >= size {
			bounds <- []int{i, size - 1}
			break
		}
		bounds <- []int{i, i + distance}
		i = i + distance
	}
	return bounds
}

func (p *Pool) Add(nums int) {
	for i := 0; i < nums; i++ {
		p.active <- 1
		p.wg.Add(1)
	}
}

func (p *Pool) Done() {
	<-p.active
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
