package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/1nterdigital/aka-im-discover/pkg/common/constant"
	"github.com/1nterdigital/aka-im-tools/errs"
)

func Load(configDirectory, configFileName, envPrefix, runtimeEnv string, cfg any) (err error) {
	if runtimeEnv == constant.KUBERNETES {
		mountPath := os.Getenv(constant.MountConfigFilePath)
		if mountPath == "" {
			return errs.ErrArgs.WrapMsg(constant.MountConfigFilePath + " env is empty")
		}

		return loadConfig(filepath.Join(mountPath, configFileName), envPrefix, cfg)
	}

	return loadConfig(filepath.Join(configDirectory, configFileName), envPrefix, cfg)
}

func loadConfig(path, envPrefix string, cfg any) (err error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = v.ReadInConfig(); err != nil {
		return errs.WrapMsg(err, "failed to read config file", "path", path, "envPrefix", envPrefix)
	}

	if err = v.Unmarshal(cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = "mapstructure"
	}); err != nil {
		return errs.WrapMsg(err, "failed to unmarshal config", "path", path, "envPrefix", envPrefix)
	}

	return nil
}
