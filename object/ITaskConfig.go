package object

type ITaskConfig interface {
	GetConfigId() string
	GetSpec() string
	IsEqual(c ITaskConfig) bool
}
