package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Column struct {
	Name    string
	Group   bool
	Default interface{}
	Type    string
}

func NewReader(r io.Reader) *Reader {
	r1 := &Reader{
		r: csv.NewReader(r),
	}
	r1.r.FieldsPerRecord = -1
	return r1
}

type Reader struct {
	cols []Column
	row  []interface{}
	err  error

	hasPeeked bool
	peekRow   []string
	peekErr   error
	r         *csv.Reader
	line      int
}

// NextTable advances to the next table and reports whether
// there is one.
func (r *Reader) NextTable() bool {
	if r.err != nil {
		return false
	}
	// Read all the current rows.
	for r.NextRow() {
	}
	_, err := r.peek()
	if err != nil {
		r.err = err
		return false
	}
	cols, err := r.readHeader()
	if err != nil {
		r.err = err
		return false
	}
	r.cols = cols
	return true
}

// Err returns any error encountered when parsing.
func (r *Reader) Err() error {
	if r.err == io.EOF {
		return nil
	}
	return r.err
}

// Columns returns the columns in the current table.
func (r *Reader) Columns() []Column {
	return r.cols
}

// NextRow advances to the next row in the current table
// and reports whether there is one.
func (r *Reader) NextRow() bool {
	if r.cols == nil || r.err != nil {
		return false
	}
	row, err := r.readRow()
	r.row = row
	if row == nil {
		r.err = err
		r.cols = nil
		return false
	}
	return true
}

// Row returns the items in the current row of the current table.
func (r *Reader) Row() []interface{} {
	return r.row
}

func (r *Reader) readRow() ([]interface{}, error) {
	row, err := r.peek()
	if err != nil {
		return nil, nil
	}
	if len(row) > 0 && strings.HasPrefix(row[0], "#") {
		// Start of next table.
		return nil, nil
	}
	r.read()
	if len(row) != len(r.cols) {
		return nil, fmt.Errorf("inconsistent number of columns at line %d; got %d want %d", r.line, len(row), len(r.cols))
	}
	rowVals := make([]interface{}, len(row))
	for i, val := range row {
		col := r.cols[i]
		if col.Default != nil && val == "" {
			rowVals[i] = col.Default
			continue
		}
		if val == "" && col.Name == "" {
			continue
		}
		x, err := convertToType(val, col.Type)
		if err != nil {
			return nil, fmt.Errorf("cannot parse %q as type %q at line %d", val, col.Type, r.line)
		}
		rowVals[i] = x
	}
	return rowVals, nil
}

func (r *Reader) readHeader() ([]Column, error) {
	var cols []Column
	var defaults []string
	for {
		row, err := r.peek()
		if err != nil {
			if len(cols) > 0 {
				err = nil
			}
			return cols, err
		}
		r.read()
		if cols == nil {
			if len(row) == 0 {
				return nil, fmt.Errorf("no columns in table header at line %d", r.line)
			}
			cols = make([]Column, len(row))
		} else if len(row) != len(cols) {
			return nil, fmt.Errorf("inconsistent table header (got %d items want %d)", len(row), len(cols))
		}
		if !strings.HasPrefix(row[0], "#") {
			for i, col := range row {
				cols[i].Name = col
			}
			break
		}
		switch row[0] {
		case "#datatype":
			for i := 1; i < len(row); i++ {
				cols[i].Type = row[i]
			}
		case "#group":
			for i := 1; i < len(row); i++ {
				// TODO parse bool?
				cols[i].Group = row[i] == "true"
			}
		case "#default":
			defaults = row
		default:
			fmt.Fprintf(os.Stderr, "unknown column annotation %q\n", row[0])
		}
	}
	if defaults != nil {
		for i := 1; i < len(defaults); i++ {
			if defaults[i] == "" {
				continue
			}
			x, err := convertToType(defaults[i], cols[i].Type)
			if err != nil {
				return nil, fmt.Errorf("cannot convert default value %q to type %q: %v", defaults[i], cols[i].Type, err)
			}
			cols[i].Default = x
		}
	}
	return cols, nil
}

func convertToType(s string, typ string) (interface{}, error) {
	switch typ {
	case "boolean":
		return strconv.ParseBool(s)
	case "long":
		return strconv.ParseInt(s, 10, 64)
	case "unsignedLong":
		return strconv.ParseUint(s, 10, 64)
	case "double":
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		if math.IsInf(x, 0) || math.IsNaN(x) {
			return s, nil
		}
		return x, nil
	case "string", "tag", "":
		return s, nil
	}
	if timeFormat := strings.TrimPrefix(typ, "dateTime:"); len(timeFormat) != len(typ) {
		layout := timeFormats[timeFormat]
		if layout == "" {
			return nil, fmt.Errorf("unknown time format %q", typ)
		}
		return time.Parse(layout, s)
	}
	fmt.Fprintf(os.Stderr, "unknown datatype %q\n", typ)
	return s, nil
}

var timeFormats = map[string]string{
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
}

func (r *Reader) read() (_r []string, _err error) {
	if r.hasPeeked {
		row, err := r.peekRow, r.peekErr
		r.peekRow, r.peekErr, r.hasPeeked = nil, nil, false
		return row, err
	}
	r.line++
	return r.r.Read()
}

func (r *Reader) peek() (_r []string, _err error) {
	if r.hasPeeked {
		return r.peekRow, r.peekErr
	}
	row, err := r.r.Read()
	r.line++
	r.peekRow, r.peekErr, r.hasPeeked = row, err, true
	return row, err
}
