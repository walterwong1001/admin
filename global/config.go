package global

import (
	"github.com/spf13/viper"
	"log"
)

const (
	DEFAULT_CONFIG_NAME = "default"
	DEFAULT_CONFIG_TYPE = "toml"
	DEFAULT_CONFIG_PATH = "../deploy"
	ENV_PREFIX          = "CONFIG"
)

var CONFIG = &config{}

func init() {
	// Working with Environment Variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix(ENV_PREFIX)

	fileName, fileType, filePath := getValue("FILE_NAME", DEFAULT_CONFIG_NAME), getValue("FILE_TYPE", DEFAULT_CONFIG_TYPE), getValue("FILE_PATH", DEFAULT_CONFIG_PATH)

	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error: config file %s.%s not found in path %s", fileName, fileType, filePath)
	}

	if err := viper.Unmarshal(CONFIG); err != nil {
		log.Fatalf("Fatal error: failed to unmarshal config: %v", err)
	}
}

// getValue retrieves the value from the environment variable or returns the default value if not set.
func getValue(envKey, defaultValue string) string {
	value := viper.GetString(envKey)
	if value == "" {
		return defaultValue
	}
	return value
}

type (
	config struct {
		Application string
		Locale      string
		Server      *server
		MySQL       *mysql
		Snowflake   *snowflake
		Jwt         *jwt
	}

	server struct {
		Port uint16
	}

	mysql struct {
		Username string
		Password string
		Host     string
		Port     uint16
		Database string
	}
	snowflake struct {
		Register string
	}

	jwt struct {
		SecretKey string `mapstructure:"secret_key"`
		Issuer    string
		Days      int // 有效天数
	}
)
