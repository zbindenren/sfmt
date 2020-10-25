package sfmt_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zbindenren/sfmt"
)

type SkiCompany struct {
	Name        string
	Country     string
	Established int
}

func (s SkiCompany) Header() []string {
	return []string{"NAME", "COUNTRY", "ESTABLISHED"}
}

func (s SkiCompany) Row() []string {
	return []string{s.Name, s.Country, fmt.Sprintf("%d", s.Established)}
}

var (
	companies = []SkiCompany{
		{
			Name:        "Black Crows",
			Country:     "France",
			Established: 2006,
		},
		{
			Name:        "Faction Skis",
			Country:     "Switzerland",
			Established: 2006,
		},
		{
			Name:        "Moment Skis",
			Country:     "USA",
			Established: 2003,
		},
	}

	expectedCSV = `NAME;COUNTRY;ESTABLISHED
Black Crows;France;2006
Faction Skis;Switzerland;2006
Moment Skis;USA;2003
`

	expectedCSVNoHeader = `Black Crows;France;2006
Faction Skis;Switzerland;2006
Moment Skis;USA;2003
`
	expectedTable = `NAME         COUNTRY     ESTABLISHED
Black Crows  France      2006
Faction Skis Switzerland 2006
Moment Skis  USA         2003
`

	expectedTableNoHeader = `Black Crows  France      2006
Faction Skis Switzerland 2006
Moment Skis  USA         2003
`

	expectedJSON = `[
  {
    "Name": "Black Crows",
    "Country": "France",
    "Established": 2006
  },
  {
    "Name": "Faction Skis",
    "Country": "Switzerland",
    "Established": 2006
  },
  {
    "Name": "Moment Skis",
    "Country": "USA",
    "Established": 2003
  }
]
`

	expectedYAML = `- name: Black Crows
  country: France
  established: 2006
- name: Faction Skis
  country: Switzerland
  established: 2006
- name: Moment Skis
  country: USA
  established: 2003
`
)

func TestSliceWriter(t *testing.T) {
	var withHeader = []struct {
		name     string
		format   sfmt.Format
		expected string
	}{
		{
			"csv",
			sfmt.CSV,
			expectedCSV,
		},
		{
			"table",
			sfmt.Table,
			expectedTable,
		},
		{
			"yaml",
			sfmt.YAML,
			expectedYAML,
		},
		{
			"json",
			sfmt.JSON,
			expectedJSON,
		},
	}

	for i := range withHeader {
		tc := withHeader[i]
		b := bytes.NewBufferString("")

		t.Run(tc.name, func(t *testing.T) {
			s := sfmt.SliceWriter{
				Writer: b,
			}

			s.Write(tc.format, companies)
			assert.Equal(t, tc.expected, b.String())
		})
	}

	var noHeader = []struct {
		name     string
		format   sfmt.Format
		expected string
	}{
		{
			"csv",
			sfmt.CSV,
			expectedCSVNoHeader,
		},
		{
			"table",
			sfmt.Table,
			expectedTableNoHeader,
		},
	}

	for i := range noHeader {
		tc := noHeader[i]
		b := bytes.NewBufferString("")

		t.Run(tc.name, func(t *testing.T) {
			s := sfmt.SliceWriter{
				Writer:    b,
				NoHeaders: true,
			}

			s.Write(tc.format, companies)
			assert.Equal(t, tc.expected, b.String())
		})
	}
}
