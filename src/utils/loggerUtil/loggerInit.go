package loggerUtil

import (
	"errors"

	"github.com/Axit88/UserService/src/config"
	"github.com/MindTickle/mt-go-logger/logger"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"

	// "io/ioutil"
	"os"
	"path"

	// "runtime"
	"strings"
)

var LambdaRequestId string
var StepFunctionExecutionId string
var SyncJobId string

type LoggerConfigs struct {
	Configs map[string]*LoggerConfig
}

type LoggerConfig struct {
	Level       string   `yaml:"level"`
	OutputPaths []string `yaml:"output_paths"`
	Appends     []string `yaml:"appends"`
}

type LogDetails struct {
	LambdaRequestId         string      `json:"lambda_request_id"`
	StepFunctionExecutionId string      `json:"step_function_execution_id"`
	SyncJobId               string      `json:"sync_job_id"`
	SyncEntityId            string      `json:"sync_entity_id"`
	LogMessage              string      `json:"log_message"`
	Other                   interface{} `json:"other"`
}

func InitLogger() (*logger.LoggerImpl, error) {
	env, err := config.GetCurrentEnv()
	if err != nil {
		return nil, err
	}

	loggerConfig, err := getParsedLoggerYamlConfig()
	if err != nil {
		return nil, err
	}

	conf := loggerConfig.Configs[env]
	if conf == nil {
		return nil, errors.New("Error in logger config file")
	}

	level, err := getLevel(conf.Level)
	if err != nil {
		// Error: Invalid level found, Setting log level to debug.
	}

	zapWriteSyncers := make([]zapcore.WriteSyncer, 0)
	for _, fileName := range conf.OutputPaths {
		zapWriteSyncers = append(zapWriteSyncers, getWriteSyncers(fileName))
	}

	logConfig := &logger.LoggerConfig{LogLevel: level, WriteSyncers: zapWriteSyncers}

	log := logger.NewLogger()
	log.UpdateConfig(logConfig)
	log.AddTags(conf.Appends)
	return log, nil
}

func getParsedLoggerYamlConfig() (*LoggerConfigs, error) {
	loggerConfigs := new(LoggerConfigs)
	pwd, _ := os.Getwd()
	currentFilePath := pwd + "/utils/loggerUtil/logger_config.yaml"
	if !(os.Getenv("LAMBDA_TASK_ROOT") == "" && os.Getenv("AWS_EXECUTION_ENV") == "") {
		lambdaPath, _ := os.Executable()
		currentFilePath = lambdaPath
	}
	yamlData, err := os.ReadFile(path.Join(path.Dir(currentFilePath), "logger_config.yaml"))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlData, loggerConfigs)
	if err != nil {
		return nil, err
	}

	return loggerConfigs, nil
}

func getLevel(level string) (logger.LogLevel, error) {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return logger.LogLevelDebug, nil
	case "INFO":
		return logger.LogLevelInfo, nil
	case "WARN":
		return logger.LogLevelWarn, nil
	case "ERROR":
		return logger.LogLevelError, nil
	case "FATAL":
		return logger.LogLevelFatal, nil
	default:
		return logger.LogLevelDebug, errors.New("unsupported level")
	}
}

func getWriteSyncers(fileName string) zapcore.WriteSyncer {
	if strings.Compare(fileName, "stdout") == 0 {
		return zapcore.AddSync(os.Stdout)
	}

	if strings.Compare(fileName, "stderr") == 0 {
		return zapcore.AddSync(os.Stderr)
	}

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    50, // megabytes
		MaxBackups: 1,
		MaxAge:     1, // days
	})
}
