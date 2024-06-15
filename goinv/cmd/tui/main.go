package main

import (
	"context"
	"fmt"
	"goinv/ent"
	"os"
	"strconv"
	"strings"

	entItem "goinv/ent/item"
	"goinv/ent/storagelocation"

	"entgo.io/ent/dialect"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

var SQLITE_DB = "file:file.db?mode=rwc&cache=shared&_fk=1"
var client *ent.Client

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	detailStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#25A065")).
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

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
	Help lipgloss.Style
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
	return &s
}

type formType int

const (
	formInitial formType = iota
	formSelectItem
	formCreateItem
)

type model struct {
	form        *huh.Form
	width       int
	height      int
	currentForm formType
	status      string
	styles      *Styles
	lg          *lipgloss.Renderer
	actionLog   []string
}

func newModel() model {
	m := model{
		form:        createInitialForm(),
		currentForm: formInitial,
		width:       80,
		height:      20,
		status:      "",
		actionLog: []string{
			"initializing...",
		},
	}

	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	return m
}

func createInitialForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("action").
				Options(
					huh.NewOptions(
						"Select",
						"Create",
					)...,
				).
				Title("What do you want to do?"),
		).WithShowErrors(false).WithShowHelp(false),
	)
}

func createSelectItemForm() *huh.Form {
	allItems, err := client.Item.Query().All(context.Background())
	if err != nil {
		log.Error(err)
		return nil
	}

	options := make([]string, len(allItems))
	for i, goItem := range allItems {
		locStr, err := goItem.QueryStorageLocation().Only(context.Background())
		if err != nil {
			log.Error(err)
			return nil
		}
		options[i] = goItem.Name + " (" + strconv.Itoa(goItem.Quantity) + ")" + " -- " + goItem.Category.String() + " -- " + locStr.Name
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("item").
				Options(huh.NewOptions(options...)...).
				Title("Select item").Height(12),

			huh.NewConfirm().
				Key("confirm").
				Title("Increase Quantity?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("nope, try again")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("No"),
		).WithShowErrors(false).WithShowHelp(false),
	)
}

func createNewItemForm() *huh.Form {
	allCategories := []string{"cable", "adapter", "device", "misc", "unknown"}

	allLocations := client.StorageLocation.Query().AllX(context.Background())
	locationOptions := make([]string, len(allLocations))
	for i, loc := range allLocations {
		locationOptions[i] = loc.Name
	}

	quantityValue := "1"
	loc := ""

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Title("Item Name"),
			huh.NewSelect[string]().
				Key("category").
				Title("Item Category").
				Options(huh.NewOptions(allCategories...)...),
			huh.NewInput().
				Key("quantity").
				Title("Quantity").Value(&quantityValue),
			huh.NewSelect[string]().
				Key("location").
				Title("Location").
				Value(&loc).
				Options(huh.NewOptions(locationOptions...)...),
		).WithShowErrors(false).WithShowHelp(false),
		huh.NewGroup(
			huh.NewConfirm().
				Key("confirm").
				Title("Create Item?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("NO!"),
		).WithShowErrors(false).WithShowHelp(false),
	)
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, 80)
		m.height = min(msg.Height, 40)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	switch m.currentForm {
	case formInitial:
		if m.form.State == huh.StateCompleted {
			action := m.form.GetString("action")
			m.status = ""

			if action == "Select" {
				m.currentForm = formSelectItem
				m.form = createSelectItemForm()
			} else if action == "Create" {
				m.currentForm = formCreateItem
				m.form = createNewItemForm()
			}
			m.form.Init()
		}
	case formSelectItem:
		if m.form.State == huh.StateCompleted {
			raw := m.form.GetString("item")
			itemName := strings.Split(raw, " (")[0]
			result := increaseQuantity(itemName)
			m.status = result
			m.actionLog = append(m.actionLog, result)

			m.currentForm = formInitial
			m.form = createInitialForm()
			m.form.Init()
		}
	case formCreateItem:
		if m.form.State == huh.StateCompleted {
			name := m.form.GetString("name")
			category := m.form.GetString("category")
			quantityStr := m.form.GetString("quantity")
			location := m.form.GetString("location")

			quantity, err := strconv.Atoi(quantityStr)
			if err != nil {
				log.Print("Invalid quantity, please enter a number")
				return m, nil
			}

			result := createNewItem(name, category, quantity, location)
			m.status = result
			m.currentForm = formInitial
			m.actionLog = append(m.actionLog, result)

			m.form = createInitialForm()
			m.form.Init()
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.form.State {
	case huh.StateCompleted:
		return m.appBoundaryView("Done!")
	default:
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		actionLogView := m.lg.NewStyle().Margin(1, 0).Width(40).Height(m.height / 2).Border(lipgloss.RoundedBorder()).Render(strings.Join(m.actionLog, "\n---\n"))

		errors := m.form.Errors()
		header := m.appBoundaryView(m.status)
		if len(errors) > 0 {
			header = m.appErrorBoundaryView(m.errorView())
		}
		body := lipgloss.JoinHorizontal(
			lipgloss.Top,
			form,
			actionLogView,
		)
		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))

		return lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			body,
			footer,
		)
	}
}

func (m model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

func (m model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

func main() {
	var err error
	client, err = ent.Open(dialect.SQLite, SQLITE_DB)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func increaseQuantity(itemName string) string {
	ctx := context.Background()
	item, err := client.Item.Query().Where(entItem.NameEQ(itemName)).Only(ctx)
	if err != nil {
		return err.Error()
	}

	updatedItem, err := client.Item.
		UpdateOne(item).
		SetQuantity(item.Quantity + 1).
		Save(ctx)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Increased quantity of %s to %d", updatedItem.Name, updatedItem.Quantity)
}

// createNewItem handles the creation of a new item in the database.
func createNewItem(name, category string, quantity int, location string) string {
	ctx := context.Background()

	// Get the storage location.
	locEnt := client.StorageLocation.Query().Where(storagelocation.NameEQ(location)).OnlyX(ctx)
	if locEnt == nil {
		return fmt.Sprintf("storage location %s not found", location)
	}

	// Create the new item.
	createdItem, err := client.Item.Create().
		SetCategory(entItem.Category(category)).
		SetName(name).
		SetQuantity(quantity).
		SetStorageLocation(locEnt).
		Save(ctx)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("created item: %s", createdItem.Name)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
