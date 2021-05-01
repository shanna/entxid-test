// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
