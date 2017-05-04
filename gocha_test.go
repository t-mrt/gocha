package gocha

import (
	"math/rand"
	"regexp"
	"regexp/syntax"
	"testing"
)

func TestNew(t *testing.T) {

	t.Run("New (error)", func(t *testing.T) {
		pattern := `[a-z`

		err, _ := New(pattern)
		_, want := syntax.Parse(pattern, syntax.Perl)

		if err.Error() != want.Error() {
			t.Errorf("%v", err.Error())
		}
	})

	t.Run("New with options", func(t *testing.T) {
		pattern := `[a-z]`
		r := rand.New(rand.NewSource(1))

		err, _ := New(pattern, Rand(r))

		if err != nil {
			t.Errorf("%v", err.Error())
		}
	})
}

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

	t.Run("fixed seed", func(t *testing.T) {
		r := rand.New(rand.NewSource(1))
		_, g := New(`[a-z]{10}`, Rand(r))
		if s := g.Gen(); s != "vbgacrjwtc" {
			t.Errorf("%v must be vbgacrjwtc", s)
		}
		if s := g.Gen(); s != "cahkfapsfc" {
			t.Errorf("%v must be cahkfapsfc", s)
		}
	})
}

var RandFromRange = randFromRange

func TestRandFromRange(t *testing.T) {

	r := rand.New(rand.NewSource(1))

	rs := []intRange{}
	r1 := intRange{
		a: 1,
		b: 2,
	}

	r2 := intRange{
		a: 10,
		b: 11,
	}

	rs = append(rs, r1)
	rs = append(rs, r2)

	if result := randFromRange(rs, r); (result != 1) && (result != 2) && (result != 10) && (result != 11) {
		t.Errorf("result:%v must 1 or 2", result)
	}
}

func BenchmarkGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, g := New(`.{1000}`)
		g.Gen()
	}
}
