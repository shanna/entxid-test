package xid

import (
	"io"
	"strconv"

	"database/sql"
	"database/sql/driver"

	rsxid "github.com/rs/xid"
)

// github.com/rs/xid currently implements these interfaces but just to be sure it remains that way:
var _ sql.Scanner = (*ID)(nil)
var _ driver.Valuer = (*ID)(nil)

type ID struct {
	rsxid.ID
}

func New() ID {
	return ID{rsxid.New()}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(u.String()))
}
