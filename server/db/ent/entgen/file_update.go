// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/harmony-development/legato/server/db/ent/entgen/file"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where adds a new predicate for the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.predicates = append(fu.mutation.predicates, ps...)
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetContenttype sets the "contenttype" field.
func (fu *FileUpdate) SetContenttype(s string) *FileUpdate {
	fu.mutation.SetContenttype(s)
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: file.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldName,
		})
	}
	if value, ok := fu.mutation.Contenttype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldContenttype,
		})
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetContenttype sets the "contenttype" field.
func (fuo *FileUpdateOne) SetContenttype(s string) *FileUpdateOne {
	fuo.mutation.SetContenttype(s)
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: file.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing File.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldName,
		})
	}
	if value, ok := fuo.mutation.Contenttype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldContenttype,
		})
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}