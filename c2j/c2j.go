package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type table struct {
	Columns map[string]column        `json:"columns,omitempty"`
	Rows    []map[string]interface{} `json:"rows"`
}

type column struct {
	Index   int         `json:"index"`
	Group   bool        `json:"group,omitempty"`
	Default interface{} `json:"default,omitempty"`
	Type    string      `json:"type,omitempty"`
}

func main() {
	r := NewReader(os.Stdin)
	var tables []*table
	for r.NextTable() {
		cols := r.Columns()
		var rows []map[string]interface{}
		for r.NextRow() {
			rowMap := make(map[string]interface{})
			for i, val := range r.Row() {
				col := cols[i]
				if val == nil && col.Name == "" {
					continue
				}
				rowMap[col.Name] = val
			}
			rows = append(rows, rowMap)
		}
		columnsMap := make(map[string]column)
		for i, col := range cols {
			if col.Name == "" && col.Default == nil {
				continue
			}
			columnsMap[col.Name] = column{
				Index: i,
				Group: col.Group,
				Type:  col.Type,
			}
		}
		tables = append(tables, &table{
			Rows:    rows,
			Columns: columnsMap,
		})
	}
	if err := r.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	data, err := json.MarshalIndent(tables, "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot marshal JSON: %v\n", err)
		os.Exit(1)
	}
	os.Stdout.Write(data)
	os.Stdout.Write([]byte{'\n'})
}
