package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"goinv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const maxWidth = 80

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
	TableHeader,
	TableRow lipgloss.Style
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
	s.TableHeader = lg.NewStyle().
		Foreground(lipgloss.Color("15")). // White text
		Background(lipgloss.Color("7")).  // Gray background
		Bold(true).
		Padding(0, 1)
	s.TableRow = lg.NewStyle().
		Padding(0, 1)
	return &s
}

type state int

const (
	statusNormal state = iota
	stateDone
)

type Model struct {
	state     state
	lg        *lipgloss.Renderer
	styles    *Styles
	form      *huh.Form
	width     int
	inventory []goinv.Item
}

func NewModel() Model {
	m := Model{width: maxWidth, inventory: getInventory()}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Placeholder("Item Name").
				Title("Item Name").
				Prompt("Name: "),

			huh.NewInput().
				Key("qty").
				Placeholder("Quantity").
				Title("Quantity").
				Prompt("Qty: ").
				Validate(func(value string) error {
					if value == "" || !isAllDigits(value) {
						return fmt.Errorf("Invalid quantity")
					}
					return nil
				}),

			huh.NewSelect[string]().
				Key("category").
				Options(huh.NewOptions(
					string(goinv.Cable),
					string(goinv.Adapter),
					string(goinv.Device),
					string(goinv.Misc),
					string(goinv.Unknown),
				)...).
				Title("Category"),

			huh.NewSelect[string]().
				Key("location").
				Options(huh.NewOptions(
					string(goinv.HalfCrate_White_1),
					string(goinv.HalfCrate_White_2),
					string(goinv.FullCrate_Black_1),
					string(goinv.FullCrate_Black_2),
					string(goinv.FullCrate_Gray_1),
					string(goinv.HalfCrate_Gray_1),
					string(goinv.HalfCrate_Gray_2),
				)...).
				Title("Location"),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("Please complete all fields")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("No"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)

	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
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

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Capture form data and create a new item
		name := m.form.GetString("name")
		qty, _ := strconv.ParseUint(m.form.GetString("qty"), 10, 64)
		category := goinv.ItemCategory(m.form.GetString("category"))
		location := goinv.StorageLocation(m.form.GetString("location"))

		newItem := goinv.Item{
			Name:     name,
			Qty:      uint(qty),
			Category: category,
			Location: location,
		}

		// Add new item to inventory
		m.inventory = append(m.inventory, newItem)

		// Print the item details to the console (for debugging)
		fmt.Println(newItem)

		// Reset form or quit when the form is done.
		// For now, we quit after creating an item.
		return m, tea.Quit
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	var formContent string
	switch m.form.State {
	case huh.StateCompleted:
		formContent = s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(
			fmt.Sprintf("%s\n\n%s", s.Highlight.Render("New Item Created!"), "Successfully created the new item. Please check your inventory database for the entry."),
		) + "\n\n"
	default:
		if m.form.State == huh.StateAborted {
			formContent = s.ErrorHeaderText.Render("Form submission failed. Please correct the errors and try again.") + "\n\n"
		} else {
			formContent = m.form.View()
		}
	}

	formView := s.Base.Render(
		s.HeaderText.Render("Create a New Item") + "\n\n" +
			formContent +
			"\n\n" + s.Help.Render("(esc to quit, enter to submit)") + "\n",
	)

	inventoryView := renderInventory(m)

	return lipgloss.JoinHorizontal(lipgloss.Top, formView, inventoryView)
}

func renderInventory(m Model) string {
	s := m.styles

	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		s.TableHeader.Render("Item Name"),
		s.TableHeader.Render("Quantity"),
		s.TableHeader.Render("Category"),
		s.TableHeader.Render("Location"),
	)

	rows := []string{header}
	for _, item := range m.inventory {
		rows = append(rows, lipgloss.JoinHorizontal(
			lipgloss.Top,
			s.TableRow.Render(item.Name),
			s.TableRow.Render(strconv.Itoa(int(item.Qty))),
			s.TableRow.Render(string(item.Category)),
			s.TableRow.Render(string(item.Location)),
		))
	}

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")). // Gray border
		MarginLeft(2).
		Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func isAllDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// getInventory fetches the inventory via the API
func getInventory() []goinv.Item {
	resp, err := http.Get("http://localhost:8080/items")
	if err != nil {
		log.Fatal("Failed to fetch inventory:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Failed to fetch inventory:", err)
		return nil
	}

	var inventory []goinv.Item
	err = json.NewDecoder(resp.Body).Decode(&inventory)
	if err != nil {
		log.Fatal("Failed to decode inventory:", err)
		return nil
	}

	return inventory
}

func main() {
	_, err := tea.NewProgram(NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
