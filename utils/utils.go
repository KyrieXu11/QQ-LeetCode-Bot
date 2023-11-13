package utils

import (
	"crypto/md5"
	"github.com/jmespath/go-jmespath"
	"net/textproto"
	"strings"
)

func Md5(s string) [16]byte {
	return md5.Sum([]byte(s))
}

func GetDictValueByKey(inter interface{}, path string) (interface{}, error) {
	return jmespath.Search(path, inter)
}

func ParseCookies(cookieString string) map[string]string {
	lines := []string{cookieString}
	if len(lines) == 0 {
		return nil
	}

	cookies := make(map[string]string)
	for _, line := range lines {
		line = textproto.TrimString(line)

		var part string
		for len(line) > 0 {
			part, line, _ = strings.Cut(line, ";")
			part = textproto.TrimString(part)
			if part == "" {
				continue
			}
			name, val, _ := strings.Cut(part, "=")
			val, ok := parseCookieValue(val, true)
			if !ok {
				continue
			}
			cookies[name] = val
		}
	}
	return cookies
}

func parseCookieValue(raw string, allowDoubleQuote bool) (string, bool) {
	// Strip the quotes, if present.
	if allowDoubleQuote && len(raw) > 1 && raw[0] == '"' && raw[len(raw)-1] == '"' {
		raw = raw[1 : len(raw)-1]
	}
	return raw, true
}

// 算了，去他妈的自己造轮子，用开源的得了，超
//type Dict map[string]interface{}
//
//func (o Dict) Get(key string) (interface{}, error) {
//	return jmespath.Search(key, o)
//}

//keys := strings.Split(key, ".")
//m := o
//for _, key := range keys {
//	if v, b := m[key]; !b {
//		return nil
//	} else {
//		switch v.(type) {
//		case Dict:
//			m = v.(Dict)
//		case map[string]interface{}:
//			m = v.(map[string]interface{})
//		default:
//			return v
//		}
//	}
//}
//return m
//}

//const (
//	DIGEST            = "0123456789"
//	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
//	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//	LETTERS           = LOWERCASE_LETTERS + UPPERCASE_LETTERS
//)

//type Parser struct {
//}
//
//type Lexer struct {
//	curChar rune
//	text    []rune
//	idx     int
//}
//
//func NewLexer(path string) {
//}
//
//func (o *Lexer) GetNextToken() (interface{}, error) {
//	for o.curChar != 0 {
//		curStr := o.curString()
//		if strings.Contains(DIGEST, curStr) {
//			return o.makeNum()
//		} else if strings.Contains(LETTERS, curStr) {
//			return o.makeIdentify()
//		}
//	}
//}
//
//func (o *Lexer) makeIdentify() (string, error) {
//	res := ""
//	for o.curChar != 0 && strings.Contains(LETTERS+"_", o.curString()) {
//		res += o.curString()
//		o.advance()
//	}
//	return res, nil
//}
//
//func (o *Lexer) makeNum() (interface{}, error) {
//	numStr := ""
//	dotCount := 0
//	for o.curChar != 0 && strings.Contains(DIGEST+".", o.curString()) {
//		if o.curChar == '.' {
//			if dotCount == 1 {
//				break
//			}
//			dotCount++
//			numStr += "."
//		} else {
//			numStr += o.curString()
//		}
//		o.advance()
//	}
//	if dotCount == 1 {
//		return strconv.ParseFloat(numStr, 10)
//	} else {
//		return strconv.Atoi(numStr)
//	}
//}
//
//func (o *Lexer) curString() string {
//	return string(o.curChar)
//}
//
//func (o *Lexer) advance() {
//	o.idx++
//}
