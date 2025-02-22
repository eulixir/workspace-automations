package editor

func (v *EditorManager) UpdateEditorTheme(theme string, settings *map[string]interface{}) {
	(*settings)["workbench.colorTheme"] = theme
}
