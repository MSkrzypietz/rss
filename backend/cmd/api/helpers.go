package main

import (
	"net/url"
	"strconv"
	"strings"
)

type envelope map[string]any

func (app *application) readString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}
	return s
}

func (app *application) readCSVStrings(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)
	if csv == "" {
		return defaultValue
	}
	return strings.Split(csv, ",")
}

func (app *application) readCSVInt64s(qs url.Values, key string, defaultValue []int64) ([]int64, error) {
	csv := app.readCSVStrings(qs, key, []string{})
	for _, s := range csv {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		defaultValue = append(defaultValue, int64(n))
	}
	return defaultValue, nil
}
