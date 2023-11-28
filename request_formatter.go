package cache_wrapper

//go:generate mockery --name=RequestFormatter --outpkg=mocks
type RequestFormatter interface {
	GetUniqKey(string, ...interface{}) []byte
	MarshalWrapI
}

type MarshalWrapI interface {
	MarshalWrapper(...interface{}) ([]byte, error)
	UnMarshalWrapper([]byte, any) ([]interface{}, error)
}
