/*
 * @Author:    thepoy
 * @Email:     thepoy@163.com
 * @File Name: encoding.go
 * @Created:   2022-02-18 11:36:56
 * @Modified:  2022-02-18 11:47:43
 */

package tools

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBKToUTF8 converts `GBK` encoded bytes to `UTF-8` encoded bytes
func GBKToUTF8(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder())

	dst, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return dst, err
}

// UTF8ToGBK converts `UTF-8` encoded bytes to `GBK` encoded bytes
func UTF8ToGBK(src []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewEncoder())

	dst, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
