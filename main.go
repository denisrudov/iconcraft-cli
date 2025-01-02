package main

import (
	"errors"
	"github.com/manifoldco/promptui"
	"strings"
	"time"
)

func main() {

	println(`Welcome to IconCraft CLI - Your Go-To Tool for Lucide Icons!  
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

Start now and supercharge your workflow with the perfect icon for your project!`)

begin:
	for {
		//ClearConsole()
		prompt := promptui.Prompt{
			Label: "Select Icon",
		}

		result, err := prompt.Run()
		if err != nil {
			return
		}

		if len(strings.TrimSpace(result)) == 0 {
			continue
		}

		search := NewSearch()
		icons := search.Perform(result)

		if len(icons) > 0 {
			var items []string
			// items = append(items, "←")
			for _, icon := range icons {
				items = append(items, icon.Name)
			}
			for {
				sPrompt := promptui.Select{
					Label: "Select Icon",
					Items: items,
				}
				selectedIndex, _, err := sPrompt.Run()
				if err != nil {
					return
				}

				icon := icons[selectedIndex]
				icon.RenderInConsole()

				sPrompt = promptui.Select{
					Label: "Copy in Clipboard",
					Items: IconActions,
				}

				_, s, err := sPrompt.Run()

				if errors.Is(err, promptui.ErrInterrupt) {
					continue
				}

				code := icon.GetAction(s)()

				err = copyToClipboard(code)
				if err != nil {
					println("failed to copy to clipboard")
				} else {
					println("Copied to clipboard...	")
				}

				time.Sleep(time.Second * 2)

				goto begin

			}

		}

	}

}
