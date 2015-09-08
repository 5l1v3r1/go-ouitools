// Package go-oui provides functions to work with MAC and OUI's
package ouidb

import (
	"fmt"
	"testing"
)

var db OuiDb

func init() {
	db = OuiDb{}
	err := db.Load("oui.txt")
	if err != nil {
		panic(err.Error())
	}
}

func lookup(t *testing.T, mac, org string) {
	v, err := db.VendorLookup(mac)
	if err != nil {
		t.Fatalf("parse: %s: %s", mac, err.Error())
	}
	if v != org {
		t.Fatalf("lookup: input %s, expect %s, got %s", mac, org, v)
	}
	fmt.Printf("    %s => %s\n", mac, v)
}

func TestLookup1(t *testing.T) {
	lookup(t, "60:03:08:a0:ec:a6", "Apple, Inc.")
}

func TestLookup2(t *testing.T) {
	lookup(t, "00:25:9c:42:c2:62", "Cisco-Linksys, LLC")
}

func TestLookup3(t *testing.T) {
	lookup(t, "00:16:e0:3d:f4:4c", "3Com Ltd")
}
