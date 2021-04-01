package sql

import (
	"io"
)

type Row []Value

// RowIter is an iterator that produces rows.
type RowIter interface {
	// Next retrieves the next row. It will return io.EOF if it's the last row.
	// After retrieving the last row, Close will be automatically closed.
	Next() (Row, error)
	// Close the iterator.
	Close() error
}

type SliceRowsIter struct {
	rows  []Row
	index int
}

func (i *SliceRowsIter) Next() (Row, error) {
	if i.index >= len(i.rows) {
		return nil, io.EOF
	}

	row := i.rows[i.index]
	i.index++

	return row, nil
}

func (i *SliceRowsIter) Close() error {
	i.rows = nil
	return nil
}

func RowsIter(rows ...Row) RowIter {
	return &SliceRowsIter{
		rows: rows,
	}
}
