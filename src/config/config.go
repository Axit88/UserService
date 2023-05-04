package config

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/Axit88/UserService/src/constants"

	"gopkg.in/yaml.v3"
)

type EnvConfigs struct {
	Configs map[string]*Config
}

type Config struct {
	UserServiceUrl *UserServiceUrl `yaml:"user_service_url"`
}

type UserServiceUrl struct {
	GrpcUrl string `yaml:"gurl"`
	RestUrl string `yaml:"rurl"`
}

func getParsedYamlConfig() (*EnvConfigs, error) {
	envConfigs := new(EnvConfigs)

	lambdaPath := "/Users/axit/Desktop/UserApi/src/config"
	filePath := path.Join(lambdaPath, "config.yaml")
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	yamlDataString := os.ExpandEnv(string(yamlData))
	err = yaml.Unmarshal([]byte(yamlDataString), envConfigs)
	if err != nil {
		return nil, err
	}
	return envConfigs, nil
}

func GetCurrentEnv() (string, error) {
	env := os.Getenv(constants.APP_ENV)
	if env == "" {
		return "", nil
	}

	if strings.HasPrefix(env, "prod") {
		env = "prod"
	}

	if env != "prod" && env != "test" {
		env = "dev"
	}
	return env, nil
}

func NewConfig() (*Config, error) {
	env, err := GetCurrentEnv()
	if err != nil {
		return nil, err
	}

	envConfigs, err := getParsedYamlConfig()
	if err != nil {
		return nil, err
	}

	conf := envConfigs.Configs[env]
	if conf == nil {
		return nil, errors.New("error in config file")
	}
	return conf, nil
}
