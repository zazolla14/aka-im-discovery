package config

import (
	"strings"
)

var (
	ShareFileName           = "share.yml"
	DiscoveryConfigFileName = "discovery.yml"
	DiscoverApiCfgFileName  = "discover-api.yml"
	LogConfigFileName       = "log.yml"
	AdminFileName           = "admin.yml"
	MongodbConfigFileName   = "mongodb.yml"
	MysqlConfigFileName     = "mysqldb.yml"
	RedisConfigFileName     = "redis.yml"
	TracerConfigFileName    = "tracer.yml"
)

var EnvPrefixMap map[string]string

func init() {
	EnvPrefixMap = make(map[string]string)
	fileNames := []string{
		ShareFileName,
		AdminFileName,
		DiscoverApiCfgFileName,
		DiscoveryConfigFileName,
		MongodbConfigFileName,
		LogConfigFileName,
		RedisConfigFileName,
		TracerConfigFileName,
	}

	for _, fileName := range fileNames {
		envKey := strings.TrimSuffix(strings.TrimSuffix(fileName, ".yml"), ".yaml")
		envKey = "DISCOVERENV_" + envKey
		envKey = strings.ToUpper(strings.ReplaceAll(envKey, "-", "_"))
		EnvPrefixMap[fileName] = envKey
	}
}

const (
	FlagConf          = "config_folder_path"
	FlagTransferIndex = "index"
)
