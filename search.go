package main

import (
	"embed"
	"encoding/json"
	"log"
	"path/filepath"
	"strings"
)

//go:embed icons/*
var iconsPath embed.FS

type Search struct {
	icons []*Icon
}

func NewSearch() *Search {
	icons, err := initializeIcons()
	if err != nil {
		log.Fatalf("Failed to initialize icons: %v", err)
	}
	return &Search{icons: icons}
}

func (s *Search) Perform(searchString string) []*Icon {
	var matchedIcons []*Icon
	for _, icon := range s.icons {
		if icon.Matches(searchString) {
			matchedIcons = append(matchedIcons, icon)
		}
	}
	return matchedIcons
}

func initializeIcons() ([]*Icon, error) {
	dirEntries, err := iconsPath.ReadDir("icons")
	if err != nil {
		return nil, err
	}

	var icons []*Icon
	id := 0
	for _, entry := range dirEntries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			icon, err := loadIcon(id+1, entry.Name())
			if err == nil {
				icons = append(icons, icon)
			}
		}
	}

	return icons, nil
}

func loadIcon(id int, jsonFileName string) (*Icon, error) {
	baseName := strings.TrimSuffix(jsonFileName, filepath.Ext(jsonFileName))
	schemaData, err := iconsPath.ReadFile(filepath.Join("icons", jsonFileName))
	if err != nil {
		return nil, err
	}

	svgData, err := iconsPath.ReadFile(filepath.Join("icons", baseName+".svg"))
	if err != nil {
		return nil, err
	}

	var schema *IconSchema
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, err
	}

	return NewIcon(id, baseName, string(svgData), schema), nil
}
