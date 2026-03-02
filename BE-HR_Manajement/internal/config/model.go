package config

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Tz       string
}

type JWT struct {
	Secret string
	Exp    int
}
