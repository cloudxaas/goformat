package goformat

import (
  "github.com/cloudxaas/gocx"
)

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
