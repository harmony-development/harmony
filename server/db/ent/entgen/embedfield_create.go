// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/embedaction"
	"github.com/harmony-development/legato/server/db/ent/entgen/embedfield"
	"github.com/harmony-development/legato/server/db/ent/entgen/embedmessage"
)

// EmbedFieldCreate is the builder for creating a EmbedField entity.
type EmbedFieldCreate struct {
	config
	mutation *EmbedFieldMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (efc *EmbedFieldCreate) SetTitle(s string) *EmbedFieldCreate {
	efc.mutation.SetTitle(s)
	return efc
}

// SetSubtitle sets the "subtitle" field.
func (efc *EmbedFieldCreate) SetSubtitle(s string) *EmbedFieldCreate {
	efc.mutation.SetSubtitle(s)
	return efc
}

// SetBody sets the "body" field.
func (efc *EmbedFieldCreate) SetBody(s string) *EmbedFieldCreate {
	efc.mutation.SetBody(s)
	return efc
}

// SetImageURL sets the "image_url" field.
func (efc *EmbedFieldCreate) SetImageURL(s string) *EmbedFieldCreate {
	efc.mutation.SetImageURL(s)
	return efc
}

// SetPresentation sets the "presentation" field.
func (efc *EmbedFieldCreate) SetPresentation(i int8) *EmbedFieldCreate {
	efc.mutation.SetPresentation(i)
	return efc
}

// AddEmbedActionIDs adds the "embed_action" edge to the EmbedAction entity by IDs.
func (efc *EmbedFieldCreate) AddEmbedActionIDs(ids ...int) *EmbedFieldCreate {
	efc.mutation.AddEmbedActionIDs(ids...)
	return efc
}

// AddEmbedAction adds the "embed_action" edges to the EmbedAction entity.
func (efc *EmbedFieldCreate) AddEmbedAction(e ...*EmbedAction) *EmbedFieldCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return efc.AddEmbedActionIDs(ids...)
}

// SetEmbedMessageID sets the "embed_message" edge to the EmbedMessage entity by ID.
func (efc *EmbedFieldCreate) SetEmbedMessageID(id int) *EmbedFieldCreate {
	efc.mutation.SetEmbedMessageID(id)
	return efc
}

// SetNillableEmbedMessageID sets the "embed_message" edge to the EmbedMessage entity by ID if the given value is not nil.
func (efc *EmbedFieldCreate) SetNillableEmbedMessageID(id *int) *EmbedFieldCreate {
	if id != nil {
		efc = efc.SetEmbedMessageID(*id)
	}
	return efc
}

// SetEmbedMessage sets the "embed_message" edge to the EmbedMessage entity.
func (efc *EmbedFieldCreate) SetEmbedMessage(e *EmbedMessage) *EmbedFieldCreate {
	return efc.SetEmbedMessageID(e.ID)
}

// Mutation returns the EmbedFieldMutation object of the builder.
func (efc *EmbedFieldCreate) Mutation() *EmbedFieldMutation {
	return efc.mutation
}

// Save creates the EmbedField in the database.
func (efc *EmbedFieldCreate) Save(ctx context.Context) (*EmbedField, error) {
	var (
		err  error
		node *EmbedField
	)
	if len(efc.hooks) == 0 {
		if err = efc.check(); err != nil {
			return nil, err
		}
		node, err = efc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmbedFieldMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = efc.check(); err != nil {
				return nil, err
			}
			efc.mutation = mutation
			node, err = efc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(efc.hooks) - 1; i >= 0; i-- {
			mut = efc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, efc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (efc *EmbedFieldCreate) SaveX(ctx context.Context) *EmbedField {
	v, err := efc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (efc *EmbedFieldCreate) check() error {
	if _, ok := efc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("entgen: missing required field \"title\"")}
	}
	if _, ok := efc.mutation.Subtitle(); !ok {
		return &ValidationError{Name: "subtitle", err: errors.New("entgen: missing required field \"subtitle\"")}
	}
	if _, ok := efc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New("entgen: missing required field \"body\"")}
	}
	if _, ok := efc.mutation.ImageURL(); !ok {
		return &ValidationError{Name: "image_url", err: errors.New("entgen: missing required field \"image_url\"")}
	}
	if _, ok := efc.mutation.Presentation(); !ok {
		return &ValidationError{Name: "presentation", err: errors.New("entgen: missing required field \"presentation\"")}
	}
	return nil
}

func (efc *EmbedFieldCreate) sqlSave(ctx context.Context) (*EmbedField, error) {
	_node, _spec := efc.createSpec()
	if err := sqlgraph.CreateNode(ctx, efc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (efc *EmbedFieldCreate) createSpec() (*EmbedField, *sqlgraph.CreateSpec) {
	var (
		_node = &EmbedField{config: efc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: embedfield.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: embedfield.FieldID,
			},
		}
	)
	if value, ok := efc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: embedfield.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := efc.mutation.Subtitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: embedfield.FieldSubtitle,
		})
		_node.Subtitle = value
	}
	if value, ok := efc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: embedfield.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := efc.mutation.ImageURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: embedfield.FieldImageURL,
		})
		_node.ImageURL = value
	}
	if value, ok := efc.mutation.Presentation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: embedfield.FieldPresentation,
		})
		_node.Presentation = value
	}
	if nodes := efc.mutation.EmbedActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   embedfield.EmbedActionTable,
			Columns: []string{embedfield.EmbedActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: embedaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := efc.mutation.EmbedMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   embedfield.EmbedMessageTable,
			Columns: []string{embedfield.EmbedMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: embedmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.embed_message_embed_field = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmbedFieldCreateBulk is the builder for creating many EmbedField entities in bulk.
type EmbedFieldCreateBulk struct {
	config
	builders []*EmbedFieldCreate
}

// Save creates the EmbedField entities in the database.
func (efcb *EmbedFieldCreateBulk) Save(ctx context.Context) ([]*EmbedField, error) {
	specs := make([]*sqlgraph.CreateSpec, len(efcb.builders))
	nodes := make([]*EmbedField, len(efcb.builders))
	mutators := make([]Mutator, len(efcb.builders))
	for i := range efcb.builders {
		func(i int, root context.Context) {
			builder := efcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmbedFieldMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, efcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, efcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, efcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (efcb *EmbedFieldCreateBulk) SaveX(ctx context.Context) []*EmbedField {
	v, err := efcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
