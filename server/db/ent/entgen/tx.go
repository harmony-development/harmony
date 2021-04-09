// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// ActionButton is the client for interacting with the ActionButton builders.
	ActionButton *ActionButtonClient
	// ActionDropdown is the client for interacting with the ActionDropdown builders.
	ActionDropdown *ActionDropdownClient
	// ActionInput is the client for interacting with the ActionInput builders.
	ActionInput *ActionInputClient
	// Channel is the client for interacting with the Channel builders.
	Channel *ChannelClient
	// EmbedAction is the client for interacting with the EmbedAction builders.
	EmbedAction *EmbedActionClient
	// EmbedField is the client for interacting with the EmbedField builders.
	EmbedField *EmbedFieldClient
	// EmbedMessage is the client for interacting with the EmbedMessage builders.
	EmbedMessage *EmbedMessageClient
	// Emote is the client for interacting with the Emote builders.
	Emote *EmoteClient
	// EmotePack is the client for interacting with the EmotePack builders.
	EmotePack *EmotePackClient
	// File is the client for interacting with the File builders.
	File *FileClient
	// FileHash is the client for interacting with the FileHash builders.
	FileHash *FileHashClient
	// FileMessage is the client for interacting with the FileMessage builders.
	FileMessage *FileMessageClient
	// ForeignUser is the client for interacting with the ForeignUser builders.
	ForeignUser *ForeignUserClient
	// Guild is the client for interacting with the Guild builders.
	Guild *GuildClient
	// GuildListEntry is the client for interacting with the GuildListEntry builders.
	GuildListEntry *GuildListEntryClient
	// Invite is the client for interacting with the Invite builders.
	Invite *InviteClient
	// LocalUser is the client for interacting with the LocalUser builders.
	LocalUser *LocalUserClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// PermissionNode is the client for interacting with the PermissionNode builders.
	PermissionNode *PermissionNodeClient
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// TextMessage is the client for interacting with the TextMessage builders.
	TextMessage *TextMessageClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserMeta is the client for interacting with the UserMeta builders.
	UserMeta *UserMetaClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Committer method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollbacker method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.ActionButton = NewActionButtonClient(tx.config)
	tx.ActionDropdown = NewActionDropdownClient(tx.config)
	tx.ActionInput = NewActionInputClient(tx.config)
	tx.Channel = NewChannelClient(tx.config)
	tx.EmbedAction = NewEmbedActionClient(tx.config)
	tx.EmbedField = NewEmbedFieldClient(tx.config)
	tx.EmbedMessage = NewEmbedMessageClient(tx.config)
	tx.Emote = NewEmoteClient(tx.config)
	tx.EmotePack = NewEmotePackClient(tx.config)
	tx.File = NewFileClient(tx.config)
	tx.FileHash = NewFileHashClient(tx.config)
	tx.FileMessage = NewFileMessageClient(tx.config)
	tx.ForeignUser = NewForeignUserClient(tx.config)
	tx.Guild = NewGuildClient(tx.config)
	tx.GuildListEntry = NewGuildListEntryClient(tx.config)
	tx.Invite = NewInviteClient(tx.config)
	tx.LocalUser = NewLocalUserClient(tx.config)
	tx.Message = NewMessageClient(tx.config)
	tx.PermissionNode = NewPermissionNodeClient(tx.config)
	tx.Profile = NewProfileClient(tx.config)
	tx.Role = NewRoleClient(tx.config)
	tx.Session = NewSessionClient(tx.config)
	tx.TextMessage = NewTextMessageClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.UserMeta = NewUserMetaClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: ActionButton.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
