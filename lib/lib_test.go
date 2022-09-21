package lib

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageRangeRegex(t *testing.T) {

	type fields struct {
		c *regexp.Regexp
	}

	type args struct {
		n map[string]bool
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "test",
			prepare: func(f *fields) {
				f.c = PageRangeRegex
			},
			args: args{
				n: map[string]bool{
					"3-4,5-9,-5":             true,
					"3-4,5-9,4-,333d,-":      false,
					"33-45,333":              true,
					"35":                     true,
					"-2":                     true,
					"1-4":                    true,
					"5-":                     true,
					"5-10":                   true,
					"1,1-2,-1,4-,-":          true,
					"1,-2,3-6,5-,-,dfhfjfdj": false,
					"3":                      true,
					"-":                      true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				if got := f.c.MatchString(k); got != v {
					t.Errorf("MatchString = %v, want %v", got, v)
				}
			}
		})
	}
}

func TestParsePageNumber(t *testing.T) {

	type res struct {
		res []Range
		err error
	}

	type fields struct {
		c func(string, int) (*[]Range, error)
	}

	type args struct {
		numPage int
		n       map[string]res
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "test success",
			prepare: func(f *fields) {
				f.c = ParsePageNumber
			},
			args: args{
				n: map[string]res{
					"3-4,5-9,-5,5-": {
						res: []Range{
							{Start: 3, End: 4},
							{Start: 5, End: 9},
							{Start: 1, End: 5},
							{Start: 5, End: 10},
						},
						err: nil,
					},
				},
				numPage: 10,
			},
		},
		{
			name: "test fail with invalid page number",
			prepare: func(f *fields) {
				f.c = ParsePageNumber
			},
			args: args{
				n: map[string]res{
					"3-4,20": {
						res: []Range{},
						err: fmt.Errorf("page range \"20\" is out of total page \"10\""),
					},
				},
				numPage: 10,
			},
		},
		{
			name: "test fail with start > end",
			prepare: func(f *fields) {
				f.c = ParsePageNumber
			},
			args: args{
				n: map[string]res{
					"8-4": {
						res: []Range{},
						err: fmt.Errorf("page range start \"8\" must be less or equal to end \"4\""),
					},
				},
				numPage: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			for k, v := range tt.args.n {
				got, err := f.c(k, tt.args.numPage)
				if err != nil {
					assert.Equal(t, v.err, err)
				} else {
					assert.Equal(t, v.res, *got)
				}
			}
		})
	}
}
