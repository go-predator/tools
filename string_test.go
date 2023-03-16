/**
 * @Author:      thepoy
 * @Email:       thepoy@163.com
 * @File Name:   string_test.go
 * @Created At:  2023-03-16 18:21:48
 * @Modified At: 2023-03-16 18:25:19
 * @Modified By: thepoy
 */

package tools

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrip(t *testing.T) {
	Convey("test strip", t, func() {
		s := " \t\n　a test string\t\t\t   　　\n"
		So(Strip(s), ShouldEqual, "a test string")
	})
}
