package clipboard

type Copier interface {
	Copy(text string) error
}

type NoopCopier struct{}

func (NoopCopier) Copy(text string) error {
	return nil
}
