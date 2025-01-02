package main

type IconSchema struct {
	Schema       string   `json:"$schema"`
	Contributors []string `json:"contributors"`
	Tags         []string `json:"tags"`
	Categories   []string `json:"categories"`
}

type Icon struct {
	Id     int
	Name   string
	Svg    string
	Schema *IconSchema
}

func NewIcon(id int, name string, svg string, schema *IconSchema) *Icon {
	return &Icon{Id: id, Name: camelCaseFromDash(name), Svg: svg, Schema: schema}

}
