package config


type DataBaseConfig struct {
	Username string
	Password string
	IP       string
	Port     int
	Database string
}

type ObjectStorageConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Examplebucket   string
}

var (
	DBconf *DataBaseConfig
	OSconf *ObjectStorageConfig
)

