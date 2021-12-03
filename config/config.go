package config

var (
	Values values
)

type values struct {
	ServerAddress string
	LoggerLevel   string
	DB            db
}

type db struct {
	Host     string
	Port     uint
	User     string
	Password string
	DBName   string
}

func NewValues() *values {
	database := db{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "gvs1!",
		DBName:   "Product",
	}
	return &values{
		ServerAddress: ":8083",
		LoggerLevel:   "debug",
		DB:            database,
	}
}
func Load() {
	Values = *NewValues()
}
