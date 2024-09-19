package builder

type Builder interface{}

type builder struct{}

func NewBuilder() Builder {
	return &builder{}
}
