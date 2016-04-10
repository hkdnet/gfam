package lib

import (
	"crypto/hmac"
	"crypto/sha1"
	"time"
)

var timeNowFunc = time.Now

type Hotp struct {
	key []byte
}

func NewHotpFromString(str string) *Hotp {
	ret := &Hotp{}
	ret.key = []byte(str)
	return ret
}

func NewHotpFromBytes(bytes []byte) *Hotp {
	ret := &Hotp{}
	ret.key = bytes
	return ret
}

func (h *Hotp) At(counter int) int {
	hs := h.hash(counter)
	sbits := dynamicTrancation(hs)
	snum := bytesToInt(sbits)
	snum = snum % 1000000
	return snum
}

func (h *Hotp) hash(counter int) []byte {
	return getHmacSha1Hash(h.key, []byte(string(counter)))
}

func getHmacSha1Hash(key, input []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(input)
	return h.Sum(nil)
}

func dynamicTrancation(b []byte) []byte {
	/*
	   DT(String) // String = String[0]...String[19]
	   Let OffsetBits be the low-order 4 bits of String[19]
	   Offset = StToNum(OffsetBits) // 0 <= OffSet <= 15
	   Let P = String[OffSet]...String[OffSet+3]
	   Return the Last 31 bits of P
	*/
	last := b[19]
	offset := last % 16
	ret := b[offset : offset+4]
	ret[0] = ret[0] % 127
	return ret
	/*
		p := b[offset : offset+3]
		ret := ""
		for e := range p {
			ret += fmt.Sprintf("%b", e)
		}
		return ret
	*/
}

func getCounter() int64 {
	t := timeNowFunc().Unix()
	return t / 30
}
