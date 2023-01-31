package likeexpr

import (
	"regexp"
)

// Convert SQL Like patterns to Go regular expressions.
// pat          - Like pattern
// escape       - Escape character
// escapeItself - Escaping the escape character itself
func ToRegexp(pat string, escape rune, escapeItself bool) string {
	src := []rune(pat)
	length := len(src)
	buf := make([]rune, 0, length)
	ret := make([]rune, 0, length+64)
	ret = append(ret, '^')

OUTER:
	for i := 0; i < length; i++ {
		c := src[i]
		switch c {
		case escape:
			if i+1 == length {
				buf = append(buf, c)
				break OUTER
			} else {
				switch src[i+1] {
				case '_', '%':
					i = i + 1
				case escape:
					i = i + 1
					if !escapeItself {
						buf = append(buf, c)
					}
				}
				buf = append(buf, src[i])
			}
		case '_':
			ret = append(ret, []rune(regexp.QuoteMeta(string(buf)))...)
			ret = append(ret, '.')
			buf = buf[0:0]
		case '%':
			ret = append(ret, []rune(regexp.QuoteMeta(string(buf)))...)
			ret = append(ret, '.', '*', '?')
			buf = buf[0:0]
		default:
			buf = append(buf, c)
		}
	}

	ret = append(ret, []rune(regexp.QuoteMeta(string(buf)))...)
	ret = append(ret, '$')

	return string(ret)
}
