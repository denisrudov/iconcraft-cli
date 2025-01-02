package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strings"
)

//go:embed iconsPath/*
var iconsPath embed.FS

type Search struct {
	icons []*Icon
	path  embed.FS
}

func (s *Search) Perform(s2 string) {

	dir, err := s.path.ReadDir("iconsPath")
	if err != nil {
		log.Fatal(err)
	}

	var found []fs.FileInfo

	for _, d := range dir {
		extension := strings.Split(d.Name(), ".")[1]
		if extension == "json" {
			found = append(found, d.(fs.FileInfo))
			fmt.Println(extension, d.Name())
		}
	}
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

	for _, d := range dir {
		extension := strings.Split(d.Name(), ".")[1]
		if extension == "json" {

			//found = append(found, d.(fs.FileInfo))
			//fmt.Println(extension, d.Name())
		}
	}

	return
}
