package domain

type List struct {
	id         int64
	name       string
	createdAt  string
	updatedAt  string
	repository ListRepository
}

func NewList(name string) *List {
	return &List{
		name: name,
	}
}

func (l *List) WithRepository(repository ListRepository) *List {
	l.repository = repository
	return l
}

func (l *List) WithId(id int64) *List {
	l.id = id
	return l
}

func (l *List) WithName(name string) *List {
	l.name = name
	return l
}

func (l *List) WithCreatedAt(createdAt string) *List {
	l.createdAt = createdAt
	return l
}

func (l *List) WithUpdatedAt(updatedAt string) *List {
	l.updatedAt = updatedAt
	return l
}

func (l *List) Id() int64 {
	return l.id
}

func (l *List) Name() string {
	return l.name
}

func (l *List) CreatedAt() string {
	return l.createdAt
}

func (l *List) UpdatedAt() string {
	return l.updatedAt
}

func (l *List) AddVideo(video Video) error {
	return l.repository.AddVideo(*l, video)
}

func (l *List) RemoveVideo(video Video) error {
	return l.repository.RemoveVideo(*l, video)
}

func (l *List) Videos() []Video {
	return l.repository.Videos(*l)
}

func (l *List) Exists() bool {
	return l.id > 0
}

func FindListById(id int64, repository ListRepository) List {
	return repository.FindById(id)
}
