package parsinglogfiles

import "regexp"

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
		i := re.FindIndex([]byte(text))
		if i[0] == 0 {
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
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
