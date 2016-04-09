package lib

import (
	"bytes"
	"testing"
	"time"
)

func TestHotp(t *testing.T) {
	_, err := Hotp("hello", "123")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestGetHmacHash(t *testing.T) {
	got := getHmacHash([]byte("Jefe"), []byte("what do ya want for nothing?"))
	want := hexToBytes("0xeffcdf6ae5eb2fa2d27416d5f184df9c259a7c79")
	if !bytes.Equal(got, want) {
		t.Errorf("want: %v\ngot: %v", want, got)
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
