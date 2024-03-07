package config

import (
	"github.com/qiniu/go-sdk/v7/auth"
)

type QiNiu struct {
	AccessKey   string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey   string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	Bucket      string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ExpiresTime int    `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
}

type QiNiuClient struct {
	Mac *auth.Credentials
}
