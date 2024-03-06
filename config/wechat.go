package config

type WechatConfig struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 是否开启日志栈
}
