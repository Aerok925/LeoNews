package configs

type DB struct {
	Host    string `env:"HOST,required"`
	Port    int    `env:"PORT,required"`
	DB      string `env:"DBNAME,required"`
	User    string `env:"USER,required"`
	Pass    string `env:"PASS,required"`
	SSL     string `env:"SSL_MODE,required"`
	MaxConn int    `env:"MAX_CONN,required"`
}
