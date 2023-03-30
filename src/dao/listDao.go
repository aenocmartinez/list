package dao

import "abix360/src/domain"

type ListDao struct{}

func (l *ListDao) Create(list domain.List) (err error) {
	return err
}

func (l *ListDao) Update(list domain.List) (err error) {
	return err
}

func (l *ListDao) Delete(list domain.List) (err error) {
	return err
}

func (l *ListDao) FindById(id int64) (list domain.List) {
	return list
}

func (l *ListDao) AddVideo(list domain.List, video domain.Video) (err error) {
	return err
}

func (l *ListDao) RemoveVideo(list domain.List, video domain.Video) (err error) {
	return err
}

func (l *ListDao) Videos(list domain.List) (videos []domain.Video) {
	return videos
}
func (l *ListDao) All() (lists []domain.List) {
	return lists
}
