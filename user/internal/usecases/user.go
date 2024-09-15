package usecase

import (
	"context"
	"errors"
	"regexp"
	"user/internal/entities"
	"user/internal/repositories"

	"github.com/google/uuid"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *entities.User) error {
	existing_user, _ := u.repo.GetUserByID(ctx, user.ID)

	if existing_user != nil {
		return errors.New("user already exists")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if !isValidUsername(user.Username) {
		return errors.New("invalid username format")
	}

	return u.repo.CreateUser(ctx, user)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,}$`)
	return re.MatchString(username)
}

func (u *userUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *entities.User) error {
	existing_user, _ := u.repo.GetUserByID(ctx, user.ID)

	if existing_user == nil {
		return errors.New("user does not exist")
	}

	return u.repo.UpdateUser(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	existing_user, _ := u.repo.GetUserByID(ctx, id)

	if existing_user == nil {
		return errors.New("user does not exist")
	}

	return u.repo.DeleteUser(ctx, id)
}
