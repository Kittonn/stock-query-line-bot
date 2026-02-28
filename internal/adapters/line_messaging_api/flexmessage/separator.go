package flexmessage

type SeperatorOption func(Component)

func NewSeperator(opts ...SeperatorOption) Component {
	c := Component{
		"type": "separator",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithSeparatorMargin(margin string) SeperatorOption {
	return func(c Component) {
		c["margin"] = margin
	}
}

func WithSeparatorColor(color string) SeperatorOption {
	return func(c Component) {
		c["color"] = color
	}
}
