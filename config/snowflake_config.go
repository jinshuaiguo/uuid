package config

type SnowflakeConfig struct {
	Workerid int64 `mapstructure:"workerid" json:"workerid" yaml:"workerid" description:"机器ID"`
}
