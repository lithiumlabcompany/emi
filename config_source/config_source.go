package config_source

type ConfigSource interface {
	Get(string) (string, error)
}
