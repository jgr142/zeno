package theme

import "github.com/gdamore/tcell/v2"

// Theme defines the application's color scheme.
type Theme struct {
	Primary    tcell.Color
	Secondary  tcell.Color
	Tertiary   tcell.Color
	Background tcell.Color
}

// New returns a new dark pastel-style theme.
func New() *Theme {
	return &Theme{
		Primary:    tcell.ColorSlateBlue,     // Muted purple-blue, good for titles or borders
		Secondary:  tcell.ColorDarkGray,      // Low-brightness gray for content
		Tertiary:   tcell.ColorCadetBlue,     // Muted teal, adds subtle color
		Background: tcell.ColorDarkSlateGray, // Deep gray-blue background, easy on the eyes
	}
}
