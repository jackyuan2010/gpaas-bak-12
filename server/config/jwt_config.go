package config

type JWTConfig struct {
	SecretKey  string `mapstructure:"scretkey" json:"scretkey" yaml:"scretkey"`
	ExpireTime int64  `mapstructure:"expiretime" json:"expiretime" yaml:"expiretime"`
	BufferTime int64  `mapstructure:"buffertime" json:"buffertime" yaml:"buffertime"`
	Issuer     string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}