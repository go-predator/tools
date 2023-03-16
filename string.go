/**
 * @Author:      thepoy
 * @Email:       thepoy@163.com
 * @File Name:   string.go
 * @Created At:  2021-11-24 19:58:22
 * @Modified At: 2023-03-16 18:35:12
 * @Modified By: thepoy
 */

package tools

import "strings"

// Trim 删除目标字符串两侧的无用字符
func Trim(s string, chars string) string {
	return strings.Trim(s, chars)
}

// Trim 删除两则空白字符，包插空格、换行、制表、全角空格
func Strip(src string) string {
	return Trim(src, " \n\t 　")
}

// ReplaceInvalidSpaces 替换html页面中各种奇葩的空白符为空格
func ReplaceInvalidSpaces(src string) string {
	// A map to store the invalid spaces
	invalidSpaces := map[string]bool{
		"\u00A0": true, // no-break space
		"\u2002": true, // en space
		"\u2003": true, // em space
		"\u2009": true, // thin space
		"\u3000": true, // thin space
	}

	// A variable to store the replaced string
	replacedString := ""

	// Loop through each character of the html and check if it is an invalid space
	for _, c := range src {
		if invalidSpaces[string(c)] {
			replacedString += " " // replace with a normal space
		} else {
			replacedString += string(c) // keep the original character
		}
	}

	return replacedString

}

// RemoveInvalidRunes 删除单句中的无效字符，如一些无意义的特殊字符
func RemoveInvalidRunes(src string, runes ...rune) string {
	var builder strings.Builder

	for _, r := range src {
		if !containsRune(r, runes...) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func containsRune(r rune, runes ...rune) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}
