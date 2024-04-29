package postgres

import (
	"events/pkg/env"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	driver "gorm.io/driver/postgres"
)

type Client struct {
	*gorm.DB

	Options *Options
}

type Options struct {
	PostgresUsername string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
}

var (
	instance *Client
)

func Initialize(env *env.Vars) error {
	options := &Options{}

	err := copier.Copy(options, env)

	if err != nil {
		return err
	}

	path := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v  sslmode=disable TimeZone=%v",
		env.PostgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDB,
		env.PostgresPort,
		env.PostgresTimezone,
	)

	prefix := fmt.Sprintf("%s.", env.PostgresSchema)

	clientDriver, err := gorm.Open(
		driver.Open(path),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: prefix,
			},
		},
	)

	if err != nil {
		return err
	}

	db, err := clientDriver.DB()

	if err != nil {
		return err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)

	instance = &Client{
		DB:      clientDriver,
		Options: options,
	}

	return nil
}

func Terminate() error {
	return nil
}
