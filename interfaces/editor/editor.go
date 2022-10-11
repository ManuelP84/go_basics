package editor

type Editor struct {
	Name string
}

func (editor Editor) Permissions() int {
	return 3
}

func (editor Editor) GetName() string {
	return editor.Name
}
