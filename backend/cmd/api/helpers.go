package main

import (
	"net/url"
	"strconv"
	"strings"
)

type envelope map[string]any

func (cfg *Config) readString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}
	return s
}

func (cfg *Config) readCSVStrings(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)
	if csv == "" {
		return defaultValue
	}
	return strings.Split(csv, ",")
}

func (cfg *Config) readCSVInt64s(qs url.Values, key string, defaultValue []int64) ([]int64, error) {
	csv := cfg.readCSVStrings(qs, key, []string{})
	for _, s := range csv {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		defaultValue = append(defaultValue, int64(n))
	}
	return defaultValue, nil
}
