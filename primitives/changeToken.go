package primitives

type ChangeToken struct {
	changed []func()
}

func NewChangeToken() *ChangeToken {
	return &ChangeToken{}
}

func (t *ChangeToken) RegisterChangeCallback(callback func()) {
	t.changed = append(t.changed, callback)
}

func (t *ChangeToken) OnReload() {
	for _, f := range t.changed {
		f()
	}
}
