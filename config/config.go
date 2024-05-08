package config

import (
	"fmt"
	"github.com/symsimmy/due/env"
	"github.com/symsimmy/due/flag"
	"github.com/symsimmy/due/value"
)

const (
	dueConfigArgName              = "config"
	dueConfigEnvName              = "GO_ENV"
	dueConfigDirEnvName           = "CONFIG_DIR"
	dueOverWriteConfigPathEnvName = "OVERWRITE_CONFIG"
	defaultConfigPath             = "./configs"
	defaultConfigFileName         = "default"

	defaultRemoteConfigEnvName = "REMOTE_CONFIG"
	defaultRemoteConfig        = "apollo"
)

var globalReader Reader

func init() {
	// get config dir
	configPath := env.Get(dueConfigDirEnvName).String()
	if len(configPath) == 0 {
		configPath = defaultConfigPath
	}
	// get path from command line
	filePath := getFileFullName(defaultConfigFileName, configPath)
	fileExtension := getFileExtension(filePath)
	path := flag.String(dueConfigArgName, filePath)

	// overwrite path from env
	var overwritePath string
	localEnv := env.Get(dueConfigEnvName).String()
	if localEnv != "" {
		overwritePath = fmt.Sprintf("%v/%v.%v", configPath, localEnv, fileExtension)
	}

	// remote overwrite path from env
	var remoteOverwritePath string
	overwriteEnv := env.Get(dueOverWriteConfigPathEnvName).String()
	if overwriteEnv != "" {
		remoteOverwritePath = overwriteEnv
	}

	remoteEnv := env.Get(defaultRemoteConfigEnvName, defaultRemoteConfig).String()

	// set local reader
	SetReader(NewReader(WithSources(NewSource(path), NewSource(overwritePath), NewSource(remoteOverwritePath)), WithRemoteSources(remoteEnv)))

	//InitConsulConfig(context.Background())
}

// SetReader 设置配置读取器
func SetReader(reader Reader) {
	if globalReader != nil {
		globalReader.Close()
	}
	globalReader = reader
}

// GetReader 获取配置读取器
func GetReader() Reader {
	return globalReader
}

// Has 是否存在配置
func Has(pattern string) bool {
	return globalReader.Has(pattern)
}

// Get 获取配置值
func Get(pattern string, def ...interface{}) value.Value {
	return globalReader.Get(pattern, def...)
}

// Set 设置配置值
func Set(pattern string, value interface{}) error {
	return globalReader.Set(pattern, value)
}

// Close 关闭配置监听
func Close() {
	if globalReader != nil {
		globalReader.Close()
	}
}
