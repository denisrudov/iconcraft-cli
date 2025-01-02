package main

type IconSchema struct {
	Schema       string   `json:"$schema"`
	Contributors []string `json:"contributors"`
	Tags         []string `json:"tags"`
	Categories   []string `json:"categories"`
}

type Icon struct {
	Id      string
	Name    string
	Svg     string
	JSX     string
	Vue     string
	Svelte  string
	Angular string
	Color   string
	Schema  IconSchema
}

func NewIcon(id string, name string, svg string, JSX string, vue string, svelte string, angular string, color string, schema IconSchema) *Icon {
	return &Icon{Id: id, Name: name, Svg: svg, JSX: JSX, Vue: vue, Svelte: svelte, Angular: angular, Color: color, Schema: schema}

}
