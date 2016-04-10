package lib

import (
	"bytes"
	"testing"
	"time"
)

func TestHotp(t *testing.T) {
	/*
		test vector from RFC 4226 for token key : 0x3132333435363738393031323334353637383930
		count: 0 hex40 (used in ootp): cc93cf1850 HOTP value: (dec6)755224 HOTP value: (dec8)84755224
		count: 1 hex40 (used in ootp): 75a48a19d4 HOTP value: (dec6)287082 HOTP value: (dec8)94287082
		count: 2 hex40 (used in ootp): 0bacb7fa08 HOTP value: (dec6)359152 HOTP value: (dec8)37359152
		count: 3 hex40 (used in ootp): 66c28227d0 HOTP value: (dec6)969429 HOTP value: (dec8)26969429
		count: 4 hex40 (used in ootp): a904c900a6 HOTP value: (dec6)338314 HOTP value: (dec8)40338314
		count: 5 hex40 (used in ootp): a37e783d7b HOTP value: (dec6)254676 HOTP value: (dec8)68254676
		count: 6 hex40 (used in ootp): bc9cd28561 HOTP value: (dec6)287922 HOTP value: (dec8)18287922
		count: 7 hex40 (used in ootp): a4fb960c0b HOTP value: (dec6)162583 HOTP value: (dec8)82162583
		count: 8 hex40 (used in ootp): 1b3c89f65e HOTP value: (dec6)399871 HOTP value: (dec8)73399871
		count: 9 hex40 (used in ootp): 1637409809 HOTP value: (dec6)520489 HOTP value: (dec8)45520489
	*/
	b := hexToBytes("0x3132333435363738393031323334353637383930")
	hotp := NewHotpFromBytes(b)

	if got, want := hotp.At(0), 755224; got != want {
		t.Errorf("want: %v\ngot: %v", want, got)
	}
}

func TestGetHmacSha1Hash(t *testing.T) {
	got := getHmacSha1Hash([]byte("Jefe"), []byte("what do ya want for nothing?"))
	want := hexToBytes("0xeffcdf6ae5eb2fa2d27416d5f184df9c259a7c79")
	if !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
	}

	got = getHmacSha1Hash(
		hexToBytes("0x0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b"),
		[]byte("Hi There"),
	)
	want = hexToBytes("0xb617318655057264e28bc0b6fb378c8ef146be00")
	if !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
	}
}

func TestDynamicTrancation(t *testing.T) {
	in := hexToBytes("0x1f8698690e02ca16618550ef7f19da8e945b555a")
	want := hexToBytes("0x50ef7f19")
	got := dynamicTrancation(in)
	if !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
	}
	sn := bytesToInt(got) % 1000000
	if sn != 872921 {
		t.Errorf("???")
	}
}

func TestGetCounter(t *testing.T) {
	timeNowFunc = func() time.Time {
		t := time.Unix(1408025400, 0)
		return t
	}
	want := int64(1408025400 / 30)
	got := getCounter()
	if got != want {
		t.Errorf("want: %v\ngot: %v", got, want)
	}
}
