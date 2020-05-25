package urlUtil

import (
	"testing"
)

func TestCombind(t *testing.T) {
	var actual string
	actual = Combind("http://ip:port/http", "add/user")
	if actual != "http://ip:port/http/add/user" {
		t.FailNow()
	}
	actual = Combind("http://ip:port/http", "/add/user")
	if actual != "http://ip:port/http/add/user" {
		t.FailNow()
	}
	actual = Combind("http://ip:port/http/", "add/user")
	if actual != "http://ip:port/http/add/user" {
		t.FailNow()
	}
	actual = Combind("http://ip:port/http/", "/add/user")
	if actual != "http://ip:port/http/add/user" {
		t.FailNow()
	}
	actual = Combind("http://ip:port/http//", "//add/user")
	if actual != "http://ip:port/http/add/user" {
		t.FailNow()
	}
}
