package pkg

type Config struct {
	Redis Redis
}

type Redis struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
	db   int    `yaml:"db"`
}

type Database struct {
	host     string `yaml:"host"`
	user     string `yaml:"user"`
	db       string `yaml:"db"`
	password string `yaml:"password"`
	port     int    `yaml:"port"`
}

type WebServer struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

func load_config() {
	//io.Copy("./dev.yaml", "./env.yaml")
	//io.copy
}
