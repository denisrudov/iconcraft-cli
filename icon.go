package main

import (
	"fmt"
	"strings"
)

var IconActions = []string{"JSX", "Component Name", "Vue", "Svelte", "Angular"}

type IconSchema struct {
	Schema       string   `json:"$schema"`
	Contributors []string `json:"contributors"`
	Tags         []string `json:"tags"`
	Categories   []string `json:"categories"`
}

type Icon struct {
	Id       int
	Name     string
	Filename string
	Svg      string
	Schema   *IconSchema
}

func (i *Icon) RenderInConsole() {
	svgString := strings.ReplaceAll(i.Svg, "currentColor", "#FFFFFF")
	toImage, _ := renderSVGToImage([]byte(svgString), 40, 40)
	renderInConsole(toImage)
}

func (i *Icon) GetAction(s string) func() string {
	var actions map[string]func() string = map[string]func() string{
		IconActions[0]: i.GetJSX,
		IconActions[1]: func() string {
			return i.Name
		},
		IconActions[2]: i.GetVue,
		IconActions[3]: i.GetSvelte,
		IconActions[4]: i.GetAngular,
	}

	return actions[s]
}

func (i *Icon) GetJSX() string {
	return fmt.Sprintf("<%s />", i.Name)
}

func (i *Icon) GetVue() string {
	return fmt.Sprintf("<%s />", i.Name)
}

func (i *Icon) GetSvelte() string {
	return fmt.Sprintf("<%s />", i.Name)
}

func (i *Icon) GetAngular() string {
	return fmt.Sprintf(`<lucide-icon name="%s"></lucide-icon>`, i.Filename)
}

func (i *Icon) matches(searchString string) bool {
	for _, tag := range i.Schema.Tags {
		if strings.Contains(tag, searchString) {
			return true
		}
	}
	return false
}

func NewIcon(id int, name string, svg string, schema *IconSchema) *Icon {
	return &Icon{Id: id, Name: camelCaseFromDash(name), Filename: name, Svg: svg, Schema: schema}

}
