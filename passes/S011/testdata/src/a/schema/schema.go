package schema

type Schema struct {
	Computed     bool
	DiffSuppressFunc func()
}
