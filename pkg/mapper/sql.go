package mapper

import (
	"database/sql"
	"time"
)

func ToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ToInt64(ni sql.NullInt64) int64 {
	if ni.Valid {
		return ni.Int64
	}
	return 0
}

func ToBool(nb sql.NullBool) bool {
	if nb.Valid {
		return nb.Bool
	}
	return false
}

func ToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: i != 0,
	}
}

func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: true, // usually you always want to store a bool
	}
}

func ToNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
