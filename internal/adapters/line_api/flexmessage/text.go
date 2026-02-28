package flexmessage

type TextOption func(Component)

func NewText(text string, opts ...TextOption) Component {
	c := Component{
		"type": "text",
		"text": text,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithTextAlign(align string) TextOption {
	return func(c Component) {
		c["align"] = align
	}
}

func WithTextColor(color string) TextOption {
	return func(c Component) {
		c["color"] = color
	}
}

func WithTextSize(size string) TextOption {
	return func(c Component) {
		c["size"] = size
	}
}

func WithTextWeight(weight string) TextOption {
	return func(c Component) {
		c["weight"] = weight
	}
}

func WithTextGravity(gravity string) TextOption {
	return func(c Component) {
		c["gravity"] = gravity
	}
}

func WithTextMargin(margin string) TextOption {
	return func(c Component) {
		c["margin"] = margin
	}
}

func WithTextFlex(flex int) TextOption {
	return func(c Component) {
		c["flex"] = flex
	}
}
