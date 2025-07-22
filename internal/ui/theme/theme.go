package theme

import "github.com/gdamore/tcell/v2"

// Theme defines the application's color scheme.
type Theme struct {
	Primary     tcell.Color // Used for borders/titles
	Secondary   tcell.Color // For background elements or less prominent UI
	Tertiary    tcell.Color // For subtle highlights
	Background  tcell.Color // Main background color
	PrimaryText tcell.Color // Main text color
	Accent      tcell.Color // Used for selection/highlight/active item
}

// New returns a new dark pastel-style theme.
func New() *Theme {
	return &Theme{
		Primary:     tcell.ColorMediumOrchid, // Muted purple, replaces SlateBlue
		Secondary:   tcell.ColorDarkGray,     // Keeps the subtle gray
		Tertiary:    tcell.ColorThistle,      // Soft lavender-gray, replaces CadetBlue
		Background:  tcell.ColorBlack,        // Dark base remains for contrast
		PrimaryText: tcell.ColorPlum,         // Soft readable pastel purple
		Accent:      tcell.ColorMediumPurple, // Highlight remains (fits well)
	}
}
