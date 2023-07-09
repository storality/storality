package config

type Config struct {
	Port int
	Headless bool
	Driver string
	Connection string
	CMD string
}

func Load(port int, headless bool, driver string, connection string, cmd string) *Config {
	return &Config{
		Port: port,
		Headless: headless,
		Driver: driver,
		Connection: connection,
		CMD: cmd,
	}
}