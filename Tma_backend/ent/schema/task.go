package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Task holds the schema definition for the Task entity.
type Task struct {
    ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
    return []ent.Field{
		field.Uint64("id").Positive(),
        field.String("title").NotEmpty(),
        field.String("description").NotEmpty(),
        field.Time("due_date").Default(time.Now),
		field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("assignments", TaskAssignment.Type),
    }
}
