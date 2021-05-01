package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/shanna/entxid-test/xid"
)

// MixinXID to be shared will all different schemas.
type MixinXID struct {
	mixin.Schema
}

// Fields of the Mixin.
func (MixinXID) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(xid.ID{}).
			DefaultFunc(xid.New).
			Unique().
			Immutable(),
	}
}
