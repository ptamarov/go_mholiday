package main

import (
	"fmt"
)

type Data map[string]any // alias for a dict into anything

// MatchNum checks if a given key is present in data and whether its value matches the expected float.
func MatchNum(key string, expected float64, data Data) error {
	v, ok := data[key]
	if !ok {
		return fmt.Errorf("key \"%s\" missing", key)
	}

	value, ok := v.(float64)
	if !ok {
		return fmt.Errorf("wrong value type for key \"%s\"", key)
	}
	if value != expected {
		return fmt.Errorf("wrong value %.2f for key \"%s\" (expected %.2f)", value, key, expected)
	}

	return nil
}

// MatchString check if a given key is present in data and whether its value matches the expected string. Matching is case insensivitve.
func MatchString(key string, expected string, data Data) error {
	v, ok := data[key]
	if !ok {
		return fmt.Errorf("key \"%s\" missing", key)
	}
	value, ok := v.(string)

	if !ok {
		return fmt.Errorf("wrong value type for key \"%s\"", key)
	}
	if value != expected {
		return fmt.Errorf("wrong value %s for key \"%s\" (expected %s)", value, key, expected)
	}
	return nil
}

// Contains checks if the expectedData is contained in a given data. It does so
// by visiting each field of the expected data recursively and checking that it is
// present in the data.
func Contains(expectedData, data Data) error {
	for k, v := range expectedData {
		switch x := v.(type) { // get the type of v
		case float64:
			if err := MatchNum(k, x, data); err != nil {
				return err
			}
		case string:
			if err := MatchString(k, x, data); err != nil {
				return err
			}
		case Data: // the value is another dictionary
			if val, ok := data[k]; !ok { // check of the current key belongs to the dictionary
				return fmt.Errorf("key \"%s\" missing", k)
			} else if subData, ok := val.(Data); ok { // check if the value is another dictionary
				if err := Contains(x, subData); err != nil { // recursively check if x is contained in subData
					return fmt.Errorf("value %s unmatched %#v: %s", k, x, err)
				}
			} else {
				return fmt.Errorf("%s unmatched in %#v", k, val)
			}
		}
	}
	return nil
}
