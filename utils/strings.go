package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/araddon/dateparse"
)

type String string

func (c String) trans(tp string, size int) (interface{}, error) {
	switch strings.ToLower(tp) {
	case "int":
		if size == 0 {
			return strconv.Atoi(c.String())
		}
		return strconv.ParseInt(c.String(), 10, size)
	case "bool":
		return strconv.ParseBool(c.String())
	case "time":
		timeTemp := c.String()
		if len(timeTemp) < 1 {
			if size >= 0 {
				return time.Unix(int64(size), 0), nil
			}
			timeTemp = "0"
		}
		return dateparse.ParseAny(timeTemp)
	default:
		return nil, errors.New("not suport")

	}

}

func (c String) Max(m int) String {
	if c.Int() > m {
	}

	return c
}

func (c String) String() string {
	return strings.TrimSpace(string(c))
}

func (c String) MustInt() (int, error) {
	result, err := c.trans("int", 0)
	return result.(int), err
}

func (c String) Int() int {
	result, err := c.trans("int", 0)
	if err != nil {
		//panic(c.String() + " type must be int")
		return 0
	}
	return result.(int)
}

func (c String) Int64() int64 {
	result, err := c.trans("int", 64)
	if err != nil {
		panic(c.String() + " type must be int64")
	}
	return result.(int64)
}

func (c String) Bool() bool {
	result, err := c.trans("bool", 0)
	if err != nil {
		return false
	}
	return result.(bool)
}

func (c String) ToTime(defaultTimeStamps ...int64) (time.Time, error) {
	defaultTimeStamp := 0
	if len(defaultTimeStamps) > 0 {
		defaultTimeStamp = int(defaultTimeStamps[0])
	}
	result, err := c.trans("time", defaultTimeStamp)
	return result.(time.Time), err
}

func (c String) ToArrString(k string) ([]string, error) {
	if len(k) < 1 {
		return nil, fmt.Errorf("empty key")
	}
	return strings.Split(c.String(), k), nil
}

func IsPrefixLowerCase(str string) bool {
	if len(str) <= 0 {
		return false
	}
	return unicode.IsLower(rune(str[0]))
}

func Unmarshal(data []byte, obj interface{}) error {
	stcMap := map[string]interface{}{}
	if err := json.Unmarshal(data, &stcMap); err != nil {
		return err
	}
	objValue := ValueOf(obj)
	stcs := []reflect.Value{objValue}
	i := 0
	for objValue.Type().Kind() == reflect.Struct {
		if i >= len(stcs) {
			break
		}
		v := stcs[i]
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			vfield := elem(v.Field(i))
			isStruct := false
			if vfield.Kind() == reflect.Struct {
				isStruct = true
			}
			if !isStruct && IsPrefixLowerCase(field.Name) {
				continue
			}
			if ok := field.Tag.Get("required"); ok != "true" {
				if isStruct {
					stcs = append(stcs, vfield)
				}
				continue
			}
			fieldName := field.Tag.Get("json")
			_, ok := stcMap[fieldName]
			if !ok {
				return fmt.Errorf("%s required but get none", fieldName)
			}
		}
		i++
	}
	return json.Unmarshal(data, obj)
}

func GenUrl(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}

	return "http://" + strings.Trim(url, "/")
}

func Shorten(o string, length int) string {
	if len(o) > length {
		return o[:length-3] + ".."
	}
	return o
}

func Split(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func EncodeMd5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
