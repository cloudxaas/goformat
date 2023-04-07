package cxformat

import (
	cx "github.com/cloudxaas/gocx"
	"github.com/valyala/bytebufferpool"
)

const digits = "0123456789"

func Uint16ToString(num uint16) string {
    if num == 0 {
        return "0"
    }
    buf := bytebufferpool.Get() // get buffer from pool
    defer bytebufferpool.Put(buf) // return buffer to pool
    for num > 0 {
        buf.WriteByte(byte(num%10) + '0')
        num /= 10
    }
    return cx.B2s(buf.B)
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
    bb := bytebufferpool.Get() // get buffer from pool
    defer bytebufferpool.Put(bb) // return buffer to pool
    for n > 0 {
        i--
        bb.WriteByte(digits[int(n%10)])
        n /= 10
    }
    if neg {
        i--
        bb.WriteByte('-')
    }
    return cx.B2s(bb.Bytes())
}