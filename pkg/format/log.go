package format

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	logFields = []string{
		"log_id",
		"user_id",
		"http_code",
		"uri",
		"cost_time",
		"perf_id",
		"ip",
		"method",
		"user_agent",
	}
)

type ServiceLogFormatter struct{}

func (slf *ServiceLogFormatter) Format(entry *log.Entry) ([]byte, error) {
	output := "[%lvl%][%time%] —"
	timestampFormat := "2006-01-02 15:04:05.999"

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)
	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	funcVal := fmt.Sprintf("%s", entry.Caller.Function)
	funcValList := strings.Split(funcVal, "/")

	funcName := funcValList[len(funcValList)-1]
	output += " func=" + funcName + ";"
	line := fmt.Sprintf("%d", entry.Caller.Line)
	output += " line=" + line + ";"

	for k, val := range entry.Data {
		if k == "userAgent" {
			continue
		}
		stringVal, ok := val.(string)
		if !ok {
			stringVal = fmt.Sprint(val)
		}

		output += " " + k + "=" + stringVal + ";"
	}
	output += " msg=" + entry.Message + ";\n"
	return []byte(output), nil
}

type GinAccessLogFormatter struct{}

func (galf *GinAccessLogFormatter) Format(entry *log.Entry) ([]byte, error) {
	output := "[%time%] —"
	timestampFormat := "2006-01-02 15:04:05.999"

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	for _, field := range logFields {
		if value, exist := entry.Data[field]; exist {
			stringVal, ok := value.(string)
			if !ok {
				stringVal = fmt.Sprint(value)
			}
			output += " " + field + "=" + stringVal + ";"
		}
	}
	output += "\n"
	return []byte(output), nil
}
