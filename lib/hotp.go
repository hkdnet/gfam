package lib

import (
	"crypto/hmac"
	"crypto/sha1"
	_ "encoding/base32"
	"time"
)

var timeNowFunc = time.Now

func Hotp(secret string, input string) ([]byte, error) {
	return nil, nil
}

func getHmacHash(key, input []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(input)
	return h.Sum(nil)
}

func getCounter() int64 {
	t := timeNowFunc().Unix()
	return t / 30
}

/*
# @param [Integer] input the number used seed the HMAC
# @option padded [Boolean] (false) Output the otp as a 0 padded string
# Usually either the counter, or the computed integer
# based on the Unix timestamp
def generate_otp(input, padded=true)
  hmac = OpenSSL::HMAC.digest(
    OpenSSL::Digest.new(digest),
    byte_secret,
    int_to_bytestring(input)
  )

  offset = hmac[-1].ord & 0xf
  code = (hmac[offset].ord & 0x7f) << 24 |
    (hmac[offset + 1].ord & 0xff) << 16 |
    (hmac[offset + 2].ord & 0xff) << 8 |
    (hmac[offset + 3].ord & 0xff)
  if padded
    (code % 10 ** digits).to_s.rjust(digits, '0')
  else
    code % 10 ** digits
  end
end
*/
