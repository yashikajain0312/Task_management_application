// Code generated by ent, DO NOT EDIT.

package taskassignment

import (
	"Tma_backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldLTE(FieldID, id))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldTaskID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldUserID, v))
}

// AssignedAt applies equality check predicate on the "assigned_at" field. It's identical to AssignedAtEQ.
func AssignedAt(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldAssignedAt, v))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNotIn(FieldTaskID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNotIn(FieldUserID, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNotIn(FieldStatus, vs...))
}

// AssignedAtEQ applies the EQ predicate on the "assigned_at" field.
func AssignedAtEQ(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldEQ(FieldAssignedAt, v))
}

// AssignedAtNEQ applies the NEQ predicate on the "assigned_at" field.
func AssignedAtNEQ(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNEQ(FieldAssignedAt, v))
}

// AssignedAtIn applies the In predicate on the "assigned_at" field.
func AssignedAtIn(vs ...time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldIn(FieldAssignedAt, vs...))
}

// AssignedAtNotIn applies the NotIn predicate on the "assigned_at" field.
func AssignedAtNotIn(vs ...time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldNotIn(FieldAssignedAt, vs...))
}

// AssignedAtGT applies the GT predicate on the "assigned_at" field.
func AssignedAtGT(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldGT(FieldAssignedAt, v))
}

// AssignedAtGTE applies the GTE predicate on the "assigned_at" field.
func AssignedAtGTE(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldGTE(FieldAssignedAt, v))
}

// AssignedAtLT applies the LT predicate on the "assigned_at" field.
func AssignedAtLT(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldLT(FieldAssignedAt, v))
}

// AssignedAtLTE applies the LTE predicate on the "assigned_at" field.
func AssignedAtLTE(v time.Time) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.FieldLTE(FieldAssignedAt, v))
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.TaskAssignment {
	return predicate.TaskAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.TaskAssignment {
	return predicate.TaskAssignment(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.TaskAssignment {
	return predicate.TaskAssignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.TaskAssignment {
	return predicate.TaskAssignment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskAssignment) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskAssignment) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TaskAssignment) predicate.TaskAssignment {
	return predicate.TaskAssignment(sql.NotPredicates(p))
}
