// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package friend

import ()

type Friend struct {
	Member1ID int64
	Member2ID int64
}

type Member struct {
	ID     int64
	UserID int64
	Name   string
	Status string
}
