package helpers

import (
	"database/sql"
	"encoding/json"
)

type JsonSqlNullString struct {
	sql.NullString
}

func (v JsonSqlNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}