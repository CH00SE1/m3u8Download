package m3u8Download

import (
	"flag"
	"fmt"
	"github.com/CH00SE1/m3u8Download/download"
	"log"
	"os"
)

/**
 * @title start
 * @author CH00SE1
 * @date 2022-10-24 17:04:30
 */
type VideoInfo struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
	Output   string `json:"output"`
	ChanSize int    `json:"chanSize"`
}

func DownloadM3u8(vl *VideoInfo) {
	flag.Parse()
	defer func() {
		r := recover()
		fmt.Println("[error]", r)
		os.Exit(-1)
	}()
	if vl.Url == "" {
		panicParameter("url is nil")
	}
	if vl.Output == "" {
		panicParameter("output is nil")
	}
	if vl.ChanSize <= 0 {
		log.Panic(`parameter 'c' must be greater than 0`)
	}
	if vl.FileName == "" {
		panicParameter("filename is nil")
	}
	downloader, err := download.NewTask(vl.Output, vl.Url, vl.FileName)
	if err != nil {
		log.Panic(err)
	}
	if err := downloader.Start(vl.ChanSize); err != nil {
		log.Panic(err)
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	log.Panic("parameter '" + name + "' is required")
}
