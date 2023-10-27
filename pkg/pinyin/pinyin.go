package pinyin

import (
	"github.com/mozillazg/go-pinyin"
	"regexp"
	"strings"
)

func GetPinyin(txt string) string {
	sb := new(strings.Builder)

	for _, c := range txt {
		if c > 128 {
			str := string([]rune{c})
			strs := pinyin.LazyConvert(str, nil)

			if len(strs) > 0 {
				sb.WriteString(strs[0])
			} else {
				sb.WriteString(str)
			}

		} else {
			sb.WriteString(string([]rune{c}))
		}
	}

	return strings.TrimSpace(sb.String())
}

func GetFirstSpell(txt string) string {
	sb := new(strings.Builder)

	for _, c := range txt {
		if c > 128 {
			str := string([]rune{c})
			strs := pinyin.LazyConvert(str, nil)

			if len(strs) > 0 {
				sb.WriteString(string(strs[0][0]))
			} else {
				sb.WriteString(str)
			}

		} else {
			r := regexp.MustCompile(`\W`)
			str := r.ReplaceAllString(string([]rune{c}), "")

			if str != "" {
				sb.WriteString(str)
			}
		}
	}

	return strings.TrimSpace(sb.String())
}
