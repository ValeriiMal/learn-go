package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	validStartStrings := []string{
		"[TRC]",
		"[DBG]",
		"[INF]",
		"[WRN]",
		"[ERR]",
		"[FTL]",
	}
	result := false
	for _,v:=range validStartStrings {
		re, err := regexp.Compile(v)
		if err != nil {
			break
		}

		i := re.FindStringIndex(text)
		if len(i) > 0 && i[0] == 0 {
			result = true
			break
		}
	}
	return result
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<(-*|=*|\**|~*)>`)
	if err != nil {
		return make([]string, 0)
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if strings.Contains(line, "\"password\"") {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)
	if err != nil {
		return text
	}
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		if strings.Contains(line, "User ") {
			re, err := regexp.Compile(`User\s+(\w+)`)
			if err != nil {
				return lines
			}
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				lines[i] = "[USR] " + matches[1] + " " + line
			}
		}
	}
	return lines
}
