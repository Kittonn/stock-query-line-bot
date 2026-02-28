package flexmessage

type ButtonOption func(Component)

func NewButton(action Component, opts ...ButtonOption) Component {
	c := Component{
		"type":   "button",
		"action": action,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithButtonStyle(style string) ButtonOption {
	return func(c Component) {
		c["style"] = style
	}
}

func WithButtonHeight(height string) ButtonOption {
	return func(c Component) {
		c["height"] = height
	}
}

func WithButtonColor(color string) ButtonOption {
	return func(c Component) {
		c["color"] = color
	}
}

func NewURIAction(label, uri string) Component {
	return Component{
		"type":  "uri",
		"label": label,
		"uri":   uri,
	}
}
