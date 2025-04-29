package services

import (
    "errors"
    "sync"

    "github.com/gitwooz/go-gin-app/models"
)

type UserService struct {
    users map[string]*models.User
    mu    sync.RWMutex
}

func NewUserService() *UserService {
    return &UserService{
        users: make(map[string]*models.User),
    }
}

func (s *UserService) CreateUser(phoneNumber string) (*models.User, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.users[phoneNumber]; exists {
        return nil, errors.New("user already exists")
    }

    user := &models.User{
        PhoneNumber: phoneNumber,
    }

    s.users[phoneNumber] = user
    return user, nil
}

func (s *UserService) GetUser(phoneNumber string) (*models.User, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    user, exists := s.users[phoneNumber]
    if !exists {
        return nil, errors.New("user not found")
    }

    return user, nil
}

func (s *UserService) UpdateUser(phoneNumber string, updatedUser models.User) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    user, exists := s.users[phoneNumber]
    if !exists {
        return errors.New("user not found")
    }

    user.PhoneNumber = updatedUser.PhoneNumber
    user.Password = updatedUser.Password
    return nil
}