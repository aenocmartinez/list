package domain

type ListRepository interface {
	Create(list List) (err error)
	Update(list List) (err error)
	Delete(list List) (err error)
	FindById(id int64) (list List)
	AddVideo(list List, video Video) (err error)
	RemoveVideo(list List, video Video) (err error)
	Videos(list List) (videos []Video)
	All() (lists []List)
}
