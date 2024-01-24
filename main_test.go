package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestIPToInt(t *testing.T) {
	ip := "192.168.1.1"
	i := ip2int(ip)
	if i != 3232235777 {
		t.Fatalf(`IP 192.168.1.1 should convert to 3232235777`)
	}
	t.Log("ip2int working fine")
}

func TestIntToIP(t *testing.T) {
	i := 3232235777
	ip := int2ip(uint64(i))
	if ip != "192.168.1.1" {
		t.Fatal(`3232235777 should convert to 192.168.1.1`)
	}
	t.Log("int2ip working fine")
}
