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

type formType int

const (
	formSelectItem formType = iota
	formCreateItem
)

type model struct {
	form        *huh.Form
	width       int
	height      int
	currentForm formType
}

func newModel() model {
	// Make initial list of items
	allItems := client.Item.Query().AllX(context.Background())
	options := make([]string, len(allItems))
	for i, goItem := range allItems {
		options[i] = goItem.Name + " (" + strconv.Itoa(goItem.Quantity) + ")" + " -- " + goItem.Category.String() + " -- " + goItem.QueryStorageLocation().OnlyX(context.Background()).Name
	}

	return model{
		form:        createSelectItemForm(),
		currentForm: formSelectItem,
	}

}

func createSelectItemForm() *huh.Form {
	allItems := client.Item.Query().AllX(context.Background())
	options := make([]string, len(allItems))
	for i, goItem := range allItems {
		options[i] = goItem.Name + " (" + strconv.Itoa(goItem.Quantity) + ")" + " -- " + goItem.Category.String() + " -- " + goItem.QueryStorageLocation().OnlyX(context.Background()).Name
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
		),
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
				Options(huh.NewOptions(locationOptions...)...),
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
		).WithHeight(20).WithShowHelp(true),
	)
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, 120)
		m.height = min(msg.Height, 40)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		case "n":
			if m.currentForm == formSelectItem {
				m.currentForm = formCreateItem
				m.form = createNewItemForm()
				return m, m.form.Init()
			}
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
	case formSelectItem:
		if m.form.State == huh.StateCompleted {
			raw := m.form.GetString("item")
			itemName := strings.Split(raw, " (")[0]
			increaseQuantity(itemName)

			m.form = createSelectItemForm()
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

			createNewItem(name, category, quantity, location)

			m.currentForm = formSelectItem
			m.form = createSelectItemForm()
			m.form.Init()
		}
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
				statusMessageStyle("Action completed!"),
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

	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
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

// createNewItem handles the creation of a new item in the database.
func createNewItem(name, category string, quantity int, location string) {
	ctx := context.Background()

	// Get the storage location.
	locEnt := client.StorageLocation.Query().Where(storagelocation.NameEQ(location)).OnlyX(ctx)
	if locEnt == nil {
		log.Fatalf("storage location not found: %s", location)
	}

	// Create the new item.
	client.Item.Create().
		SetCategory(entItem.Category(category)).
		SetName(name).
		SetQuantity(quantity).
		SetStorageLocation(locEnt).
		SaveX(ctx)

	log.Info("created item:", name)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
