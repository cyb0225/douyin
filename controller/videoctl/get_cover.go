package videoctl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 截图获取封面
func GetCover(video string, playUrl string) (string, error) {
	coverName := creatCoverFileName(video)
	reader := ReadFrameAsJpeg(playUrl, 1)
	coverUrl, err := OS.PutCoverObject(coverName, &reader)
	if err != nil {
		return "", err
	}

	return coverUrl, nil
}

// 参数是视频地址， 视频要截图的帧数
func ReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

// 根据video生成JPEG文件名
func creatCoverFileName(video string) (cover string) {
	names := strings.Split(video, ".")
	cover = names[0] + ".jpeg"
	return
}
