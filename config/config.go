package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	// RotateLogs RotateLogs `mapstructure:"rotateLogs" json:"rotateLogs" yaml:"rotateLogs"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	// Captcha     Captcha     `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"` // 跨域配置
	// Storage     Storage     `mapstructure:"storage" json:"storage" yaml:"storage"`
	// Minio       Minio       `mapstructure:"minio" json:"minio" yaml:"minio"`
	// MeiliSearch MeiliSearch `mapstructure:"meiliSearch" json:"meiliSearch" yaml:"meiliSearch"`
	// QuestDB     QuestDB     `mapstructure:"questdb" json:"questdb" yaml:"questdb"`
	// Mqtt        Mqtt        `mapstructure:"mqtt" json:"mqtt" yaml:"mqtt"`
	Wechat WechatConfig `mapstructure:"wechat" json:"wechat" yaml:"wechat"`
}
