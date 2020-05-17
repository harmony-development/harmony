// Code generated by sqlc. DO NOT EDIT.

package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addAttachmentStmt, err = db.PrepareContext(ctx, addAttachment); err != nil {
		return nil, fmt.Errorf("error preparing query AddAttachment: %w", err)
	}
	if q.addMessageStmt, err = db.PrepareContext(ctx, addMessage); err != nil {
		return nil, fmt.Errorf("error preparing query AddMessage: %w", err)
	}
	if q.addSessionStmt, err = db.PrepareContext(ctx, addSession); err != nil {
		return nil, fmt.Errorf("error preparing query AddSession: %w", err)
	}
	if q.addUserToGuildStmt, err = db.PrepareContext(ctx, addUserToGuild); err != nil {
		return nil, fmt.Errorf("error preparing query AddUserToGuild: %w", err)
	}
	if q.createChannelStmt, err = db.PrepareContext(ctx, createChannel); err != nil {
		return nil, fmt.Errorf("error preparing query CreateChannel: %w", err)
	}
	if q.createGuildStmt, err = db.PrepareContext(ctx, createGuild); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGuild: %w", err)
	}
	if q.createGuildInviteStmt, err = db.PrepareContext(ctx, createGuildInvite); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGuildInvite: %w", err)
	}
	if q.deleteChannelStmt, err = db.PrepareContext(ctx, deleteChannel); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteChannel: %w", err)
	}
	if q.deleteGuildStmt, err = db.PrepareContext(ctx, deleteGuild); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteGuild: %w", err)
	}
	if q.deleteInviteStmt, err = db.PrepareContext(ctx, deleteInvite); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteInvite: %w", err)
	}
	if q.deleteMessageStmt, err = db.PrepareContext(ctx, deleteMessage); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMessage: %w", err)
	}
	if q.getAttachmentsStmt, err = db.PrepareContext(ctx, getAttachments); err != nil {
		return nil, fmt.Errorf("error preparing query GetAttachments: %w", err)
	}
	if q.getChannelsStmt, err = db.PrepareContext(ctx, getChannels); err != nil {
		return nil, fmt.Errorf("error preparing query GetChannels: %w", err)
	}
	if q.getGuildMembersStmt, err = db.PrepareContext(ctx, getGuildMembers); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildMembers: %w", err)
	}
	if q.getGuildOwnerStmt, err = db.PrepareContext(ctx, getGuildOwner); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildOwner: %w", err)
	}
	if q.getGuildPictureStmt, err = db.PrepareContext(ctx, getGuildPicture); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildPicture: %w", err)
	}
	if q.getMessageStmt, err = db.PrepareContext(ctx, getMessage); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessage: %w", err)
	}
	if q.getMessageAuthorStmt, err = db.PrepareContext(ctx, getMessageAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessageAuthor: %w", err)
	}
	if q.getMessageDateStmt, err = db.PrepareContext(ctx, getMessageDate); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessageDate: %w", err)
	}
	if q.getMessagesStmt, err = db.PrepareContext(ctx, getMessages); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessages: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.guildsForUserStmt, err = db.PrepareContext(ctx, guildsForUser); err != nil {
		return nil, fmt.Errorf("error preparing query GuildsForUser: %w", err)
	}
	if q.incrementInviteStmt, err = db.PrepareContext(ctx, incrementInvite); err != nil {
		return nil, fmt.Errorf("error preparing query IncrementInvite: %w", err)
	}
	if q.openInvitesStmt, err = db.PrepareContext(ctx, openInvites); err != nil {
		return nil, fmt.Errorf("error preparing query OpenInvites: %w", err)
	}
	if q.removeUserFromGuildStmt, err = db.PrepareContext(ctx, removeUserFromGuild); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveUserFromGuild: %w", err)
	}
	if q.resolveGuildIDStmt, err = db.PrepareContext(ctx, resolveGuildID); err != nil {
		return nil, fmt.Errorf("error preparing query ResolveGuildID: %w", err)
	}
	if q.sessionToUserIDStmt, err = db.PrepareContext(ctx, sessionToUserID); err != nil {
		return nil, fmt.Errorf("error preparing query SessionToUserID: %w", err)
	}
	if q.setGuildNameStmt, err = db.PrepareContext(ctx, setGuildName); err != nil {
		return nil, fmt.Errorf("error preparing query SetGuildName: %w", err)
	}
	if q.setGuildPictureStmt, err = db.PrepareContext(ctx, setGuildPicture); err != nil {
		return nil, fmt.Errorf("error preparing query SetGuildPicture: %w", err)
	}
	if q.userInGuildStmt, err = db.PrepareContext(ctx, userInGuild); err != nil {
		return nil, fmt.Errorf("error preparing query UserInGuild: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAttachmentStmt != nil {
		if cerr := q.addAttachmentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAttachmentStmt: %w", cerr)
		}
	}
	if q.addMessageStmt != nil {
		if cerr := q.addMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addMessageStmt: %w", cerr)
		}
	}
	if q.addSessionStmt != nil {
		if cerr := q.addSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addSessionStmt: %w", cerr)
		}
	}
	if q.addUserToGuildStmt != nil {
		if cerr := q.addUserToGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserToGuildStmt: %w", cerr)
		}
	}
	if q.createChannelStmt != nil {
		if cerr := q.createChannelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createChannelStmt: %w", cerr)
		}
	}
	if q.createGuildStmt != nil {
		if cerr := q.createGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGuildStmt: %w", cerr)
		}
	}
	if q.createGuildInviteStmt != nil {
		if cerr := q.createGuildInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGuildInviteStmt: %w", cerr)
		}
	}
	if q.deleteChannelStmt != nil {
		if cerr := q.deleteChannelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteChannelStmt: %w", cerr)
		}
	}
	if q.deleteGuildStmt != nil {
		if cerr := q.deleteGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteGuildStmt: %w", cerr)
		}
	}
	if q.deleteInviteStmt != nil {
		if cerr := q.deleteInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteInviteStmt: %w", cerr)
		}
	}
	if q.deleteMessageStmt != nil {
		if cerr := q.deleteMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMessageStmt: %w", cerr)
		}
	}
	if q.getAttachmentsStmt != nil {
		if cerr := q.getAttachmentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAttachmentsStmt: %w", cerr)
		}
	}
	if q.getChannelsStmt != nil {
		if cerr := q.getChannelsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getChannelsStmt: %w", cerr)
		}
	}
	if q.getGuildMembersStmt != nil {
		if cerr := q.getGuildMembersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildMembersStmt: %w", cerr)
		}
	}
	if q.getGuildOwnerStmt != nil {
		if cerr := q.getGuildOwnerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildOwnerStmt: %w", cerr)
		}
	}
	if q.getGuildPictureStmt != nil {
		if cerr := q.getGuildPictureStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildPictureStmt: %w", cerr)
		}
	}
	if q.getMessageStmt != nil {
		if cerr := q.getMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageStmt: %w", cerr)
		}
	}
	if q.getMessageAuthorStmt != nil {
		if cerr := q.getMessageAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageAuthorStmt: %w", cerr)
		}
	}
	if q.getMessageDateStmt != nil {
		if cerr := q.getMessageDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageDateStmt: %w", cerr)
		}
	}
	if q.getMessagesStmt != nil {
		if cerr := q.getMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessagesStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.guildsForUserStmt != nil {
		if cerr := q.guildsForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing guildsForUserStmt: %w", cerr)
		}
	}
	if q.incrementInviteStmt != nil {
		if cerr := q.incrementInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing incrementInviteStmt: %w", cerr)
		}
	}
	if q.openInvitesStmt != nil {
		if cerr := q.openInvitesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing openInvitesStmt: %w", cerr)
		}
	}
	if q.removeUserFromGuildStmt != nil {
		if cerr := q.removeUserFromGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeUserFromGuildStmt: %w", cerr)
		}
	}
	if q.resolveGuildIDStmt != nil {
		if cerr := q.resolveGuildIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing resolveGuildIDStmt: %w", cerr)
		}
	}
	if q.sessionToUserIDStmt != nil {
		if cerr := q.sessionToUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing sessionToUserIDStmt: %w", cerr)
		}
	}
	if q.setGuildNameStmt != nil {
		if cerr := q.setGuildNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setGuildNameStmt: %w", cerr)
		}
	}
	if q.setGuildPictureStmt != nil {
		if cerr := q.setGuildPictureStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setGuildPictureStmt: %w", cerr)
		}
	}
	if q.userInGuildStmt != nil {
		if cerr := q.userInGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userInGuildStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                      DBTX
	tx                      *sql.Tx
	addAttachmentStmt       *sql.Stmt
	addMessageStmt          *sql.Stmt
	addSessionStmt          *sql.Stmt
	addUserToGuildStmt      *sql.Stmt
	createChannelStmt       *sql.Stmt
	createGuildStmt         *sql.Stmt
	createGuildInviteStmt   *sql.Stmt
	deleteChannelStmt       *sql.Stmt
	deleteGuildStmt         *sql.Stmt
	deleteInviteStmt        *sql.Stmt
	deleteMessageStmt       *sql.Stmt
	getAttachmentsStmt      *sql.Stmt
	getChannelsStmt         *sql.Stmt
	getGuildMembersStmt     *sql.Stmt
	getGuildOwnerStmt       *sql.Stmt
	getGuildPictureStmt     *sql.Stmt
	getMessageStmt          *sql.Stmt
	getMessageAuthorStmt    *sql.Stmt
	getMessageDateStmt      *sql.Stmt
	getMessagesStmt         *sql.Stmt
	getUserStmt             *sql.Stmt
	guildsForUserStmt       *sql.Stmt
	incrementInviteStmt     *sql.Stmt
	openInvitesStmt         *sql.Stmt
	removeUserFromGuildStmt *sql.Stmt
	resolveGuildIDStmt      *sql.Stmt
	sessionToUserIDStmt     *sql.Stmt
	setGuildNameStmt        *sql.Stmt
	setGuildPictureStmt     *sql.Stmt
	userInGuildStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                      tx,
		tx:                      tx,
		addAttachmentStmt:       q.addAttachmentStmt,
		addMessageStmt:          q.addMessageStmt,
		addSessionStmt:          q.addSessionStmt,
		addUserToGuildStmt:      q.addUserToGuildStmt,
		createChannelStmt:       q.createChannelStmt,
		createGuildStmt:         q.createGuildStmt,
		createGuildInviteStmt:   q.createGuildInviteStmt,
		deleteChannelStmt:       q.deleteChannelStmt,
		deleteGuildStmt:         q.deleteGuildStmt,
		deleteInviteStmt:        q.deleteInviteStmt,
		deleteMessageStmt:       q.deleteMessageStmt,
		getAttachmentsStmt:      q.getAttachmentsStmt,
		getChannelsStmt:         q.getChannelsStmt,
		getGuildMembersStmt:     q.getGuildMembersStmt,
		getGuildOwnerStmt:       q.getGuildOwnerStmt,
		getGuildPictureStmt:     q.getGuildPictureStmt,
		getMessageStmt:          q.getMessageStmt,
		getMessageAuthorStmt:    q.getMessageAuthorStmt,
		getMessageDateStmt:      q.getMessageDateStmt,
		getMessagesStmt:         q.getMessagesStmt,
		getUserStmt:             q.getUserStmt,
		guildsForUserStmt:       q.guildsForUserStmt,
		incrementInviteStmt:     q.incrementInviteStmt,
		openInvitesStmt:         q.openInvitesStmt,
		removeUserFromGuildStmt: q.removeUserFromGuildStmt,
		resolveGuildIDStmt:      q.resolveGuildIDStmt,
		sessionToUserIDStmt:     q.sessionToUserIDStmt,
		setGuildNameStmt:        q.setGuildNameStmt,
		setGuildPictureStmt:     q.setGuildPictureStmt,
		userInGuildStmt:         q.userInGuildStmt,
	}
}
