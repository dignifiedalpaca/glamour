package styles

import "github.com/charmbracelet/glamour/ansi"

type palette struct {
	base          *string
	surface       *string
	overlay       *string
	muted         *string
	subtle        *string
	text          *string
	love          *string
	gold          *string
	rose          *string
	pine          *string
	foam          *string
	iris          *string
	highlightLow  *string
	highlightMed  *string
	highlightHigh *string
	chromaTheme   string
}

var (
	// RosePineStyleConfig is the main Rosé Pine style.
	RosePineStyleConfig = generateStyle(rosePine)
	// RosePineMoonStyleConfig is the Rosé Pine Moon style.
	RosePineMoonStyleConfig = generateStyle(rosePineMoon)
	// RosePineDawnStyleConfig is the Rosé Pine Dawn style
	RosePineDawnStyleConfig = generateStyle(rosePineDawn)
	rosePine                = palette{
		base:          stringPtr("#191724"),
		surface:       stringPtr("#1f1d2e"),
		overlay:       stringPtr("#26233a"),
		muted:         stringPtr("#6e6a86"),
		subtle:        stringPtr("#908caa"),
		text:          stringPtr("#e0def4"),
		love:          stringPtr("#eb6f92"),
		gold:          stringPtr("#f6c177"),
		rose:          stringPtr("#ebbcba"),
		pine:          stringPtr("#31748f"),
		foam:          stringPtr("#9ccfd8"),
		iris:          stringPtr("#c4a7e7"),
		highlightLow:  stringPtr("#21202e"),
		highlightMed:  stringPtr("#403d52"),
		highlightHigh: stringPtr("#403d52"),
		chromaTheme:   "rose-pine",
	}
	rosePineMoon = palette{
		base:          stringPtr("#232136"),
		surface:       stringPtr("#2a273f"),
		overlay:       stringPtr("#393552"),
		muted:         stringPtr("#6e6a86"),
		subtle:        stringPtr("#908caa"),
		text:          stringPtr("#e0def4"),
		love:          stringPtr("#eb6f92"),
		gold:          stringPtr("#f6c177"),
		rose:          stringPtr("#ea9a97"),
		pine:          stringPtr("#3e8fb0"),
		foam:          stringPtr("#9ccfd8"),
		iris:          stringPtr("#c4a7e7"),
		highlightLow:  stringPtr("#2a283e"),
		highlightMed:  stringPtr("#44415a"),
		highlightHigh: stringPtr("#56526e"),
		chromaTheme:   "rose-pine-moon",
	}
	rosePineDawn = palette{
		base:          stringPtr("#191724"),
		surface:       stringPtr("#1f1d2e"),
		overlay:       stringPtr("#26233a"),
		muted:         stringPtr("#6e6a86"),
		subtle:        stringPtr("#908caa"),
		text:          stringPtr("#e0def4"),
		love:          stringPtr("#eb6f92"),
		gold:          stringPtr("#f6c177"),
		rose:          stringPtr("#ebbcba"),
		pine:          stringPtr("#31748f"),
		foam:          stringPtr("#9ccfd8"),
		iris:          stringPtr("#c4a7e7"),
		highlightLow:  stringPtr("#21202e"),
		highlightMed:  stringPtr("#403d52"),
		highlightHigh: stringPtr("#403d52"),
		chromaTheme:   "rose-pine-dawn",
	}
)

func generateStyle(palette palette) ansi.StyleConfig {
	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "\n",
				BlockSuffix: "\n",
				Color:       palette.text,
			},
			Margin: uintPtr(defaultMargin),
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  palette.subtle,
				Italic: boolPtr(true),
			},
			Indent: uintPtr(defaultMargin),
		},
		List: ansi.StyleList{
			LevelIndent: defaultMargin,
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color: palette.subtle,
				},
			},
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockSuffix: "\n",
				Color:       palette.love,
				Bold:        boolPtr(true),
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "# ",
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "## ",
			},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "### ",
			},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "#### ",
			},
		},
		H5: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "##### ",
			},
		},
		H6: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "###### ",
			},
		},
		Strikethrough: ansi.StylePrimitive{
			CrossedOut: boolPtr(true),
			Color:      palette.muted,
		},
		Emph: ansi.StylePrimitive{
			Color:  palette.iris,
			Italic: boolPtr(true),
		},
		Strong: ansi.StylePrimitive{
			Bold:  boolPtr(true),
			Color: palette.rose,
		},
		HorizontalRule: ansi.StylePrimitive{
			Color:  palette.highlightHigh,
			Format: "\n--------\n",
		},
		Item: ansi.StylePrimitive{
			BlockPrefix: "• ",
		},
		Enumeration: ansi.StylePrimitive{
			BlockPrefix: ". ",
			Color:       palette.gold,
		},
		Task: ansi.StyleTask{
			StylePrimitive: ansi.StylePrimitive{},
			Ticked:         "[✓] ",
			Unticked:       "[ ] ",
		},
		Link: ansi.StylePrimitive{
			Color:     palette.foam,
			Underline: boolPtr(true),
		},
		LinkText: ansi.StylePrimitive{
			Color: palette.foam,
		},
		Image: ansi.StylePrimitive{
			Color:     palette.pine,
			Underline: boolPtr(true),
		},
		ImageText: ansi.StylePrimitive{
			Color:  palette.pine,
			Format: "Image: {{.text}} →",
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: palette.gold,
			},
		},
		CodeBlock: ansi.StyleCodeBlock{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color: palette.rose,
				},
				Margin: uintPtr(defaultMargin),
			},
			Theme: palette.chromaTheme,
		},
		Table: ansi.StyleTable{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{},
			},
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: "\n🠶 ",
		},
	}
}
