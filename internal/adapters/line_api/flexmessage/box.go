package flexmessage

type BoxOption func(Component)

func NewBox(layout string, contents []Component, opts ...BoxOption) Component {
	c := Component{
		"type":     "box",
		"layout":   layout,
		"contents": contents,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithBoxMargin(margin string) BoxOption {
	return func(c Component) {
		c["margin"] = margin
	}
}

func WithBoxSpacing(spacing string) BoxOption {
	return func(c Component) {
		c["spacing"] = spacing
	}
}

func WithBoxAlignItems(alignItems string) BoxOption {
	return func(c Component) {
		c["alignItems"] = alignItems
	}
}

func WithBoxPaddingAll(padding string) BoxOption {
	return func(c Component) {
		c["paddingAll"] = padding
	}
}

func WithBoxPaddingBottom(padding string) BoxOption {
	return func(c Component) {
		c["paddingBottom"] = padding
	}
}

func WithBackgroundColor(color string) BoxOption {
	return func(c Component) {
		c["backgroundColor"] = color
	}
}
