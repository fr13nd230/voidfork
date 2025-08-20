package lib

type InitRepo interface {
	Init() error
}
type InitConfig struct {
	Path string
}

type CatFileCmd interface {
	CatFile() error
}
type CatFileConfig struct {
	InitPath string
}