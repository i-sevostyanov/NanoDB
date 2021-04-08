package plan

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

//go:generate mockgen -source=create.go -destination ./create_mock_test.go -package plan_test

type DatabaseCreator interface {
	CreateDatabase(name string) (sql.Database, error)
}

type CreateDatabase struct {
	creator DatabaseCreator
	name    string
}

func NewCreateDatabase(creator DatabaseCreator, name string) *CreateDatabase {
	return &CreateDatabase{
		creator: creator,
		name:    name,
	}
}

func (d *CreateDatabase) RowIter() (sql.RowIter, error) {
	if _, err := d.creator.CreateDatabase(d.name); err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	return sql.RowsIter(), nil
}

type TableCreator interface {
	CreateTable(name string, scheme sql.Scheme) (sql.Table, error)
}

type CreateTable struct {
	creator TableCreator
	name    string
	scheme  sql.Scheme
}

func NewCreateTable(creator TableCreator, name string, scheme sql.Scheme) *CreateTable {
	return &CreateTable{
		creator: creator,
		name:    name,
		scheme:  scheme,
	}
}

func (d *CreateTable) RowIter() (sql.RowIter, error) {
	if _, err := d.creator.CreateTable(d.name, d.scheme); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return sql.RowsIter(), nil
}
