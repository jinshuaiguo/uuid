package config

type Server struct {
	Snowflake SnowflakeConfig `mapstructure:"snowflake" json:"snowflake" yaml:"snowflake" description:"雪花算法配置"`
	Address   AddressConfig   `mapstructure:"address" json:"address" yaml:"address" description:"地址配置"`
}
