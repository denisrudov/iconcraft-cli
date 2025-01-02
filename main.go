package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

var version = "dev"

func main() {
	displayWelcomeMessage()

	for {
		iconName := getUserInput("Search Icon")
		if iconName == "" {
			return
		}

		search := NewSearch()
		icons := search.Perform(iconName)

		if len(icons) == 0 {
			fmt.Println("No icons found. Try again.")
			continue
		}

		handleIconSelection(icons)
	}
}

func displayWelcomeMessage() {
	fmt.Printf(`Welcome to IconCraft CLI (v.%s) - Your Go-To Tool for Lucide Icons!

Effortlessly search and explore a vast database of icons from [Lucide](https://lucide.dev/).

📂 **Features:**  
- 🔍 **Search:** Quickly find icons by name or keyword.  
- 📜 **Explore:** View the names of matching icons in an easy-to-read list.  
- 📋 **Export:** Copy your chosen icon's code in your preferred framework:  
  - JSX  
  - React  
  - Angular  
  - Vue  
  - Svelte  

Start now and supercharge your workflow with the perfect icon for your project!
Ctrl-c Back / Exit
`, version)
}

func getUserInput(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(result)
}

func handleIconSelection(icons []*Icon) {
	iconNames := extractIconNames(icons)

	for {
		selectedIndex, _, err := promptSelection("Select Icon", iconNames)
		if err != nil {
			return
		}

		selectedIcon := icons[selectedIndex]
		selectedIcon.RenderInConsole()

		action, err := promptUserAction()
		if err != nil {
			continue
		}

		handleIconAction(selectedIcon, action)
		time.Sleep(2 * time.Second)
	}
}

func extractIconNames(icons []*Icon) []string {
	names := make([]string, len(icons))
	for i, icon := range icons {
		names[i] = icon.Name
	}
	return names
}

func promptSelection(label string, items []string) (int, string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	return prompt.Run()
}

func promptUserAction() (string, error) {
	prompt := promptui.Select{
		Label: "Copy in Clipboard",
		Items: IconActions,
	}
	_, action, err := prompt.Run()
	if errors.Is(err, promptui.ErrInterrupt) {
		return "", err
	}
	return action, nil
}

func handleIconAction(icon *Icon, action string) {
	code := icon.GetAction(action)()

	if err := copyToClipboard(code); err != nil {
		fmt.Println("Failed to copy to clipboard.")
	} else {
		fmt.Println("Copied to clipboard!")
	}
}
