package utils

import "unsafe"

// StringToBytes StringToBytes
/*
 * @description: 字符串零拷贝转切片
 * @version:
 * @auth: mt
 * @param {string} s
 * @return {[]byte}
 */
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesToString BytesToString
/**
 * @description: 切片零拷贝转字符串
 * @version:
 * @auth: mt
 * @param {[]byte} b
 * @return {string}
 */
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
