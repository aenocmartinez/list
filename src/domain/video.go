package domain

type Video struct {
	id        int64
	name      string
	url       string
	createdAt string
	updatedAt string
	list      List
}

func NewVideo(name, url string) *Video {
	return &Video{
		name: name,
		url:  url,
	}
}

func (v *Video) WithId(id int64) *Video {
	v.id = id
	return v
}

func (v *Video) WithName(name string) *Video {
	v.name = name
	return v
}

func (v *Video) WithUrl(url string) *Video {
	v.url = url
	return v
}

func (v *Video) WithCreatedAt(createdAt string) *Video {
	v.createdAt = createdAt
	return v
}

func (v *Video) WithUpdatedAt(updatedAt string) *Video {
	v.updatedAt = updatedAt
	return v
}

func (v *Video) WithList(list List) *Video {
	v.list = list
	return v
}

func (v *Video) Id() int64 {
	return v.id
}

func (v *Video) Name() string {
	return v.name
}

func (v *Video) Url() string {
	return v.url
}

func (v *Video) CreatedAt() string {
	return v.createdAt
}

func (v *Video) UpdatedAt() string {
	return v.updatedAt
}

func (v *Video) ListId() int64 {
	return v.list.Id()
}

func (v *Video) ListName() string {
	return v.list.Name()
}

func (v *Video) ListVideos() []Video {
	return v.list.Videos()
}
