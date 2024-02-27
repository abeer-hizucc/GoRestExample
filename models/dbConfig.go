package models

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
	TimeZone string
}
