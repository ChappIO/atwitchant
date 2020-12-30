package command

type Command struct {
	Name        string
	Description string
	Flags       func()
	Validate    func() bool
	Run         func()
}
