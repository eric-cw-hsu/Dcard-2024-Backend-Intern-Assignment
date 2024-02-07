package databases

type DatabaseConfigs struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type BaseDatabase interface {
	Connect()
	Close()
	GetConnectionString() string
}
