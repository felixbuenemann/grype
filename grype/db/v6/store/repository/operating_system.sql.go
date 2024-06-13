// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: operating_system.sql

package repository

import (
	"context"
	"database/sql"
)

const createOperatingSystem = `-- name: CreateOperatingSystem :one
INSERT INTO operating_systems (name, major_version, minor_version, codename)
VALUES (?, ?, ?, ?)
RETURNING os_id
`

type CreateOperatingSystemParams struct {
	Name         string         `db:"name" json:"name"`
	MajorVersion string         `db:"major_version" json:"majorVersion"`
	MinorVersion sql.NullString `db:"minor_version" json:"minorVersion"`
	Codename     sql.NullString `db:"codename" json:"codename"`
}

func (q *Queries) CreateOperatingSystem(ctx context.Context, arg CreateOperatingSystemParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createOperatingSystem,
		arg.Name,
		arg.MajorVersion,
		arg.MinorVersion,
		arg.Codename,
	)
	var os_id int64
	err := row.Scan(&os_id)
	return os_id, err
}

const listOperatingSystems = `-- name: ListOperatingSystems :many

SELECT os_id, name, major_version, minor_version, codename FROM operating_systems
ORDER BY name, major_version, minor_version
`

// TODO: this ends up creating interface{} values for input params, which is not great at runtime safety-wise... but this is a single flex statement, replacing all of the other ones below
// -- name: ListOperatingSystems :many
// SELECT * FROM operating_systems
// WHERE (name = COALESCE(NULLIF(@name, ”), name))
//
//	AND (major_version = COALESCE(NULLIF(@major_version, ''), major_version))
//	AND (minor_version = COALESCE(?, minor_version))
//	AND (codename = COALESCE(?, codename))
//
// ORDER BY name, major_version, minor_version;
func (q *Queries) ListOperatingSystems(ctx context.Context) ([]OperatingSystem, error) {
	rows, err := q.db.QueryContext(ctx, listOperatingSystems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OperatingSystem{}
	for rows.Next() {
		var i OperatingSystem
		if err := rows.Scan(
			&i.OsID,
			&i.Name,
			&i.MajorVersion,
			&i.MinorVersion,
			&i.Codename,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOperatingSystemsByCodename = `-- name: ListOperatingSystemsByCodename :many
SELECT os_id, name, major_version, minor_version, codename FROM operating_systems
WHERE codename = ?
ORDER BY name, major_version, minor_version
`

func (q *Queries) ListOperatingSystemsByCodename(ctx context.Context, codename sql.NullString) ([]OperatingSystem, error) {
	rows, err := q.db.QueryContext(ctx, listOperatingSystemsByCodename, codename)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OperatingSystem{}
	for rows.Next() {
		var i OperatingSystem
		if err := rows.Scan(
			&i.OsID,
			&i.Name,
			&i.MajorVersion,
			&i.MinorVersion,
			&i.Codename,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOperatingSystemsByName = `-- name: ListOperatingSystemsByName :many
SELECT os_id, name, major_version, minor_version, codename FROM operating_systems
WHERE name = ?
ORDER BY name, major_version, minor_version
`

func (q *Queries) ListOperatingSystemsByName(ctx context.Context, name string) ([]OperatingSystem, error) {
	rows, err := q.db.QueryContext(ctx, listOperatingSystemsByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OperatingSystem{}
	for rows.Next() {
		var i OperatingSystem
		if err := rows.Scan(
			&i.OsID,
			&i.Name,
			&i.MajorVersion,
			&i.MinorVersion,
			&i.Codename,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOperatingSystemsByNameAndExactVersion = `-- name: ListOperatingSystemsByNameAndExactVersion :many
SELECT os_id, name, major_version, minor_version, codename FROM operating_systems
WHERE name = ?
  AND major_version = ?
  AND minor_version = ?
ORDER BY name, major_version, minor_version
`

type ListOperatingSystemsByNameAndExactVersionParams struct {
	Name         string         `db:"name" json:"name"`
	MajorVersion string         `db:"major_version" json:"majorVersion"`
	MinorVersion sql.NullString `db:"minor_version" json:"minorVersion"`
}

func (q *Queries) ListOperatingSystemsByNameAndExactVersion(ctx context.Context, arg ListOperatingSystemsByNameAndExactVersionParams) ([]OperatingSystem, error) {
	rows, err := q.db.QueryContext(ctx, listOperatingSystemsByNameAndExactVersion, arg.Name, arg.MajorVersion, arg.MinorVersion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OperatingSystem{}
	for rows.Next() {
		var i OperatingSystem
		if err := rows.Scan(
			&i.OsID,
			&i.Name,
			&i.MajorVersion,
			&i.MinorVersion,
			&i.Codename,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOperatingSystemsByNameAndMajorVersion = `-- name: ListOperatingSystemsByNameAndMajorVersion :many
SELECT os_id, name, major_version, minor_version, codename FROM operating_systems
WHERE name = ?
  AND major_version = ?
ORDER BY name, major_version, minor_version
`

type ListOperatingSystemsByNameAndMajorVersionParams struct {
	Name         string `db:"name" json:"name"`
	MajorVersion string `db:"major_version" json:"majorVersion"`
}

func (q *Queries) ListOperatingSystemsByNameAndMajorVersion(ctx context.Context, arg ListOperatingSystemsByNameAndMajorVersionParams) ([]OperatingSystem, error) {
	rows, err := q.db.QueryContext(ctx, listOperatingSystemsByNameAndMajorVersion, arg.Name, arg.MajorVersion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OperatingSystem{}
	for rows.Next() {
		var i OperatingSystem
		if err := rows.Scan(
			&i.OsID,
			&i.Name,
			&i.MajorVersion,
			&i.MinorVersion,
			&i.Codename,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
