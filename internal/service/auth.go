package service

import (
	"context"
	"errors"

	"github.com/amirabaris/go-auth/internal/config"
	"github.com/amirabaris/go-auth/internal/crypto"
	"github.com/amirabaris/go-auth/internal/db"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	q   *db.Queries
	cfg *config.Config
}

func New(q *db.Queries, cfg *config.Config) *Service {
	return &Service{q: q, cfg: cfg}
}

func (s *Service) Register(ctx context.Context, email, password string) (string, error) {
	email, err := validateCredentials(email, password)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if _, err := s.q.GetUserByEmail(ctx, email); err == nil {
		return "", ErrEmailAlreadyExists
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return "", err
	}

	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	user, err := s.q.CreateUser(ctx, db.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return "", err
	}

	token, error := GenerateToken(user.ID, s.cfg.JWTSecret)
	if error != nil {
		return "", err
	}

	return token, nil

}

func (s *Service) Login(ctx context.Context, email, password string) {

}
