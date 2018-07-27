package aetable

import (
	"errors"
	"regexp"
	"strings"
)

type AETableConfig struct {
	Title    string
	XLSXFile string
	Sheet    string
}

func ParseConfig(str string) (*AETableConfig, error) {
	str = strings.TrimSpace(str)
	if !strings.HasPrefix(str, "{aetable") {
		return nil, nil
	}

	hasBrackets := strings.Contains(str, "{") && strings.Contains(str, "}")
	if !hasBrackets {
		err := errors.New("Not a valid AETable config")
		return nil, err
	}

	content := str[strings.Index(str, "{")+1+len("aetable") : strings.Index(str, "}")]

	r := regexp.MustCompile(`\S+=\".+?\"`)
	// fields := strings.Fields(content)
	fields := r.FindAllString(content, -1)

	config := AETableConfig{}
	for index := 0; index < len(fields); index++ {
		f := strings.TrimSpace(fields[index])
		equalIndex := strings.Index(f, "=")
		key := f[:equalIndex]
		value := f[equalIndex+2 : len(f)-1]
		switch key {
		case "title":
			config.Title = value
		case "file":
			config.XLSXFile = value
		case "sheet":
			config.Sheet = value
		}
		// log.Println(key, value, "-")

	}
	// log.Println(fields, len(fields))
	return &config, nil
}
