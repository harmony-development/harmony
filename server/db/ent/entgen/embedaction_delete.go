// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/embedaction"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// EmbedActionDelete is the builder for deleting a EmbedAction entity.
type EmbedActionDelete struct {
	config
	hooks    []Hook
	mutation *EmbedActionMutation
}

// Where adds a new predicate to the EmbedActionDelete builder.
func (ead *EmbedActionDelete) Where(ps ...predicate.EmbedAction) *EmbedActionDelete {
	ead.mutation.predicates = append(ead.mutation.predicates, ps...)
	return ead
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ead *EmbedActionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ead.hooks) == 0 {
		affected, err = ead.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmbedActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ead.mutation = mutation
			affected, err = ead.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ead.hooks) - 1; i >= 0; i-- {
			mut = ead.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ead.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ead *EmbedActionDelete) ExecX(ctx context.Context) int {
	n, err := ead.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ead *EmbedActionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: embedaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: embedaction.FieldID,
			},
		},
	}
	if ps := ead.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ead.driver, _spec)
}

// EmbedActionDeleteOne is the builder for deleting a single EmbedAction entity.
type EmbedActionDeleteOne struct {
	ead *EmbedActionDelete
}

// Exec executes the deletion query.
func (eado *EmbedActionDeleteOne) Exec(ctx context.Context) error {
	n, err := eado.ead.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{embedaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eado *EmbedActionDeleteOne) ExecX(ctx context.Context) {
	eado.ead.ExecX(ctx)
}
