package pkg

import (
	"fmt"
	"os"

	_"github.com/bytedance/sonic/option"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload" // add underscore
)

func InitMySql() (*sqlx.DB, error) {
	
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbname := os.Getenv("DB_NAME")

	
	// uname:pass@tcp(host:port)/dbname
	// connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	
	var options []any
	options = append(options, os.Getenv("DB_USER"))
	options = append(options, os.Getenv("DB_PASS"))
	options = append(options, os.Getenv("DB_HOST"))
	options = append(options, os.Getenv("DB_PORT"))
	options = append(options, os.Getenv("DB_NAME"))
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",options...)
	
	return sqlx.Connect("mysql", connectionStr)
}
