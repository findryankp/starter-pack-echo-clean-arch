package configs

type AppConfig struct {
	DBUSERNAME string
	DBPASS     string
	DBHOST     string
	DBPORT     string
	DBNAME     string
	JWTKEY     string
}

func InitConfig() *AppConfig {
	return InitEnv()
}
