package config

type JWTConfig struct {
	SecretKey  string `mapstructure:"scretkey" json:"scretkey" yaml:"scretkey"`
	ExpiresTime int64  `mapstructure:"expirestime" json:"expirestime" yaml:"expirestime"`
	BufferTime int64  `mapstructure:"buffertime" json:"buffertime" yaml:"buffertime"`
	Issuer     string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}