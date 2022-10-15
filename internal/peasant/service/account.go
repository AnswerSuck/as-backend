package service

import (
	"context"
	"time"

	"github.com/ysomad/answersuck/apperror"
	"github.com/ysomad/answersuck/internal/peasant/domain"
	"github.com/ysomad/answersuck/internal/peasant/service/dto"

	"github.com/ysomad/answersuck/cryptostr"
)

type passwordEncodeComparer interface {
	Encode(plain string) (string, error)
	Compare(plain, encoded string) (bool, error)
}

type accountService struct {
	repo     accountRepository
	password passwordEncodeComparer

	emailVerifCodeLifetime time.Duration
}

func NewAccountService(r accountRepository, p passwordEncodeComparer, emailVerifCodeLifetime time.Duration) (*accountService, error) {
	if r == nil || p == nil || emailVerifCodeLifetime == 0 {
		return nil, apperror.ErrNilArgs
	}

	return &accountService{
		repo:                   r,
		password:               p,
		emailVerifCodeLifetime: emailVerifCodeLifetime,
	}, nil
}

func (s *accountService) Create(ctx context.Context, args dto.AccountCreateArgs) (*domain.Account, error) {
	// TODO: Check if password is not banned

	// TODO: Check if username is not banned

	// TODO: Check if email is real or not banned

	encodedPass, err := s.password.Encode(args.PlainPassword)
	if err != nil {
		return nil, err
	}

	emailVerifCode, err := cryptostr.RandomBase64(32)
	if err != nil {
		return nil, err
	}

	a, err := s.repo.Create(
		ctx,
		dto.AccountSaveArgs{
			Email:           args.Email,
			Username:        args.Username,
			EncodedPassword: encodedPass,
		},
		dto.EmailVerifSaveArgs{
			Code:      emailVerifCode,
			ExpiresAt: time.Now().Add(s.emailVerifCodeLifetime),
		})
	if err != nil {
		return nil, err
	}

	// TODO: Send email with verification code

	return a, nil
}

func (s *accountService) GetByID(ctx context.Context, accountID string) (*domain.Account, error) {
	return s.repo.GetByID(ctx, accountID)
}

func (s *accountService) DeleteByID(ctx context.Context, accountID string) error {
	// TODO: log out all sessions
	return s.repo.DeleteByID(ctx, accountID)
}
