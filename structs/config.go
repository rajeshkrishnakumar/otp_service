package structs

type Config struct {
	SmsDomain  string `mapstructure:"SmsDomain"`
	SmsTokten  string `mapstructure:"SmsTokten"`
	Env        string `mapstructure:"environment"`
	NiyamUrl   string `mapstructure:"niyamUrl"`
	REDIS_HOST string `mapstructure:"redis_host"`
	REDIS_PORT string `mapstructure:"redis_port"`
	REDIS_DB   int    `mapstructure:"redis_db"`
}
