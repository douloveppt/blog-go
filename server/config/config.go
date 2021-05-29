package config

import "time"

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    //jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` //过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    //缓冲时间
}

type Zap struct {
	Level         string        `json:"level" yaml:"level" mapstructure:"level"`                           //日志级别
	Format        string        `json:"format" yaml:"format" mapstructure:"format"`                        //日志格式
	Prefix        string        `json:"prefix" yaml:"prefix" mapstructure:"prefix"`                        //日志前缀
	Director      string        `json:"director" yaml:"director" mapstructure:"director"`                  //日志目录
	LinkName      string        `json:"linkName" yaml:"link-name" mapstructure:"link-name"`                //软连接名称
	ShowLine      bool          `json:"showLine" yaml:"showLine" mapstructure:"show-line"`                 //显示行
	EncodeLevel   string        `json:"encodeLevel" yaml:"encode-level" mapstructure:"encode-level"`       //编码级
	StacktraceKey string        `json:"stacktraceKey" yaml:"stacktrace-key" mapstructure:"stacktrace-key"` //栈名
	LogInConsole  bool          `json:"logInConsole" yaml:"log-in-console" mapstructure:"log-in-console"`  //是否输出控制台
	MaxAge        time.Duration `json:"maxAge" yaml:"max-age" mapstructure:"max-age"`                      //保存时间
	RotationTime  time.Duration `json:"rotationTime" yaml:"rotation-time" mapstructure:"rotation-time"`    //切割时间
}

type Redis struct {
	DB       int    `json:"db" yaml:"db" mapstructure:"db"`                   //数据库
	Addr     string `json:"addr" yaml:"addr" mapstructure:"addr"`             //服务地址：端口
	Password string `json:"password" yaml:"password" mapstructure:"password"` //密码
}

type Captcha struct {
	KeyLong   int `mapstructure:"key_long" yaml:"key_long" json:"key_long"`       //验证码长度
	ImgWidth  int `mapstructure:"img_width" yaml:"img_width" json:"img_width"`    //图片宽度
	ImgHeight int `mapstructure:"img_height" yaml:"img_height" json:"img_height"` //图片高度
}

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             //服务地址：端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       //配置
	Dbname       string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`                       //数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 //用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 //密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` //
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` //
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  //
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     //
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` //本地文件路径
}
