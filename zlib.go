/*
 * @Author:    thepoy
 * @Email:     thepoy@163.com
 * @File Name: zlib.go
 * @Created:   2021-11-24 19:58:48
 * @Modified:  2021-11-24 19:59:00
 */

package tools

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zlib"
)

func Compress(src []byte) []byte {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(src)
	w.Close()
	return buf.Bytes()
}

func Decompress(src []byte) ([]byte, error) {
	srcReader := bytes.NewReader(src)

	r, err := zlib.NewReader(srcReader)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()
	return buf.Bytes(), nil
}
