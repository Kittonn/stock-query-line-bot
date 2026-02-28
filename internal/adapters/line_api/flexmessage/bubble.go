package flexmessage

type BubbleOption func(Component)

func NewBubble(opts ...BubbleOption) Component {
	c := Component{
		"type": "bubble",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithBubbleHeader(header Component) BubbleOption {
	return func(c Component) {
		c["header"] = header
	}
}

func WithBubbleHero(hero Component) BubbleOption {
	return func(c Component) {
		c["hero"] = hero
	}
}

func WithBubbleBody(body Component) BubbleOption {
	return func(c Component) {
		c["body"] = body
	}
}

func WithBubbleFooter(footer Component) BubbleOption {
	return func(c Component) {
		c["footer"] = footer
	}
}

func WithBubbleStyles(styles Component) BubbleOption {
	return func(c Component) {
		c["styles"] = styles
	}
}

func WithBubbleSize(size string) BubbleOption {
	return func(c Component) {
		c["size"] = size
	}
}
