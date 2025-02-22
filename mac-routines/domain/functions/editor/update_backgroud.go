package editor

func (v *EditorManager) UpdateEditorBackground(url string, settings *map[string]interface{}) {
	bgSettings, ok := (*settings)["background.fullscreen"].(map[string]interface{})
	if !ok {
		bgSettings = map[string]interface{}{
			"opacity":  0.85,
			"position": "center",
			"size":     "cover",
		}
	}

	bgSettings["images"] = []string{url}

	(*settings)["background.fullscreen"] = bgSettings
}
