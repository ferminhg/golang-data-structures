package command

type CommandOp interface {
	SaveBackup()
	Undo()
	Execute() bool
}

type Command struct {
	app    Application
	editor Editor
	backup string
}

func NewCommand(app Application, editor Editor) *Command {
	return &Command{
		app:    app,
		editor: editor,
	}
}

func (c *Command) SaveBackup() {
	c.backup = c.editor.text
}

func (c *Command) Undo() {
	c.editor.text = c.backup
}

func (c *Command) Execute() bool {
	panic("implement me")
	return false
}

type CopyCommand struct {
	*Command
}

func NewCopyCommand(app Application, editor Editor) *CopyCommand {
	return &CopyCommand{
		Command: &Command{
			app:    app,
			editor: editor,
		},
	}
}

func (c *CopyCommand) Execute() bool {
	c.app.clipboard = c.editor.GetSelection()
	return false
}

type CutCommand struct {
	*Command
}

func NewCutCommand(app Application, editor Editor) *CutCommand {
	return &CutCommand{
		Command: &Command{
			app:    app,
			editor: editor,
		},
	}
}

func (c *CutCommand) Execute() bool {
	c.SaveBackup()
	c.app.clipboard = c.editor.GetSelection()
	c.editor.DeleteSelection()
	return true
}

type PasteCommand struct {
	*Command
}

func NewPasteCommand(app Application, editor Editor) *PasteCommand {
	return &PasteCommand{
		Command: &Command{
			app:    app,
			editor: editor,
		},
	}
}

func (c *PasteCommand) Execute() bool {
	c.SaveBackup()
	c.editor.ReplaceSelection(c.app.clipboard)
	return true
}

type Application struct {
	clipboard    string
	editors      []Editor
	activeEditor Editor
	history      CommandHistory
}

func (a *Application) CreateEditor() {
	a.editors = append(a.editors, Editor{})
	a.activeEditor = a.editors[len(a.editors)-1]
}

func (a *Application) ExecuteCommand(command CommandOp) {
	if command.Execute() {
		a.history.Push(command)
	}
}

func (a *Application) Undo() {
	cmd := a.history.Pop()
	if cmd != nil {
		cmd.Undo()
	}
}

type Editor struct {
	text string
}

func (e *Editor) GetSelection() string {
	return e.text
}
func (e *Editor) DeleteSelection() {
	e.text = ""
}
func (e *Editor) ReplaceSelection(text string) {
	e.text += text
}

type CommandHistory struct {
	history []CommandOp
}

func (c *CommandHistory) Push(cmd CommandOp) {
	c.history = append(c.history, cmd)
}

func (c *CommandHistory) Pop() CommandOp {
	if len(c.history) == 0 {
		return nil
	}
	cmd := c.history[len(c.history)-1]
	c.history = c.history[:len(c.history)-1]
	return cmd
}
