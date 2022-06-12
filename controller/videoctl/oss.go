// 使用对象存储技术保存视频和封面数据，并对用户提供访问的url

package videoctl

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"

	"github.com/2103561941/douyin/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ObjectStorage struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Examplebucket   string
	Url             string
	Client          *oss.Client
	Bucket          *oss.Bucket
}

// 配置个人的oss仓库信息
var (
	OS *ObjectStorage
)

// 初始化 oss 对象存储库
func InitOss() error {
	OS = &ObjectStorage{
		Endpoint:        config.OSconf.Endpoint,
		AccessKeyId:     config.OSconf.AccessKeyId,
		AccessKeySecret: config.OSconf.AccessKeySecret,
		Examplebucket:   config.OSconf.Examplebucket,
	}

	if err := OS.CreatStorySpace(); err != nil {
		return err
	}
	return nil
}

// 创建oss存储空间
func (obj *ObjectStorage) CreatStorySpace() error {
	// 创建OSSClient实例。
	client, err := oss.New(obj.Endpoint, obj.AccessKeyId, obj.AccessKeySecret)
	if err != nil {
		return err
	}
	// 创建名为examplebucket的存储空间，并设置存储类型为标准存储类型、读写权限ACL为公共读oss.ACLPublicRead、数据容灾类型为同城冗余存储oss.RedundancyZRS。
	err = client.CreateBucket(obj.Examplebucket, oss.StorageClass(oss.StorageStandard), oss.ACL(oss.ACLPublicRead), oss.RedundancyType(oss.RedundancyZRS))
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(obj.Examplebucket)
	if err != nil {
		return err
	}

	obj.Client = client
	obj.Bucket = bucket

	obj.Url = "https://" + obj.Bucket.BucketName + "." + obj.Endpoint + "/"

	return nil
}

// 传输存储对象，保存到oss的bucket中, 保存视频
func (obj *ObjectStorage) PutVideoObject(objectKey string, data *multipart.FileHeader) (string, error) {
	src, err := data.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	savePath := "video" + "/" + objectKey

	err = obj.Bucket.PutObject(savePath, src)
	if err != nil {
		return "", err
	}

	url := obj.Url + savePath

	return url, nil
}

// 传输存储对象，保存到oss的bucket中, 保存封面
func (obj *ObjectStorage) PutCoverObject(objectKey string, data *io.Reader) (string, error) {

	savePath := "cover" + "/" + objectKey

	// 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
	reader := bytes.NewReader([]byte("hello world"))
	if reader == nil {
		return "", errors.New("reader := bytes.NewReader(buffer) failed")
	}

	err := obj.Bucket.PutObject(savePath, *data)
	if err != nil {
		return "", err
	}

	url := obj.Url + savePath

	return url, nil
}
