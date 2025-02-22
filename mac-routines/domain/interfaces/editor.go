package interfaces

type Editor interface {
	UpdateEditorSettings(theme, wallpaper string) error
}
