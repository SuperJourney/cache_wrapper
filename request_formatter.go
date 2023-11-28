package cache_wrapper

//go:generate mockery --name=RequestFormatter --outpkg=mocks
type RequestFormatter interface {
	GetUniqKey(...interface{}) ([]byte, error)
	MarshalWrapI
}

type MarshalWrapI interface {
	MarshalWrapper(...interface{}) ([]byte, error)
	UnMarshalWrapper([]byte, interface{}) ([]interface{}, error)
}
