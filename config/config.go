package config

import (
	"os"
	"strconv"
)

func GetStringValue(name, defaultValue string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	return val
}

func GetIntValue(name string, defaultValue int) int {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	i64, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int(i64)
}

func GetBoolValue(name string, defaultValue bool) bool {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}

	return b
}
