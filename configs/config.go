package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	Host     string
	User     string
	Password string
	Port     string
	DBNAME   string
}

func ImportSetting() setting {
	var result setting
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	return setting{}
	// }
	// result.Host = os.Getenv("poshost")
	// result.User = os.Getenv("posuser")
	// result.Password = os.Getenv("pospw")
	// result.Port = os.Getenv("posport")
	// result.DBNAME = os.Getenv("dbname")

	result.Host = "aws-0-ap-southeast-1.pooler.supabase.com"
	result.User = "postgres.wgemmubheztwipfiszza"
	result.Password = "TfSBWEL8an6bJ7zQ"
	result.Port = "5432"
	result.DBNAME = "postgres"
	return result
}

func ConnectDB(s setting) (*gorm.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBNAME)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "newsapp.",
		},
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
