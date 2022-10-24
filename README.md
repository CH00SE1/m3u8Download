# [作者](https://github.com/CH00SE1/)  说明

### 包引入

```yaml
import (
"github.com/CH00SE1/m3u8Download"
)
```

### 使用案例

```go
func main() {

// 下载启动
download_m3u8_video.DownloadM3u8(&download_m3u8_video.VideoInfo{

FileName: `乌龙院`,

Url:      `https://mgtv.sd-play.com/20220928/ikNW1Fld/1200kb/hls/index.m3u8`,

ChanSize: 100,

Output: "D:\\m3u8_video\\go\\" + time.Now().Format("2006-01-02") + "\\",

})

}

```