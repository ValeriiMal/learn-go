package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
    count := 0
    for _,line := range lines {
        if re.MatchString(line) {
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
