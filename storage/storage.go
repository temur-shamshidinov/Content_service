package storage

import "github.com/temur-shamshidinov/Content_service/storage/postgres"

type StorageI interface{
	GetContentRepo() postgres.ContentRepoI
}

type storage struct {
	contentRepo postgres.ContentRepoI
}

func NewStorage () StorageI {
	return &storage{
		contentRepo: postgres.NewContentRepo(),
	}
}

func (s * storage) GetContentRepo() postgres.ContentRepoI {

	return s.contentRepo
}