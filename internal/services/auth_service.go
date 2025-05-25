package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/vCif3r/ecommerce-api/internal/models"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
	"time"
)

type AuthService struct {
	repo         *repositories.AuthRepository
	jwtSecret    string
	tokenExpires time.Duration
}

func NewAuthService(repo *repositories.AuthRepository, jwtSecret string, tokenExpires time.Duration) *AuthService {
	return &AuthService{
		repo:         repo,
		jwtSecret:    jwtSecret,
		tokenExpires: tokenExpires,
	}
}

func (s *AuthService) Register(userReq *models.RegisterRequest) (*models.User, error) {
	// Verificar si el usuario ya existe
	existingUser, err := s.repo.FindUserByEmail(userReq.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Crear el usuario
	user := &models.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Email:     userReq.Email,
		Password:  string(hashedPassword),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(loginReq *models.LoginRequest) (*models.AuthResponse, error) {
	// Buscar usuario por email
	user, err := s.repo.FindUserByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Verificar contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generar token JWT
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, err
	}


	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"email": user.Email,
		"exp":  time.Now().Add(s.tokenExpires).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}