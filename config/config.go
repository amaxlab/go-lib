package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConfigLoader struct {}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (c *ConfigLoader) String(name, defaultValue string) string {
	val, ok := os.LookupEnv(c.getEnvVarName(name))
	if ! ok {
		return defaultValue
	}

	return val
}

func (c *ConfigLoader) Int(name string, defaultValue int) int {
	val, ok := os.LookupEnv(c.getEnvVarName(name))
	if !ok {
		return defaultValue
	}

	i64, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int(i64)
}

func (c *ConfigLoader) Bool(name string, defaultValue bool) bool {
	val, ok := os.LookupEnv(c.getEnvVarName(name))
	if !ok {
		return defaultValue
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}

	return b
}

func (c *ConfigLoader) getEnvVarName(name string) string {
	return fmt.Sprintf("APP_%s", strings.ToUpper(strings.Replace(name, "-", "_", -1)))
}