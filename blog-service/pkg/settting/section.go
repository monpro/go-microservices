package settting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type AppSettingS struct {
	DefaultContextTimeout time.Duration
	DefaultPageSize       int
	MaxPageSize           int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

func (s *Setting) ReadSection(key string, value interface{}) error {
	err := s.vp.UnmarshalKey(key, value)
	if err != nil {
		return err
	}
	return nil
}
