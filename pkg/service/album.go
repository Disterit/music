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

func (s *AlbumService) GetAlbums(artistId int) ([]music.Album, error) {
	return s.repo.GetAlbums(artistId)
}

func (s *AlbumService) GetAlbum(artistId, albumId int) (music.Album, error) {
	return s.repo.GetAlbum(artistId, albumId)
}

func (s *AlbumService) UpdateAlbum(album music.Album) error {
	return s.repo.UpdateAlbum(album)
}

func (s *AlbumService) DeleteAlbum(artistId, albumId int) error {
	return s.repo.DeleteAlbum(artistId, albumId)
}
