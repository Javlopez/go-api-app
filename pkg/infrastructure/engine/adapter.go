package engine

type Adapter interface {
	Connect() Adapter
}
