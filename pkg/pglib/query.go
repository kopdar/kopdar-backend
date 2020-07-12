package pglib

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

// List of error values.
var (
	ErrNotFound = errors.New("data not found")
)

const (
	ptrStruct = iota
	ptrSlicePtrStruct
)

func Query(ctx context.Context, db *sqlx.DB, query string, params interface{}, out interface{}) error {
	kind, err := findKind(out)
	if err != nil {
		return err
	}

	newQuery, args, err := db.BindNamed(query, params)
	if err != nil {
		return err
	}

	rows, err := db.QueryxContext(ctx, newQuery, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	switch kind {
	case ptrStruct:
		if !rows.Next() {
			return ErrNotFound
		}
		return rows.StructScan(out)
	case ptrSlicePtrStruct:
		rval := reflect.ValueOf(out).Elem()
		rval.Set(reflect.Zero(rval.Type()))
		for rows.Next() {
			n := rval.Len()
			rval.Set(reflect.Append(rval, reflect.New(rval.Type().Elem().Elem())))
			if err := rows.StructScan(rval.Index(n).Interface()); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}


func findKind(val interface{}) (int, error) {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Ptr {
		return 0, fmt.Errorf("invalid type %T - must be pointer", val)
	}

	t = t.Elem()
	kind := ptrStruct
	if t.Kind() == reflect.Slice {
		t = t.Elem()
		if t.Kind() != reflect.Ptr {
			return 0, fmt.Errorf("invalid type %T - elem of slice must be pointer", val)
		}

		t = t.Elem()
		kind = ptrSlicePtrStruct
	}

	if t.Kind() != reflect.Struct {
		return 0, fmt.Errorf("invalid type %T - %s should be struct", val, t.String())
	}

	return kind, nil
}
