package validator

import "regexp"

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	Errors map[string]string
}

// Return Validator instance with an empty errors map.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Return true if the errors map doesn't contain any entries.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// Add an error message to errors map if no entry already exists for given key.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Add an error message to errors map only if validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Return true if a specific value is in a list of strings.
func In(val string, list ...string) bool {
	for i := range list {
		if list[i] == val {
			return true
		}
	}
	return false
}

// Return true if a string value matches a specific regexp pattern.
func Matches(val string, rx *regexp.Regexp) bool {
	return rx.MatchString(val)
}

// Return true if all string values in a slice are unique.
func Unique(vals []string) bool {
	uniqueValues := make(map[string]bool)
	for _, val := range vals {
		uniqueValues[val] = true
	}
	return len(uniqueValues) == len(vals)
}
