package repository_psql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/answersuck/vault/internal/adapter/repository/psql"
	"github.com/answersuck/vault/internal/domain/account"
	"github.com/answersuck/vault/internal/domain/session"
	"github.com/answersuck/vault/pkg/strings"
)

var sessionRepo *psql.SessionRepo

type sessionRepoTestSuite struct {
	suite.Suite
}

func insertTestSession(s *session.Session) (*session.Session, error) {
	id, err := strings.NewUnique(session.SessionIdLen)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	s.Id = id
	s.CreatedAt = now
	s.ExpiresAt = now.Local().Add(time.Minute * time.Duration(2)).Unix()
	s.MaxAge = 1
	s.UserAgent = "ua"
	s.IP = "192.0.0.1"

	_, err = postgresClient.Pool.Exec(
		context.Background(),
		`INSERT INTO session (id, account_id, max_age, user_agent, ip, expires_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		s.Id, s.AccountId, s.MaxAge, s.UserAgent, s.IP, s.ExpiresAt, s.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func TestSessionRepoTestSuite(t *testing.T) { suite.Run(t, new(sessionRepoTestSuite)) }

func (s *sessionRepoTestSuite) TestSave() {
	a, err := insertTestAccount(account.Account{})
	s.NoError(err)

	sessionId, err := strings.NewUnique(session.SessionIdLen)
	s.NoError(err)

	type args struct {
		ctx context.Context
		s   *session.Session
	}
	now := time.Now()
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "session successfully saved",
			args: args{
				ctx: context.Background(),
				s: &session.Session{
					Id:        sessionId,
					AccountId: a.Id,
					UserAgent: "ua",
					IP:        "192.0.0.1",
					MaxAge:    1,
					ExpiresAt: now.Local().Add(time.Minute * time.Duration(2)).Unix(), // 2 min
					CreatedAt: now,
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "session already exist",
			args: args{
				ctx: context.Background(),
				s: &session.Session{
					Id:        sessionId,
					AccountId: a.Id,
				},
			},
			wantErr: true,
			err:     session.ErrAlreadyExist,
		},
		{
			name: "session account not found",
			args: args{
				ctx: context.Background(),
				s: &session.Session{
					AccountId: "fd550c9a-b3a5-4c03-9f66-9e7ffc0c9523",
				},
			},
			wantErr: true,
			err:     session.ErrAccountNotFound,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := sessionRepo.Save(tt.args.ctx, tt.args.s)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.NoError(t, err)

			var s session.Session
			err = postgresClient.Pool.QueryRow(
				context.Background(),
				"SELECT id, account_id, user_agent, ip, max_age, expires_at, created_at FROM session WHERE id = $1",
				tt.args.s.Id,
			).Scan(&s.Id, &s.AccountId, &s.UserAgent, &s.IP, &s.MaxAge, &s.ExpiresAt, &s.CreatedAt)
			assert.NoError(t, err)
			// hardcoded since assert.Equal is not working with time.Time
			assert.Equal(t, tt.args.s.Id, s.Id)
			assert.Equal(t, tt.args.s.AccountId, s.AccountId)
			assert.Equal(t, tt.args.s.IP, s.IP)
			assert.Equal(t, tt.args.s.MaxAge, s.MaxAge)
			assert.Equal(t, tt.args.s.ExpiresAt, s.ExpiresAt)
			assert.Equal(t, tt.args.s.CreatedAt.Unix(), s.CreatedAt.Unix())
		})
	}
}

func (s *sessionRepoTestSuite) TestFindById() {
	a, err := insertTestAccount(account.Account{})
	s.NoError(err)

	sess, err := insertTestSession(&session.Session{AccountId: a.Id})
	s.NoError(err)

	type args struct {
		ctx       context.Context
		sessionId string
	}
	tests := []struct {
		name    string
		args    args
		want    *session.Session
		wantErr bool
		err     error
	}{
		{
			name: "session found by id",
			args: args{
				ctx:       context.Background(),
				sessionId: sess.Id,
			},
			want:    sess,
			wantErr: false,
			err:     nil,
		},
		{
			name: "session not found",
			args: args{
				ctx:       context.Background(),
				sessionId: "",
			},
			want:    sess,
			wantErr: true,
			err:     session.ErrNotFound,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := sessionRepo.FindById(tt.args.ctx, tt.args.sessionId)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.NoError(t, err)

			assert.Equal(t, tt.want.Id, got.Id)
			assert.Equal(t, tt.want.AccountId, got.AccountId)
			assert.Equal(t, tt.want.MaxAge, got.MaxAge)
			assert.Equal(t, tt.want.UserAgent, got.UserAgent)
			assert.Equal(t, tt.want.IP, got.IP)
			assert.Equal(t, tt.want.ExpiresAt, got.ExpiresAt)
			assert.Equal(t, tt.want.CreatedAt.Unix(), got.CreatedAt.Unix())
		})
	}
}

func (s *sessionRepoTestSuite) TestFindWithAccountDetails() {
	a, err := insertTestAccount(account.Account{Verified: true})
	s.NoError(err)

	sess, err := insertTestSession(&session.Session{AccountId: a.Id})
	s.NoError(err)

	type args struct {
		ctx       context.Context
		sessionId string
	}
	tests := []struct {
		name    string
		args    args
		want    *session.WithAccountDetails
		wantErr bool
		err     error
	}{
		{
			name: "session with details found",
			args: args{
				ctx:       context.Background(),
				sessionId: sess.Id,
			},
			want: &session.WithAccountDetails{
				Session:         *sess,
				AccountVerified: a.Verified,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "session with details not found",
			args: args{
				ctx:       context.Background(),
				sessionId: "",
			},
			want: &session.WithAccountDetails{
				Session:         *sess,
				AccountVerified: a.Verified,
			},
			wantErr: true,
			err:     session.ErrNotFound,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := sessionRepo.FindWithAccountDetails(tt.args.ctx, tt.args.sessionId)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.NoError(t, err)

			assert.Equal(t, tt.want.Session.Id, got.Session.Id)
			assert.Equal(t, tt.want.Session.AccountId, got.Session.AccountId)
			assert.Equal(t, tt.want.Session.MaxAge, got.Session.MaxAge)
			assert.Equal(t, tt.want.Session.UserAgent, got.Session.UserAgent)
			assert.Equal(t, tt.want.Session.IP, got.Session.IP)
			assert.Equal(t, tt.want.Session.ExpiresAt, got.Session.ExpiresAt)
			assert.Equal(t, tt.want.Session.CreatedAt.Unix(), got.Session.CreatedAt.Unix())
			assert.Equal(t, tt.want.AccountVerified, got.AccountVerified)
		})
	}
}

func (s *sessionRepoTestSuite) TestFindAll() {
	a, err := insertTestAccount(account.Account{})
	s.NoError(err)

	a1, err := insertTestAccount(account.Account{})
	s.NoError(err)

	var sessions []*session.Session
	for i := 0; i < 10; i++ {
		sess, err := insertTestSession(&session.Session{AccountId: a.Id})
		s.NoError(err)
		sessions = append(sessions, sess)
	}

	type args struct {
		ctx       context.Context
		accountId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []*session.Session
		err     error
	}{
		{
			name: "account has sessions",
			args: args{
				ctx:       context.Background(),
				accountId: a.Id,
			},
			want:    sessions,
			wantErr: false,
			err:     nil,
		},
		{
			name: "account has no sessions",
			args: args{
				ctx:       context.Background(),
				accountId: a1.Id,
			},
			want:    make([]*session.Session, 0),
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := sessionRepo.FindAll(tt.args.ctx, tt.args.accountId)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.NoError(t, err)

			assert.Equal(t, len(tt.want), len(got))
			if len(got) > 0 {
				for _, sess := range got {
					assert.Equal(t, a.Id, sess.AccountId)
				}
			}
		})
	}
}

func (s *sessionRepoTestSuite) TestDelete() {
	a, err := insertTestAccount(account.Account{})
	s.NoError(err)

	sess, err := insertTestSession(&session.Session{AccountId: a.Id})
	s.NoError(err)

	type args struct {
		ctx       context.Context
		sessionId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "session successfully deleted",
			args: args{
				ctx:       context.Background(),
				sessionId: sess.Id,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "session successfully deleted",
			args: args{
				ctx:       context.Background(),
				sessionId: "",
			},
			wantErr: true,
			err:     session.ErrNotFound,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := sessionRepo.Delete(tt.args.ctx, tt.args.sessionId)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.NoError(t, err)
		})
	}
}

//
// func Test_sessionRepo_DeleteWithExcept(t *testing.T) {
// 	type fields struct {
// 		l *zap.Logger
// 		c *postgres.Client
// 	}
// 	type args struct {
// 		ctx       context.Context
// 		accountId string
// 		sessionId string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &sessionRepo{
// 				l: tt.fields.l,
// 				c: tt.fields.c,
// 			}
// 			if err := r.DeleteWithExcept(tt.args.ctx, tt.args.accountId, tt.args.sessionId); (err != nil) != tt.wantErr {
// 				t.Errorf("sessionRepo.DeleteWithExcept() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
//
// func Test_sessionRepo_DeleteAll(t *testing.T) {
// 	type fields struct {
// 		l *zap.Logger
// 		c *postgres.Client
// 	}
// 	type args struct {
// 		ctx       context.Context
// 		accountId string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &sessionRepo{
// 				l: tt.fields.l,
// 				c: tt.fields.c,
// 			}
// 			if err := r.DeleteAll(tt.args.ctx, tt.args.accountId); (err != nil) != tt.wantErr {
// 				t.Errorf("sessionRepo.DeleteAll() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
