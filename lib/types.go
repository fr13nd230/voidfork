package lib

type InitRepo interface {
	Init() (error)
}
type InitConfig struct {
	Path string
}
