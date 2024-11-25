package service

import (
	"music"
	"music/pkg/repository"
)

type AlbumService struct {
	repo repository.Album
}

func NewAlbumService(repo repository.Album) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) CreateAlbum(album music.Album) (int, error) {
	return s.repo.CreateAlbum(album)
}
