// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	v1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/ent/entgen/channel"
	"github.com/harmony-development/legato/server/db/ent/entgen/message"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
	"github.com/harmony-development/legato/server/db/ent/entgen/user"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where adds a new predicate for the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.predicates = append(mu.mutation.predicates, ps...)
	return mu
}

// SetCreatedat sets the "createdat" field.
func (mu *MessageUpdate) SetCreatedat(t time.Time) *MessageUpdate {
	mu.mutation.SetCreatedat(t)
	return mu
}

// SetNillableCreatedat sets the "createdat" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableCreatedat(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetCreatedat(*t)
	}
	return mu
}

// SetEditedat sets the "editedat" field.
func (mu *MessageUpdate) SetEditedat(t time.Time) *MessageUpdate {
	mu.mutation.SetEditedat(t)
	return mu
}

// SetNillableEditedat sets the "editedat" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableEditedat(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetEditedat(*t)
	}
	return mu
}

// ClearEditedat clears the value of the "editedat" field.
func (mu *MessageUpdate) ClearEditedat() *MessageUpdate {
	mu.mutation.ClearEditedat()
	return mu
}

// SetMetadata sets the "metadata" field.
func (mu *MessageUpdate) SetMetadata(v *v1.Metadata) *MessageUpdate {
	mu.mutation.SetMetadata(v)
	return mu
}

// ClearMetadata clears the value of the "metadata" field.
func (mu *MessageUpdate) ClearMetadata() *MessageUpdate {
	mu.mutation.ClearMetadata()
	return mu
}

// SetOverride sets the "override" field.
func (mu *MessageUpdate) SetOverride(v *v1.Override) *MessageUpdate {
	mu.mutation.SetOverride(v)
	return mu
}

// ClearOverride clears the value of the "override" field.
func (mu *MessageUpdate) ClearOverride() *MessageUpdate {
	mu.mutation.ClearOverride()
	return mu
}

// SetContent sets the "Content" field.
func (mu *MessageUpdate) SetContent(v *v1.Content) *MessageUpdate {
	mu.mutation.SetContent(v)
	return mu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mu *MessageUpdate) SetUserID(id uint64) *MessageUpdate {
	mu.mutation.SetUserID(id)
	return mu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableUserID(id *uint64) *MessageUpdate {
	if id != nil {
		mu = mu.SetUserID(*id)
	}
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MessageUpdate) SetUser(u *User) *MessageUpdate {
	return mu.SetUserID(u.ID)
}

// SetChannelID sets the "channel" edge to the Channel entity by ID.
func (mu *MessageUpdate) SetChannelID(id uint64) *MessageUpdate {
	mu.mutation.SetChannelID(id)
	return mu
}

// SetNillableChannelID sets the "channel" edge to the Channel entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableChannelID(id *uint64) *MessageUpdate {
	if id != nil {
		mu = mu.SetChannelID(*id)
	}
	return mu
}

// SetChannel sets the "channel" edge to the Channel entity.
func (mu *MessageUpdate) SetChannel(c *Channel) *MessageUpdate {
	return mu.SetChannelID(c.ID)
}

// SetParentID sets the "parent" edge to the Message entity by ID.
func (mu *MessageUpdate) SetParentID(id uint64) *MessageUpdate {
	mu.mutation.SetParentID(id)
	return mu
}

// SetNillableParentID sets the "parent" edge to the Message entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableParentID(id *uint64) *MessageUpdate {
	if id != nil {
		mu = mu.SetParentID(*id)
	}
	return mu
}

// SetParent sets the "parent" edge to the Message entity.
func (mu *MessageUpdate) SetParent(m *Message) *MessageUpdate {
	return mu.SetParentID(m.ID)
}

// AddReplyIDs adds the "replies" edge to the Message entity by IDs.
func (mu *MessageUpdate) AddReplyIDs(ids ...uint64) *MessageUpdate {
	mu.mutation.AddReplyIDs(ids...)
	return mu
}

// AddReplies adds the "replies" edges to the Message entity.
func (mu *MessageUpdate) AddReplies(m ...*Message) *MessageUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddReplyIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MessageUpdate) ClearUser() *MessageUpdate {
	mu.mutation.ClearUser()
	return mu
}

// ClearChannel clears the "channel" edge to the Channel entity.
func (mu *MessageUpdate) ClearChannel() *MessageUpdate {
	mu.mutation.ClearChannel()
	return mu
}

// ClearParent clears the "parent" edge to the Message entity.
func (mu *MessageUpdate) ClearParent() *MessageUpdate {
	mu.mutation.ClearParent()
	return mu
}

// ClearReplies clears all "replies" edges to the Message entity.
func (mu *MessageUpdate) ClearReplies() *MessageUpdate {
	mu.mutation.ClearReplies()
	return mu
}

// RemoveReplyIDs removes the "replies" edge to Message entities by IDs.
func (mu *MessageUpdate) RemoveReplyIDs(ids ...uint64) *MessageUpdate {
	mu.mutation.RemoveReplyIDs(ids...)
	return mu
}

// RemoveReplies removes "replies" edges to Message entities.
func (mu *MessageUpdate) RemoveReplies(m ...*Message) *MessageUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveReplyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Createdat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldCreatedat,
		})
	}
	if value, ok := mu.mutation.Editedat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldEditedat,
		})
	}
	if mu.mutation.EditedatCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: message.FieldEditedat,
		})
	}
	if value, ok := mu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldMetadata,
		})
	}
	if mu.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: message.FieldMetadata,
		})
	}
	if value, ok := mu.mutation.Override(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldOverride,
		})
	}
	if mu.mutation.OverrideCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: message.FieldOverride,
		})
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldContent,
		})
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ChannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ChannelTable,
			Columns: []string{message.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: channel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ChannelTable,
			Columns: []string{message.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
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
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
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
	if mu.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
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
	if nodes := mu.mutation.RemovedRepliesIDs(); len(nodes) > 0 && !mu.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetCreatedat sets the "createdat" field.
func (muo *MessageUpdateOne) SetCreatedat(t time.Time) *MessageUpdateOne {
	muo.mutation.SetCreatedat(t)
	return muo
}

// SetNillableCreatedat sets the "createdat" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableCreatedat(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetCreatedat(*t)
	}
	return muo
}

// SetEditedat sets the "editedat" field.
func (muo *MessageUpdateOne) SetEditedat(t time.Time) *MessageUpdateOne {
	muo.mutation.SetEditedat(t)
	return muo
}

// SetNillableEditedat sets the "editedat" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableEditedat(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetEditedat(*t)
	}
	return muo
}

// ClearEditedat clears the value of the "editedat" field.
func (muo *MessageUpdateOne) ClearEditedat() *MessageUpdateOne {
	muo.mutation.ClearEditedat()
	return muo
}

// SetMetadata sets the "metadata" field.
func (muo *MessageUpdateOne) SetMetadata(v *v1.Metadata) *MessageUpdateOne {
	muo.mutation.SetMetadata(v)
	return muo
}

// ClearMetadata clears the value of the "metadata" field.
func (muo *MessageUpdateOne) ClearMetadata() *MessageUpdateOne {
	muo.mutation.ClearMetadata()
	return muo
}

// SetOverride sets the "override" field.
func (muo *MessageUpdateOne) SetOverride(v *v1.Override) *MessageUpdateOne {
	muo.mutation.SetOverride(v)
	return muo
}

// ClearOverride clears the value of the "override" field.
func (muo *MessageUpdateOne) ClearOverride() *MessageUpdateOne {
	muo.mutation.ClearOverride()
	return muo
}

// SetContent sets the "Content" field.
func (muo *MessageUpdateOne) SetContent(v *v1.Content) *MessageUpdateOne {
	muo.mutation.SetContent(v)
	return muo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (muo *MessageUpdateOne) SetUserID(id uint64) *MessageUpdateOne {
	muo.mutation.SetUserID(id)
	return muo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableUserID(id *uint64) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetUserID(*id)
	}
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MessageUpdateOne) SetUser(u *User) *MessageUpdateOne {
	return muo.SetUserID(u.ID)
}

// SetChannelID sets the "channel" edge to the Channel entity by ID.
func (muo *MessageUpdateOne) SetChannelID(id uint64) *MessageUpdateOne {
	muo.mutation.SetChannelID(id)
	return muo
}

// SetNillableChannelID sets the "channel" edge to the Channel entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableChannelID(id *uint64) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetChannelID(*id)
	}
	return muo
}

// SetChannel sets the "channel" edge to the Channel entity.
func (muo *MessageUpdateOne) SetChannel(c *Channel) *MessageUpdateOne {
	return muo.SetChannelID(c.ID)
}

// SetParentID sets the "parent" edge to the Message entity by ID.
func (muo *MessageUpdateOne) SetParentID(id uint64) *MessageUpdateOne {
	muo.mutation.SetParentID(id)
	return muo
}

// SetNillableParentID sets the "parent" edge to the Message entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableParentID(id *uint64) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetParentID(*id)
	}
	return muo
}

// SetParent sets the "parent" edge to the Message entity.
func (muo *MessageUpdateOne) SetParent(m *Message) *MessageUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddReplyIDs adds the "replies" edge to the Message entity by IDs.
func (muo *MessageUpdateOne) AddReplyIDs(ids ...uint64) *MessageUpdateOne {
	muo.mutation.AddReplyIDs(ids...)
	return muo
}

// AddReplies adds the "replies" edges to the Message entity.
func (muo *MessageUpdateOne) AddReplies(m ...*Message) *MessageUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddReplyIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MessageUpdateOne) ClearUser() *MessageUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// ClearChannel clears the "channel" edge to the Channel entity.
func (muo *MessageUpdateOne) ClearChannel() *MessageUpdateOne {
	muo.mutation.ClearChannel()
	return muo
}

// ClearParent clears the "parent" edge to the Message entity.
func (muo *MessageUpdateOne) ClearParent() *MessageUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// ClearReplies clears all "replies" edges to the Message entity.
func (muo *MessageUpdateOne) ClearReplies() *MessageUpdateOne {
	muo.mutation.ClearReplies()
	return muo
}

// RemoveReplyIDs removes the "replies" edge to Message entities by IDs.
func (muo *MessageUpdateOne) RemoveReplyIDs(ids ...uint64) *MessageUpdateOne {
	muo.mutation.RemoveReplyIDs(ids...)
	return muo
}

// RemoveReplies removes "replies" edges to Message entities.
func (muo *MessageUpdateOne) RemoveReplies(m ...*Message) *MessageUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveReplyIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Message.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Createdat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldCreatedat,
		})
	}
	if value, ok := muo.mutation.Editedat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldEditedat,
		})
	}
	if muo.mutation.EditedatCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: message.FieldEditedat,
		})
	}
	if value, ok := muo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldMetadata,
		})
	}
	if muo.mutation.MetadataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: message.FieldMetadata,
		})
	}
	if value, ok := muo.mutation.Override(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldOverride,
		})
	}
	if muo.mutation.OverrideCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: message.FieldOverride,
		})
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldContent,
		})
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ChannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ChannelTable,
			Columns: []string{message.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: channel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ChannelTable,
			Columns: []string{message.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
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
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
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
	if muo.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
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
	if nodes := muo.mutation.RemovedRepliesIDs(); len(nodes) > 0 && !muo.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.RepliesTable,
			Columns: []string{message.RepliesColumn},
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
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
