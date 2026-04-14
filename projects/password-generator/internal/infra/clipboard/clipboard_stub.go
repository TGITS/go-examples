package clipboard

func New() Copier {
	return NoopCopier{}
}
