package service

import "awesomeGoProject/internal/repository"

type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}
