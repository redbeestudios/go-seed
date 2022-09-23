package pkg

type ServerConfig struct {
	Address      string `json:"address"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	IdleTimeout  int    `json:"idle_timeout"`
}

type RestClientConfig struct {
	BaseUrl               string `json:"base_url"`
	TimeoutMilliseconds   int    `json:"timeout_milliseconds"`
	RetryCount            int    `json:"retry_count"`
	RetryWaitMilliseconds int    `json:"retry_wait_milliseconds"`
}

type RedisConfig struct {
	BaseUrl string `json:"base_url"`
}
