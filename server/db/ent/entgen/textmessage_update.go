// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/textmessage"
)

// TextMessageUpdate is the builder for updating TextMessage entities.
type TextMessageUpdate struct {
	config
	hooks    []Hook
	mutation *TextMessageMutation
}

// Where adds a new predicate for the TextMessageUpdate builder.
func (tmu *TextMessageUpdate) Where(ps ...predicate.TextMessage) *TextMessageUpdate {
	tmu.mutation.predicates = append(tmu.mutation.predicates, ps...)
	return tmu
}

// SetContent sets the "content" field.
func (tmu *TextMessageUpdate) SetContent(s string) *TextMessageUpdate {
	tmu.mutation.SetContent(s)
	return tmu
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (tmu *TextMessageUpdate) SetMessageID(id uint64) *TextMessageUpdate {
	tmu.mutation.SetMessageID(id)
	return tmu
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (tmu *TextMessageUpdate) SetNillableMessageID(id *uint64) *TextMessageUpdate {
	if id != nil {
		tmu = tmu.SetMessageID(*id)
	}
	return tmu
}

// SetMessage sets the "message" edge to the Message entity.
func (tmu *TextMessageUpdate) SetMessage(m *Message) *TextMessageUpdate {
	return tmu.SetMessageID(m.ID)
}

// Mutation returns the TextMessageMutation object of the builder.
func (tmu *TextMessageUpdate) Mutation() *TextMessageMutation {
	return tmu.mutation
}

// ClearMessage clears the "message" edge to the Message entity.
func (tmu *TextMessageUpdate) ClearMessage() *TextMessageUpdate {
	tmu.mutation.ClearMessage()
	return tmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmu *TextMessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tmu.hooks) == 0 {
		affected, err = tmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TextMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmu.mutation = mutation
			affected, err = tmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tmu.hooks) - 1; i >= 0; i-- {
			mut = tmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmu *TextMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := tmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmu *TextMessageUpdate) Exec(ctx context.Context) error {
	_, err := tmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmu *TextMessageUpdate) ExecX(ctx context.Context) {
	if err := tmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tmu *TextMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   textmessage.Table,
			Columns: textmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: textmessage.FieldID,
			},
		},
	}
	if ps := tmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textmessage.FieldContent,
		})
	}
	if tmu.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   textmessage.MessageTable,
			Columns: []string{textmessage.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tmu.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   textmessage.MessageTable,
			Columns: []string{textmessage.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{textmessage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TextMessageUpdateOne is the builder for updating a single TextMessage entity.
type TextMessageUpdateOne struct {
	config
	hooks    []Hook
	mutation *TextMessageMutation
}

// SetContent sets the "content" field.
func (tmuo *TextMessageUpdateOne) SetContent(s string) *TextMessageUpdateOne {
	tmuo.mutation.SetContent(s)
	return tmuo
}

// SetMessageID sets the "message" edge to the Message entity by ID.
func (tmuo *TextMessageUpdateOne) SetMessageID(id uint64) *TextMessageUpdateOne {
	tmuo.mutation.SetMessageID(id)
	return tmuo
}

// SetNillableMessageID sets the "message" edge to the Message entity by ID if the given value is not nil.
func (tmuo *TextMessageUpdateOne) SetNillableMessageID(id *uint64) *TextMessageUpdateOne {
	if id != nil {
		tmuo = tmuo.SetMessageID(*id)
	}
	return tmuo
}

// SetMessage sets the "message" edge to the Message entity.
func (tmuo *TextMessageUpdateOne) SetMessage(m *Message) *TextMessageUpdateOne {
	return tmuo.SetMessageID(m.ID)
}

// Mutation returns the TextMessageMutation object of the builder.
func (tmuo *TextMessageUpdateOne) Mutation() *TextMessageMutation {
	return tmuo.mutation
}

// ClearMessage clears the "message" edge to the Message entity.
func (tmuo *TextMessageUpdateOne) ClearMessage() *TextMessageUpdateOne {
	tmuo.mutation.ClearMessage()
	return tmuo
}

// Save executes the query and returns the updated TextMessage entity.
func (tmuo *TextMessageUpdateOne) Save(ctx context.Context) (*TextMessage, error) {
	var (
		err  error
		node *TextMessage
	)
	if len(tmuo.hooks) == 0 {
		node, err = tmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TextMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmuo.mutation = mutation
			node, err = tmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tmuo.hooks) - 1; i >= 0; i-- {
			mut = tmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmuo *TextMessageUpdateOne) SaveX(ctx context.Context) *TextMessage {
	node, err := tmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmuo *TextMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := tmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmuo *TextMessageUpdateOne) ExecX(ctx context.Context) {
	if err := tmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tmuo *TextMessageUpdateOne) sqlSave(ctx context.Context) (_node *TextMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   textmessage.Table,
			Columns: textmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: textmessage.FieldID,
			},
		},
	}
	id, ok := tmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TextMessage.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := tmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: textmessage.FieldContent,
		})
	}
	if tmuo.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   textmessage.MessageTable,
			Columns: []string{textmessage.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tmuo.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   textmessage.MessageTable,
			Columns: []string{textmessage.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TextMessage{config: tmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{textmessage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
