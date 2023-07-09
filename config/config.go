package config

type Config struct {
	Port int
	Headless bool
	DataDir string
	CMD string
}

func Load(port int, headless bool, dataDir string, cmd string) *Config {
	return &Config{
		Port: port,
		Headless: headless,
		DataDir: dataDir,
		CMD: cmd,
	}
}