package ChainResponsability

type ComponentWithContextualHelp interface {
	ShowHelp() string
}

type Container struct {
	*Component
	children []ComponentWithContextualHelp
}

func (c *Container) add(child ComponentWithContextualHelp) {
	c.children = append(c.children, child)
	switch v := child.(type) {
	case *Component:
		v.container = c
	}
}

type Component struct {
	tooltipText string
	container   *Container
}

func (c *Component) ShowHelp() string {
	if c.tooltipText != "" {
		return c.tooltipText
	}
	return c.container.ShowHelp()
}

type Dialog struct {
	Container
	wikiPageUrl string
}

func NewDialog(wikiPageUrl string) *Dialog {
	return &Dialog{wikiPageUrl: wikiPageUrl}
}
func (d *Dialog) ShowHelp() string {
	return d.Component.ShowHelp()
}

type Panel struct {
	Container
	modalHelpText string
	x, y, z, w    int
}

func NewPanel(modalHelpText string, x, y, z, w int) *Panel {
	return &Panel{modalHelpText: modalHelpText, x: x, y: y, z: z, w: w}
}

func (p *Panel) ShowHelp() string {
	return p.modalHelpText
}

func buildPanel() *Panel {
	dialog := NewDialog("https://www.wikipedia.com")
	panel := NewPanel("Panel Help", 0, 0, 400, 800)
	okButton := Component{tooltipText: "This is an OK button that..."}
	cancelButton := Component{tooltipText: "Cancel"}

	panel.add(&okButton)
	panel.add(&cancelButton)
	panel.add(dialog)
	return panel
}
