package dao

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"database/sql"
)

type ListDao struct {
	db *database.ConnectDB
}

func NewListDao() *ListDao {
	return &ListDao{
		db: database.Instance(),
	}
}

func (l *ListDao) Create(list *domain.List) (err error) {
	var query bytes.Buffer
	query.WriteString("INSERT INTO lists (name) VALUES (?)")

	source := l.db.Source().(*database.MySQL)
	conn := source.Conn().(*sql.DB)

	stmt, err := conn.Prepare(query.String())
	if err != nil {
		return err
	}

	rs, err := stmt.Exec(list.Name())
	if err != nil {
		return err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return err
	}

	list.WithId(id)

	return nil
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
