package lib

import (
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
