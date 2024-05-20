package schema

import (
    "time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// TaskAssignment holds the schema definition for the TaskAssignment entity.
type TaskAssignment struct {
    ent.Schema
}

// Fields of the TaskAssignment.
func (TaskAssignment) Fields() []ent.Field {
    return []ent.Field{
		field.Uint64("id").Positive(),
        field.Uint64("task_id"),
        field.Uint64("user_id"),
		field.Enum("status").Values("pending", "completed").Default("pending"),
        field.Time("assigned_at").Default(time.Now),
    }
}

// Edges of the TaskAssignment.
func (TaskAssignment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("task", Task.Type).
            Ref("assignments").
            Unique().
            Required().
            Field("task_id"),
        edge.From("user", User.Type).
            Ref("assignments").
            Unique().
            Required().
            Field("user_id"),
    }
}