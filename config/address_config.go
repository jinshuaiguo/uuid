package config

type AddressConfig struct {
	Prot int64 `mapstructure:"prot" json:"prot" yaml:"prot" description:"端口号"`
}
