package util

import (
	"os"
	"strings"
)

const (
	env_flag = "pome"
	prod_env = "prod"
	dev_env  = "dev"
)

var (
	cur_env = dev_env
)

func init() {
	cur_env = strings.ToLower(os.Getenv(env_flag))
	cur_env = strings.TrimSpace(cur_env)

	if len(cur_env) == 0 {
		cur_env = dev_env
	}
}

func IsProduct() bool {
	return cur_env == prod_env
}

func IsDev() bool {
	return cur_env == dev_env
}

func GetEnv() string {
	return cur_env
}
