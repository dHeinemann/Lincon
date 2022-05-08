/*
 * Lincon, a hyperlink converter
 * Copyright (C) 2022 David Heinemann
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"dheinemann.com/lincon/path"
	"github.com/antchfx/htmlquery"
)

var addMissingExtensions bool = true

func main() {
	pathFlag := flag.String("path", "", "Path to file (or directory) to convert.")
	basePathFlag := flag.String("base", "", "Path to base (root) of website. Only used when PATH is a file.")
	flag.Parse()

	if *pathFlag == "" {
		printHelp()
		return
	}

	fileInfo, err := os.Stat(*pathFlag)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if fileInfo.IsDir() {
		convertAll(*pathFlag)
	} else if canConvert(fileInfo) {
		convertLinks(*pathFlag, *basePathFlag)
	}
}

func printHelp() {
	fmt.Println("Lincon v0.1.0")
	fmt.Println()
	fmt.Println("usage: lincon -path=PATH [-base=BASE]")
	fmt.Println()
	flag.PrintDefaults()
}

func canConvert(fileInfo os.FileInfo) bool {
	return !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), "html")
}

func convertAll(basePath string) {
	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if canConvert(fileInfo) {
			convertLinks(path, basePath)
		}
		return nil
	})
}

func convertLinks(filePath string, basePath string) {
	doc, err := htmlquery.LoadDoc(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, target := range targets {
		targetNodes := htmlquery.Find(doc, target.getSelector())
		for _, node := range targetNodes {
			for i := 0; i < len(node.Attr); i++ {
				nodeVal := node.Attr[i].Val
				if node.Attr[i].Key == target.attribute {
					if len(nodeVal) > 0 && strings.HasPrefix(nodeVal, "/") {
						uri := path.ConvertToRelative(node.Attr[i].Val, filePath, basePath)
						if path.MissingExtension(uri) {
							uri = path.RestoreExtension(uri)
						}

						node.Attr[i].Val = uri
					}
				}
			}
		}
	}

	fmt.Println(filePath)
	f, err := os.OpenFile(filePath, os.O_RDWR, os.ModeAppend)
	defer f.Close()

	f.WriteString(htmlquery.OutputHTML(doc, true))
}
