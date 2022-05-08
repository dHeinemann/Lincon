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

import "fmt"

// Represents an HTML node targeted for conversion.
type target struct {
	tag       string
	attribute string
}

// Get an XPath selector.
func (this target) getSelector() string {
	return fmt.Sprintf("//%s[@%s]", this.tag, this.attribute)
}

// List of nodes targeted for conversion.
var targets = [...]target{
	{
		tag:       "link",
		attribute: "href",
	},
	{
		tag:       "script",
		attribute: "src",
	},
	{
		tag:       "a",
		attribute: "href",
	},
	{
		tag:       "img",
		attribute: "src",
	},
}
