package task

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
)

func Convert(yamlCode string, indent string) string {
	t := map[string]interface{}{}

	err := yaml.Unmarshal([]byte(yamlCode), &t)

	if err != nil {
		return err.Error()
	}

	fmt.Printf("%v", t)

	result := formatMap(t, 0, indent, "Config")
	println(result)
	return result
}

func formatMap(m map[string]interface{}, index int, indent string, outName string) string {
	result := make([]string, 0)

	result = append(result, getHeader(index, indent, outName))

	for key, value := range m {
		if reflect.TypeOf(value).Kind() == reflect.String {
			result = append(result, getBody(index, indent, key, value))
		} else if reflect.TypeOf(value).Kind() == reflect.Map {
			newValue := getMap(value.(map[interface{}]interface{}))
			result = append(result, formatMap(newValue, index+1, indent, key))
		} else if reflect.TypeOf(value).Kind() == reflect.Array {
			result = append(result, getArrayBody(index, indent, key, value.([]interface{})))
		}
	}

	result = append(result, getFooter(index, indent, outName))

	return strings.Join(result, "\n")
}

func getMap(m map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		newKey := k.(string)
		result[newKey] = v
	}
	return result
}

func getHeader(index int, indent, outName string) string {
	var result string
	for i:=0; i<index; i++ {
		result += indent
	}
	if index == 0 {
		result += "type "
	}
	result += outName
	result += " struct {"

	return result
}

func getBody(index int, indent, key string, value interface{}) string {
	var result string
	for i:=0; i<index+1; i++ {
		result += indent
	}
	result += key
	result += " "
	result += reflect.TypeOf(value).String()
	result += " "
	result += "`yaml:\"" + key + "\"`"

	return result
}

func getArrayBody(index int, indent, key string, value []interface{}) string {
	var result string
	for i:=0; i<index+1; i++ {
		result += indent
	}
	result += key
	result += " []"
	result += reflect.TypeOf(value).String()
	result += " "
	result += "`yaml:\"" + key + "\"`"

	return result

}

func getFooter(index int, indent, outName string) string {
	var result string
	for i:=0; i<index; i++ {
		result += indent
	}
	result += "} `yaml:\"" + outName + "\"`"

	return result
}




