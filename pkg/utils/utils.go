package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// EnsurePwdDir : ensure dir exist
func EnsurePwdDir(targetDir string) (string, error) {
	_, err := os.Stat(targetDir)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			err = os.MkdirAll(targetDir, 0775)
		}
	}
	return targetDir, err
}

// GetParentDir : get start path parent
func GetParentDir() string {
	pwdPath, _ := os.Getwd()
	return filepath.Dir(pwdPath)
}

// GetProjectDir : get start path
func GetProjectDir() string {
	pwdPath, _ := os.Getwd()
	return pwdPath
}

// ToInterfaceSlice : trans any types data list to interface
func ToInterfaceSlice(slice any) []any {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}

	ret := make([]any, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// If returns interface type value
// condition? trueVal: falseVal
func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}

// GetRandom 根据max min，获取(min, max)的值
func GetRandom(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

// GetRandomNum
func GetRandomNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

// 在start, end的范围内生成n个随机数
func GetRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}

	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn(end-start) + start

		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

// 指定的字符串模板和变量map，生成目标字符串
// 示例：
// ProcessStringTemplate("hello {{.target}}", map[string]interface{}{"target":"world"})
// return hello wolrd
func ProcessStringTemplate(templateString string, vars any) (string, error) {
	// 创建模板
	tmpl, err := template.New("tmpl").Parse(templateString)
	if err != nil {
		return "", err
	}
	// 执行变量重写
	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, vars)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

// PageCompute 分页请求
func PageCompute(page, pageSize, length int) (start, end int) {
	from := (page - 1) * pageSize
	if from < 0 || pageSize < 0 {
		return 0, length
	}
	if from > length {
		return 0, 0
	}
	if from+pageSize > length {
		return from, length
	}
	return from, from + pageSize
}

// func GenerateBatchID(jobType string, notifiedType string) string {
// 	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
// 	uuid := GenUUID()
// 	return jobType + "_" + notifiedType + "_" + GenMD5(timestamp+uuid)
// }

func ToJSONString(object any) string {
	jsonBytes, err := json.Marshal(object)
	if nil != err {
		return ""
	}
	return string(jsonBytes)
}

// IsOddNum 是否是奇数
func IsOddNum(num int64) bool {
	return num%2 == 1
}

// HasKeyWord str是否存在关键字key
func HasKeyWord(str, key string) bool {
	r := regexp.MustCompile(key)
	return r.FindStringIndex(str) != nil
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GenSHA256(str string) string {
	cipherByte := sha256.Sum256([]byte(str))
	token := hex.EncodeToString(cipherByte[:])
	return token
}

// figure out which endpoint the request comes from
func IsMobileOrPC(UserAgent string) string {
	endpointMap := map[string]string{
		"mobile":  "mobile",
		"iOS":     "mobile",
		"Android": "mobile",
		"Mac":     "pc",
		"Windows": "pc",
	}
	patternMap := map[string]string{
		"mobile":  "(?i)Android|webOS|iPhone|iPod|BlackBerry",
		"Mac":     "(?i)Macintosh|macOS",
		"iOS":     "(?i)iPhone|iPad|iPod",
		"Windows": "(?i)Windows",
		"Android": "(?i)Android",
	}

	for i := range patternMap {
		if match, _ := regexp.MatchString(patternMap[i], UserAgent); match {
			return endpointMap[i]
		}
	}

	return "mobile"
}

func GetTerminalByUserAgent(UserAgent string) string {
	patternMap := map[string]string{
		"Mac":     "(?i)Macintosh|macOS",
		"iOS":     "(?i)iPhone|iPad|iPod",
		"Windows": "(?i)Windows",
		"Android": "(?i)Android",
	}

	for i, v := range patternMap {
		if match, _ := regexp.MatchString(v, UserAgent); match {
			return i
		}
	}

	return ""
}

func Contain(obj any, target any) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func GeneratePlatformAuthorization(ak string, sk string) string {
	timestamp := time.Now().UnixNano() / 1e6
	timestampStr := strconv.FormatInt(timestamp, 10)
	content := fmt.Sprintf("%s%s%s", ak, sk, timestampStr)
	sign := GenSHA256(content)
	newContent := fmt.Sprintf("%s.%s.%s", ak, sign, timestampStr)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(newContent))
	return encodedStr
}

func GeneratePlatformAuthorizationWithTimestamp(ak string, sk string, timestamp int64) string {
	timestampStr := strconv.FormatInt(timestamp, 10)
	content := fmt.Sprintf("%s%s%s", ak, sk, timestampStr)
	sign := GenSHA256(content)
	newContent := fmt.Sprintf("%s.%s.%s", ak, sign, timestampStr)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(newContent))
	return encodedStr
}

func ParsePlatformAuthorization(authorization string) (string, string, error) {
	dataByte, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		return "", "", err
	}
	dataStr := string(dataByte)
	dataSlice := strings.Split(dataStr, ".")
	if len(dataSlice) != 3 {
		return "", "", errors.New("authorization format error")
	}
	ak := dataSlice[0]
	timestampStr := dataSlice[2]
	return ak, timestampStr, nil
}

// MinInt64 返回给定两个 int64 类型整数中的较小值
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 返回两个 int64 类型数值中的最大值
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// @probability: 事件发生的概率
func HappenByProbability(probability float32) bool {
	if probability > 1 {
		return false
	}
	// 随机生成一个[0, max)的数
	max := 10000
	num := rand.Intn(max)
	if float32(num) <= float32(max)*probability {
		return true
	}
	return false
}

func IsToday(sendTime int64) bool {
	if sendTime <= 0 {
		return true
	}
	second := sendTime / 1000
	t := time.Unix(second, 0)
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// RemoveBrackets RemoveBrackets：去除中括号， @content: 待去除的字符串
func RemoveBrackets(content string) string {
	content = strings.ReplaceAll(content, "[", "")
	content = strings.ReplaceAll(content, "]", "")
	return content
}

func ObjectToLogStr(obj any) string {
	logStr, _ := json.Marshal(obj)
	// strans := fmt.Sprintf("%v", obj)
	// strans = strings.ReplaceAll(strans, "\n", "--*--")
	// return strans
	return string(logStr)
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func IsNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}
