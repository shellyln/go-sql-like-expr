package likeexpr_test

import (
	"testing"

	"github.com/shellyln/go-sql-like-expr/likeexpr"
)

func Test1(t *testing.T) {
	tests := []struct {
		name         string
		src          string
		escape       rune
		escapeItself bool
		want         string
		wantErr      bool
	}{{
		name:    "1",
		src:     ``,
		escape:  '\\',
		want:    `^$`,
		wantErr: false,
	}, {
		name:    "2",
		src:     `foobar`,
		escape:  '\\',
		want:    `^foobar$`,
		wantErr: false,
	}, {
		name:    "3",
		src:     `foo%`,
		escape:  '\\',
		want:    `^foo.*?$`,
		wantErr: false,
	}, {
		name:    "4",
		src:     `%bar`,
		escape:  '\\',
		want:    `^.*?bar$`,
		wantErr: false,
	}, {
		name:    "5",
		src:     `foob_r`,
		escape:  '\\',
		want:    `^foob.r$`,
		wantErr: false,
	}, {
		name:    "6",
		src:     `foo\%`,
		escape:  '\\',
		want:    `^foo%$`,
		wantErr: false,
	}, {
		name:    "7",
		src:     `foob\%r`,
		escape:  '\\',
		want:    `^foob%r$`,
		wantErr: false,
	}, {
		name:    "8",
		src:     `\%bar`,
		escape:  '\\',
		want:    `^%bar$`,
		wantErr: false,
	}, {
		name:    "9",
		src:     `foob\_r`,
		escape:  '\\',
		want:    `^foob_r$`,
		wantErr: false,
	}, {
		name:    "10",
		src:     `fooba\_`,
		escape:  '\\',
		want:    `^fooba_$`,
		wantErr: false,
	}, {
		name:    "11",
		src:     `\foobar`,
		escape:  '\\',
		want:    `^\\foobar$`,
		wantErr: false,
	}, {
		name:    "12",
		src:     `foo\bar`,
		escape:  '\\',
		want:    `^foo\\bar$`,
		wantErr: false,
	}, {
		name:         "13a",
		src:          `foobar\`,
		escape:       '\\',
		escapeItself: false,
		want:         `^foobar\\$`,
		wantErr:      false,
	}, {
		name:         "13b",
		src:          `foobar\`,
		escape:       '\\',
		escapeItself: true,
		want:         `^foobar\\$`,
		wantErr:      false,
	}, {
		name:         "14a",
		src:          `^(.+?[A-Z]{3,5})(?:\s\d|\\)$`,
		escape:       '\\',
		escapeItself: false,
		want:         `^\^\(\.\+\?\[A-Z\]\{3,5\}\)\(\?:\\s\\d\|\\\\\)\$$`,
		wantErr:      false,
	}, {
		name:         "14b",
		src:          `^(.+?[A-Z]{3,5})(?:\s\d|\\)$`,
		escape:       '\\',
		escapeItself: true,
		want:         `^\^\(\.\+\?\[A-Z\]\{3,5\}\)\(\?:\\s\\d\|\\\)\$$`,
		wantErr:      false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := likeexpr.ToRegexp(tt.src, tt.escape, tt.escapeItself)
			if v != tt.want {
				t.Errorf("%v: v = %v, want = %v", tt.name, v, tt.want)
			}
		})
	}
}
