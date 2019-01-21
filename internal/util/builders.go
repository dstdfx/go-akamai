package util

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// BuildServiceEndpoint concatenates given params into service endpoint.
// It builds `HTTPS` endpoint by default. To overwrite it, add `http://` to host param.
func BuildServiceEndpoint(host, service, version string) string {
	if strings.HasPrefix(host, "http://") ||
		strings.HasPrefix(host, "https://") {
		return strings.Join([]string{host, service, version}, "/")
	}
	return fmt.Sprintf("https://%s/%s/%s",
		host,
		service,
		version)
}

// BuildQueryParameters converts provided options struct to the string of URL parameters.
func BuildQueryParameters(opts interface{}) (string, error) {
	optsValue := reflect.ValueOf(opts)
	if optsValue.Kind() != reflect.Struct {
		return "", errors.New("provided options is not a structure")
	}
	optsType := reflect.TypeOf(opts)

	params := url.Values{}

	for i := 0; i < optsValue.NumField(); i++ {
		fieldValue := optsValue.Field(i)
		fieldType := optsType.Field(i)

		queryTag := fieldType.Tag.Get("param")
		if queryTag != "" {
			if isZero(fieldValue) {
				continue
			}

			tags := strings.Split(queryTag, ",")
		loop:
			switch fieldValue.Kind() {
			case reflect.Ptr:
				fieldValue = fieldValue.Elem()
				goto loop
			case reflect.String:
				params.Add(tags[0], fieldValue.String())
			case reflect.Int:
				params.Add(tags[0], strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Bool:
				params.Add(tags[0], strconv.FormatBool(fieldValue.Bool()))
			}
		}
	}

	return params.Encode(), nil
}

// isZero checks if provided value is zero.
func isZero(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	z := reflect.Zero(v.Type())

	return v.Interface() == z.Interface()
}

