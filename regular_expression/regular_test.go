package regular_expression

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	r := IsMatch("aa", "a")
	if r {
		t.Fatalf("want false, got true")
	}
	r = IsMatch("aa", "aa")
	if !r {
		t.Fatalf("want true, got false")
	}
	r = IsMatch("aaa", "aa")
	if r {
		t.Fatalf("want false, got true")
	}
	r = IsMatch("aa", "a*")
	if !r {
		t.Fatalf("want true, got false")
	}
	r = IsMatch("aa", ".*")
	if !r {
		t.Fatalf("want false, got true")
	}
	r = IsMatch("ab", ".*")
	if !r {
		t.Fatalf("want true, got false")
	}
	r = IsMatch("aab", "c*a*b")
	if !r {
		t.Fatalf("want true, got false")
	}
	r = IsMatch("aa", "a*a")
	if !r {
		t.Fatalf("want true, got false")
	}
}
