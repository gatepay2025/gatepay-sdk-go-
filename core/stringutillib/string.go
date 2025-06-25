package stringutillib

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func JoinStrSepByDoubleQuot(strList []string) string {
	return "\"" + strings.Join(strList, "\", \"") + "\""
}

func IsStringInSlice(str string, strs []string) bool {
	for _, item := range strs {
		if str == item {
			return true
		}
	}
	return false
}

func IsSameStringSlice(strSlice1, strSlice2 []string) bool {
	if len(strSlice1) != len(strSlice2) {
		return false
	}
	if len(strSlice1) == 0 {
		return true
	}

	sort.Slice(strSlice1, func(i, j int) bool {
		return strSlice1[i] < strSlice1[j]
	})
	sort.Slice(strSlice2, func(i, j int) bool {
		return strSlice2[i] < strSlice2[j]
	})
	for idx := range strSlice1 {
		if strSlice1[idx] != strSlice2[idx] {
			return false
		}
	}
	return true
}

func ObjToJsonStr(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}

func JsonStrToObject(data string, obj interface{}) error {
	return json.Unmarshal([]byte(data), obj)
}

func IsEmptyStr(str string) bool {
	return strings.TrimSpace(str) == ""
}

func IsEmptyJsonStr(str string) bool {
	str = strings.TrimSpace(str)
	return str == "" || str == "{}" || str == "[]"
}

func AddOneBasedIntStr(numStr string, add int64) (string, error) {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return "", err
	}
	num += add
	return strconv.FormatInt(num, 10), nil
}

func SplitNumsStringToInt64Slice(str, separator string) ([]int64, error) {
	if len(str) == 0 {
		return make([]int64, 0), nil
	}
	numStrs := strings.Split(str, separator)
	res := make([]int64, len(numStrs))
	for index, num := range numStrs {
		intNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			return nil, err
		}
		res[index] = intNum
	}
	return res, nil
}

func IsValidUrl(rawUrl string) bool {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return false
	}

	// 检查协议是否为 http 或 https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}
	if len(parsedURL.Host) == 0 && len(parsedURL.Path) == 0 {
		return false
	}
	return true
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)

	for i := range randomString {
		// 使用 crypto/rand 生成一个随机索引
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		randomString[i] = charset[num.Int64()]
	}

	return string(randomString)
}

func IsAlphanumeric(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(s)
}

func StringToInt64(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

func AnyToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	default:
		// 处理其他类型
		return fmt.Sprintf("%v", value)
	}
}

func SplitToArray(str string) []string {
	if !IsEmptyStr(str) {
		list := strings.Split(str, ",")
		if len(list) > 0 {
			return list
		}
	}
	return []string{}
}
