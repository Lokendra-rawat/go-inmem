package util

import (
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	GoEnv                       string `mapstructure:"GO_ENV"`
	GrpcPort                    string `mapstructure:"GRPC_PORT"`
	HttpPort                    string `mapstructure:"HTTP_PORT"`
	ConsumerHttpPort            string `mapstructure:"CONSUMER_HTTP_PORT"`
	DBConsumerGroup             string `mapstructure:"DB_CONSUMER_GROUP"`
	RankerStoreUpdatesTopic     string `mapstructure:"RANKER_STORE_UPDATES_TOPIC"`
	DatabaseUrlRr3              string `mapstructure:"DATABASE_URL_RR3"`
	DatabaseUrl                 string `mapstructure:"DATABASE_URL"`
	RecoModelBucket             string `mapstructure:"RECO_MODEL_BUCKET"`
	AwsAccessKeyId              string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey          string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion                   string `mapstructure:"AWS_REGION"`
	KafkaDSBrokers              string `mapstructure:"KAFKA_SIMZI_SERVER"`
	KafkaDSUser                 string `mapstructure:"KAFKA_SIMZI_USER"`
	KafkaDSPassword             string `mapstructure:"KAFKA_SIMZI_PASSWORD"`
	FeatureStoreConsumerAddress string `mapstructure:"FEATURE_STORE_CONSUMER_ADDRESS"`
}

type configType struct {
	Data     *Config
	Mu       sync.Mutex
	IsLoaded bool
}

var config configType

func LoadConfig() {
	config.Mu.Lock()
	defer config.Mu.Unlock()
	if !config.IsLoaded {
		viper.AutomaticEnv()

		dir, _ := os.Getwd()
		log.Info().Str("Component", "util load config").Msgf("cwd %s", dir)
		if !strings.HasSuffix(os.Args[0], ".test") {
			log.Info().Msg("normal run")
			viper.AddConfigPath(".")
			viper.SetConfigName("app")
			viper.SetConfigType("env")
		} else {
			log.Info().Msg("run under go test")
			viper.AddConfigPath("..")
			viper.SetConfigName("app.test")
			viper.SetConfigType("env")
		}

		err := viper.ReadInConfig()
		if err != nil {
			log.Panic().Str("Component", "util load config").Err(err).Msg("Failed to load config")
		}

		config.Data = &Config{}
		err = viper.Unmarshal(&config.Data)
		if err != nil {
			log.Panic().Str("Component", "util load config").Err(err).Msg("Failed to load config")
		}

		config.IsLoaded = true
	}
}

func GetConfig() Config {
	LoadConfig()
	return *config.Data
}

func OverrideEnv(propName, value string) {
	LoadConfig()
	v := reflect.ValueOf(config.Data).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("mapstructure")

		if tag == propName {
			fieldValue := v.Field(i)
			if fieldValue.IsValid() && fieldValue.CanSet() {
				nv := reflect.ValueOf(value)
				if fieldValue.Type() == nv.Type() {
					fieldValue.Set(nv)
					return
				} else {
					panic("type mismatch: expected " + fieldValue.Type().Name() + " but got " + nv.Type().Name())
				}
			} else {
				panic("cannot set the field")
			}
		}
	}

	panic("property name not found")
}
