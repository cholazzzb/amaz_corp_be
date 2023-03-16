package database

import _ "embed"

//go:embed user/schema.sql
var DdlUser string
