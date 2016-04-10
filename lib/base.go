package lib

import (
	"encoding/binary"
	"math"
	"strconv"
	"strings"
)

func hexToBytes(hex string) []byte {
	if strings.Index(hex, "0x") == 0 {
		hex = hex[2:]
	}
	if len(hex)%2 != 0 {
		hex = "0" + hex
	}
	ret := []byte{}
	for i, l := 0, len(hex); i < l; i += 2 {
		h := hex[i : i+2]
		b, _ := strconv.ParseInt(h, 16, 0)
		ret = append(ret, byte(b))
	}
	return ret
}

func bytesToInt(b []byte) int {
	l := len(b)
	ret := 0
	for i := l - 1; i >= 0; i-- {
		base := int(math.Pow(256, float64(l-i-1)))
		ret += int(b[i]) * base
	}
	return ret
}

func intToBytes(i int) []byte {
	ret := make([]byte, 4)
	binary.PutVarint(ret, int64(i))
	return ret
}
