package validator

import (
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Add(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// check adds an error to the map if the validation check is false
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.Add(key, message)
	}
}

func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique(value []string) bool {
	UniqueValues := make(map[string]bool)
	for _, value := range value {
		UniqueValues[value] = true
	}
	return len(value) == len(UniqueValues)
}
