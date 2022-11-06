package m3u8Download

/**
 * @title start
 * @author CH00SE1
 * @date 2022-10-24 17:04:30
 */
import (
	"flag"
	"fmt"
	"github.com/CH00SE1/m3u8Download/download"
	"log"
	"os"
)

type VideoInfo struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
	Output   string `json:"output"`
	ChanSize int    `json:"chanSize"`
}

func (vl VideoInfo) DownloadM3u8() {
	flag.Parse()
	defer func() {
		r := recover()
		fmt.Println("[error]", r)
		os.Exit(-1)
	}()
	if vl.Url == "" {
		log.Panic("url is nil")
	}
	if vl.Output == "" {
		log.Panic("output is nil")
	}
	if vl.ChanSize <= 0 {
		log.Panic("chansize is nil")
	}
	if vl.FileName == "" {
		log.Panic("filename is nil")
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
