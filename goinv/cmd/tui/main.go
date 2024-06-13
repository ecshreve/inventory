package main

import (
	"context"
	"fmt"
	"goinv/ent"
	"os"
	"strconv"
	"strings"

	entItem "goinv/ent/item"

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

type model struct {
	form  *huh.Form
	width int
}

func newModel() model {
	// Make initial list of items
	allItems := client.Item.Query().AllX(context.Background())
	options := make([]string, len(allItems))
	for i, goItem := range allItems {
		options[i] = goItem.Name + " (" + strconv.Itoa(goItem.Quantity) + ")" + " -- " + goItem.Category.String() + " -- " + goItem.QueryStorageLocation().OnlyX(context.Background()).Name
	}

	return model{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("item").
					Options(huh.NewOptions(options...)...).
					Title("Select item").Height(12),

				huh.NewConfirm().
					Key("confirm").
					Title("Increase Quantity?"),
			),
		),
	}

}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, 120)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		case "enter":

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
		raw := m.form.GetString("item")
		itemName := strings.Split(raw, " (")[0]
		increaseQuantity(itemName)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.form.State {
	case huh.StateNormal:
		return m.form.View()
	case huh.StateCompleted:
		return appStyle.Render(
			titleStyle.Render("Success!") + "\n\n" +
				statusMessageStyle("Quantity increased!"),
		)
	default:
		return m.form.View()
	}
}

func main() {
	var err error
	client, err = ent.Open(dialect.SQLite, SQLITE_DB)
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// increaseQuantity increases the quantity of the selected item.
func increaseQuantity(itemName string) {
	// Get the item from the database.
	goItem := client.Item.Query().Where(entItem.NameEQ(itemName)).OnlyX(context.Background())

	// Increase the quantity of the item.
	goItem.Update().SetQuantity(goItem.Quantity + 1).SaveX(context.Background())
}
