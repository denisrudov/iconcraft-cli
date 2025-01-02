package main

import (
	"embed"
	"encoding/json"
	"log"
	"strings"
)

//go:embed icons/*
var iconsPath embed.FS

type Search struct {
	icons []*Icon
	path  embed.FS
}

func (s *Search) Perform(searchString string) (icons []*Icon) {
	icons = make([]*Icon, 0)

	for _, icon := range s.icons {
		for _, tag := range icon.Schema.Tags {
			if strings.Contains(tag, searchString) {
				icons = append(icons, icon)
			}
		}
	}
	return
}

func NewSearch() *Search {
	return &Search{path: iconsPath, icons: initializeIcons()}
}

func initializeIcons() (icons []*Icon) {
	icons = make([]*Icon, 0)

	dir, err := iconsPath.ReadDir("icons")
	if err != nil {
		log.Fatal(err)
	}

	for id, d := range dir {
		fileSplit := strings.Split(d.Name(), ".")
		extension := fileSplit[1]
		if extension == "json" {
			jsonFile, err := iconsPath.ReadFile("icons/" + d.Name())
			svgFile, svgErr := iconsPath.ReadFile("icons/" + fileSplit[0] + ".svg")
			if err != nil || svgErr != nil {
				continue
			}
			var schema *IconSchema
			err = json.Unmarshal(jsonFile, &schema)
			if err != nil {
				continue
			}
			icons = append(icons, NewIcon(id, fileSplit[0], string(svgFile), schema))
		}
	}

	return
}
