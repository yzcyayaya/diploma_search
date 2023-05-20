package model

//文档顶层,每个model都应该具有
type Document interface {
	GetIndex() string
	GetTitle() string
	GetGenres() []string
	GetDocId() string
}
