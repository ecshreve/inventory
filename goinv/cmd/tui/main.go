package main

import (
	"fmt"
	"goinv"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var inventory goinv.Inventory

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help,
	ListItem lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	s.ListItem = lg.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(0, 1, 0, 1)
	return &s
}

type item goinv.Item

func (i item) Title() string       { return fmt.Sprintf("%d -- %s", i.Qty, i.Name) }
func (i item) Description() string { return fmt.Sprintf("%s -- %s", i.Category, i.Location) }
func (i item) FilterValue() string { return fmt.Sprintf("%s %s", i.Name, i.Category) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s -- %s", i.Name, i.Category)

	fn := NewStyles(lipgloss.DefaultRenderer()).ListItem.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

const maxWidth = 120

type state int

const (
	stateMain state = iota
	locationSelected
	categorySelected
	itemSelected
)

type Model struct {
	state    state
	lg       *lipgloss.Renderer
	styles   *Styles
	width    int
	errors   []error
	list     list.Model
	invItems map[string]goinv.Item
}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.invItems = m.getItemsFromInventory()

	listItems := make([]list.Item, 0, len(m.invItems))
	for _, invItem := range m.invItems {
		listItems = append(listItems, item(invItem))
	}
	m.list = list.New(listItems, itemDelegate{}, 40, 40)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Do stuff here
	var cmd tea.Cmd

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

type genericItem struct {
	name string
}

func (g genericItem) FilterValue() string { return g.name }

func (m Model) View() string {
	s := m.styles

	categories := make(map[string]struct{})
	for _, item := range m.invItems {
		categories[string(item.Category)] = struct{}{}
	}

	var categoriesItems []list.Item
	for category := range categories {
		categoriesItems = append(categoriesItems, genericItem{name: category})
	}
	categoriesList := list.New(categoriesItems, list.NewDefaultDelegate(), 40, 40)

	errors := m.errors
	header := m.appBoundaryView("Inventory Management System")
	if len(errors) > 0 {
		header = m.appErrorBoundaryView(m.errorView())
	}
	body := lipgloss.JoinHorizontal(lipgloss.Top, categoriesList.View(), m.list.View(), s.Status.Render(""))

	footer := m.appBoundaryView("Press 'q' to quit")
	if len(errors) > 0 {
		footer = m.appErrorBoundaryView("")
	}

	return s.Base.Render(header + "\n" + body + "\n\n" + footer)
}

func (m Model) errorView() string {
	var s string
	for _, err := range m.errors {
		s += err.Error()
	}
	return s
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

// getItemsFromInventory fetches items from the inventory and returns them
// as a slice of list items.
func (m Model) getItemsFromInventory() map[string]goinv.Item {
	items, err := inventory.GetItems()
	if err != nil {
		m.errors = append(m.errors, err)
		return nil
	}

	invItems := make(map[string]goinv.Item)
	for _, item := range items {
		invItems[item.Name] = item
	}

	return invItems
}

func main() {
	var err error
	inventory, err = goinv.NewGormInventory()
	if err != nil {
		fmt.Println("Failed to initialize inventory:", err)
		os.Exit(1)
	}

	_, err = tea.NewProgram(NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
