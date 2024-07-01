package config

type Server struct {
	Snowflake SnowflakeConfig `mapstructure:"snowflake" json:"snowflake" yaml:"snowflake" description:"雪花算法配置"`
}
