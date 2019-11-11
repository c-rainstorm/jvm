package formatter

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"jvm/pkg/constants"
	"strings"
)

type LogrusFormatter struct {
	TimeFormatter  string
	FieldPrefix    string
	FieldSuffix    string
	FieldSeparator string
}

var DefaultLogrusFormatter = LogrusFormatter{
	TimeFormatter:  "2006-01-02 15:04:05",
	FieldPrefix:    "[",
	FieldSuffix:    "]",
	FieldSeparator: constants.Space,
}

func NewLogrusFormatter() *LogrusFormatter {
	return &DefaultLogrusFormatter
}

func (this *LogrusFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logItems := []string{
		this.around(entry.Time.Format(this.TimeFormatter)),
		this.around(strings.ToUpper(entry.Level.String())),
		entry.Message,
	}

	serialized := strings.Join(logItems, this.FieldSeparator)

	return append(bytes.NewBufferString(serialized).Bytes(), '\n'), nil
}

func (this *LogrusFormatter) around(msg string) string {
	return this.FieldPrefix + msg + this.FieldSuffix
}
