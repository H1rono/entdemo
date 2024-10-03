package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"regexp"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	nameRegexp := regexp.MustCompile("[a-zA-Z_]+$")
	return []ent.Field{
		field.String("name").Match(nameRegexp),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
