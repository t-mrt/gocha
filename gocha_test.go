package gocha

import (
	"regexp"
	"testing"
)

func TestGen(t *testing.T) {

	patterns := []string{
		`a`,
		`ab`,
		`a|b`,
		`a*`,
		`a?`,
		`a+`,
		`a{1,3}`,
		`a{3}`,
		`a{3,}`,
		`[xyz]`,
		`[^xyz]`,
		`[[:alpha:]]`,
		`[[:^alpha:]]`,
		`\pN`,
		`\p{Greek}`,
		`\PN`,
		`\P{Greek}`,
		`x*?`,
		`x+?`,
		`x??`,
		`x{n,m}?`,
		`x{n,}?`,
		`x{n}?`,
		`(re)`,
		`\d`,
		`\D`,
		`.`,
		`[カコヵか][ッー]{1,3}?[フヒふひ]{1,3}[ィェー]{1,3}[ズス][ドクグュ][リイ][プブぷぶ]{1,3}[トドォ]{1,2}`,
		`[あ-お]{10}`,
		`(?i:[a-z]{10})`,
		`$`,
		`(?i)[^\W]`,
		`[[:alpha:]]`,
	}

	for _, pattern := range patterns {

		_, g := New(pattern)
		s := g.Gen()

		if m, _ := regexp.MatchString(pattern, s); !m {
			t.Errorf("%v does not match to %v", s, pattern)
		}
	}

	_, g := New(``)
	s := g.Gen()

	if m, _ := regexp.MatchString(`\A\z`, s); !m {
		t.Errorf("null regexp")
	}
}
