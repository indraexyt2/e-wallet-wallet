package cmd

import (
	"testing"
	"time"
)

func TestWaktu(t *testing.T) {
	waktuSekarang := time.Now()
	waktuRFC3339 := waktuSekarang.Format(time.RFC3339)
	println(waktuRFC3339)
}
