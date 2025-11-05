package usecase

import (
	"context"

	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/jwt"
	"easy-go-monitor/internal/user/domain"
)

type (
	// LoginUserUseCase input port
	LoginUserUseCase interface {
		Execute(ctx context.Context, input LoginUserInput) (LoginUserOutput, error)
	}

	// LoginUserInput input data
	LoginUserInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// LoginUserOutput output data
	LoginUserOutput struct {
		Token string `json:"token"`
	}

	loginUserInteractor struct {
		repo domain.UserRepository
		jwt  jwt.JWTService
	}
)

func NewLoginUserInteractor(repo domain.UserRepository, jwt jwt.JWTService) LoginUserUseCase {
	return &loginUserInteractor{
		repo: repo,
		jwt:  jwt,
	}
}

func (i *loginUserInteractor) Execute(ctx context.Context, input LoginUserInput) (LoginUserOutput, error) {

	if input.Email == "" || input.Password == "" {
		return LoginUserOutput{}, codes.ErrInvalidCredentials
	}

	user, err := i.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return LoginUserOutput{}, codes.ErrInvalidCredentials
	}

	if err := user.Authenticate(input.Password); err != nil {
		return LoginUserOutput{}, codes.ErrInvalidCredentials
	}

	token, err := i.jwt.GenerateToken(user.ID)
	if err != nil {
		return LoginUserOutput{}, err
	}

	return LoginUserOutput{Token: token}, nil
}
