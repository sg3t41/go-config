package filetype

type ConfigFileType int

const (
	YAML ConfigFileType = iota
	JSON
	INI
)
