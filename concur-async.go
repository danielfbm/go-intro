package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func syncRequests(print bool) {

	fmt.Println("Starting SYNC")
	start := time.Now()

	baidu, _ := http.Get("https://baidu.com")

	bing, _ := http.Get("https://bing.com")

	fmt.Println("ended: ", time.Now().Sub(start).Seconds(), "s")

	if print {
		fmt.Println("Baidu data header  ---- \n", baidu.Header)
		fmt.Println("Bing data length ---- \n", bing.Header)
	}

}

func asyncRequest(print bool) {
	var baiduRes, bingRes *http.Response

	var group sync.WaitGroup
	fmt.Println("Starting ASYNC")
	start := time.Now()

	group.Add(1)
	go func() {
		baiduRes, _ = http.Get("https://baidu.com")
		group.Done()
	}()

	group.Add(1)
	go func() {
		bingRes, _ = http.Get("https://bing.com")
		group.Done()
	}()
	group.Wait()

	fmt.Println("ended: ", time.Now().Sub(start).Seconds(), "s")

	if print {
		fmt.Println("Baidu data header  ---- \n", baiduRes.Header)
		fmt.Println("Bing data length ---- \n", bingRes.Header)
		// baiduData, _ := ioutil.ReadAll(baiduRes.Body)
		// bingData, _ := ioutil.ReadAll(bingRes.Body)
		//
		// fmt.Println(string(baiduData))
		//
		// fmt.Println(string(bingData))
	}
}

func main() {
	print := false

	if len(os.Args) > 1 {
		asyncRequest(print)
	} else {
		syncRequests(print)
	}

}
