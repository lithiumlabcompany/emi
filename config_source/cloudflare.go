package config_source

type CloudFlare struct {
}

func (c CloudFlare) Get(key string) (string, error) {
	return "", nil
}
