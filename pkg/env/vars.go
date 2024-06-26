package env

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Vars struct {
	//
	// ENV
	//

	ASYNQConcurrency int `mapstructure:"ASYNQ_CONCURRENCY"`

	HTTPServerPort string `mapstructure:"HTTP_SERVER_PORT"`

	// PostgresUser     string `mapstructure:"POSTGRES_USER"`
	// PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	// PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	// PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	// PostgresDB       string `mapstructure:"POSTGRES_DB"`
	// PostgresSchema   string `mapstructure:"POSTGRES_SCHEMA"`
	// PostgresTimezone string `mapstructure:"POSTGRES_TIMEZONE"`

	RedisDB       int    `mapstructure:"REDIS_DB"`
	RedisDBAsynq  int    `mapstructure:"REDIS_DB_ASYNQ"`
	RedisDBCache  int    `mapstructure:"REDIS_DB_CACHE"`
	RedisDBEvents int    `mapstructure:"REDIS_DB_EVENTS"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`

	WebSocketPort string `mapstructure:"WEB_SOCKET_PORT"`

	// Flags

	InitAll     bool `mapstructure:"init-all"`
	InitWorkers bool `mapstructure:"init-workers"`
}

func init() {
	flag.Bool("init-all", false, "initialize everything")
	flag.Bool("init-workers", false, "initialize workers")
}

func Load() (*Vars, error) {
	vars := &Vars{}

	// if err := BindEnv(vars); err != nil {
	//   return nil, err
	// }

	// env

	viper.AutomaticEnv()

	// flags

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	// unmarshal

	if err := viper.Unmarshal(vars); err != nil {
		return nil, err
	}

	// success

	return vars, nil
}

// func BindEnv(i interface{}) error {
//   r := reflect.TypeOf(i)

//   for r.Kind() == reflect.Ptr {
//     r = r.Elem()
//   }

//   for i := 0; i < r.NumField(); i++ {
//     env := r.Field(i).Tag.Get("mapstructure")
//     if err := viper.BindEnv(env); err != nil {
//       return err
//     }
//   }

//   return viper.Unmarshal(i)
// }

func Setup() {
}
