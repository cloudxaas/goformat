package cxformat

import (
  "github.com/cloudxaas/gocx"
)

const digits = "0123456789"

func Uint16ToString(num uint16) string {
    if num == 0 {
        return "0"
    }
    buf := [5]byte{} // maximum uint16 digits are 5
    i := len(buf)
    for num > 0 {
        i--
        buf[i] = byte(num%10) + '0'
        num /= 10
    }
    return cx.B2s(buf[i:])
}

func Int16ToString(n int16) string {
    if n == 0 {
        return "0"
    }
    
    var buf [6]byte // maximum length of int16 string is 6 digits
    i := len(buf)
    neg := false
    if n < 0 {
        neg = true
        n = -n
    }
    for n > 0 {
        i--
        buf[i] = digits[int(n%10)]
        n /= 10
    }
    if neg {
        i--
        buf[i] = '-'
    }
    return cx.B2s(buf[i:])
}
