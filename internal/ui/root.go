package ui

type (
	page int
	size int
)

const (
	viewPage page = iota
)

type state struct {}

type model struct {
	page  page
	state state
}
