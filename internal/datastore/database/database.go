package database

import _ "embed"

//go:embed user/schema.sql
var DdlUser string

//go:embed friend/schema.sql
var DdlFriend string
