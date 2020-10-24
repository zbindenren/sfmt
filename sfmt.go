// Package sfmt formats slices.
package sfmt

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v2"
)

// RowHeader is an interface to format slices.
type RowHeader interface {
	Header() []string
	Row() []string
}

// SliceWriter writes slices to Writer with a specified Format.
type SliceWriter struct {
	Writer    io.Writer
	NoHeaders bool
}

func (s SliceWriter) Write(f Format, v interface{}) error {
	rows := []RowHeader{}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(v)

		for i := 0; i < slice.Len(); i++ {
			sf, ok := slice.Index(i).Interface().(RowHeader)
			if !ok {
				return errors.New("fuck")
			}

			rows = append(rows, sf)
		}
	default:
		return errors.New("argument is not a slice")
	}

	return s.write(f, rows)
}

// nolint: gocyclo
func (s SliceWriter) write(f Format, r []RowHeader) error {
	if len(r) == 0 {
		return nil
	}

	switch f {
	case JSON:
		enc := json.NewEncoder(s.Writer)
		enc.SetIndent("", "  ")

		if err := enc.Encode(r); err != nil {
			return err
		}

		return nil
	case YAML:
		out, err := yaml.Marshal(r)
		if err != nil {
			return err
		}

		s.Writer.Write(out)

		return nil
	case Table:
		tw := tabwriter.NewWriter(s.Writer, 0, 0, 1, ' ', 0)
		format := strings.TrimRight(strings.Repeat("%s\t", len(r[0].Row())), "\t") + "\n"

		if !s.NoHeaders {
			header := []interface{}{}
			for _, s := range r[0].Header() {
				header = append(header, s)
			}

			fmt.Fprintf(tw, format, header...)
		}

		for i := range r {
			row := []interface{}{}
			for _, s := range r[i].Row() {
				row = append(row, s)
			}

			fmt.Fprintf(tw, format, row...)
		}

		tw.Flush()
	case CSV:
		cw := csv.NewWriter(s.Writer)
		cw.Comma = ';'

		if !s.NoHeaders {
			if err := cw.Write(r[0].Header()); err != nil {
				return err
			}
		}

		for i := range r {
			if err := cw.Write(r[i].Row()); err != nil {
				return err
			}
		}

		cw.Flush()
	default:
		return errors.New("unsupported format")
	}

	return nil
}
