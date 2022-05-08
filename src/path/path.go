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

package path

import "strings"

func IsRelative(uri string) bool {
	return !strings.HasPrefix(uri, "/")
}

func ConvertToRelative(uri string, currentPath string, basePath string) string {
	if IsRelative(uri) {
		return uri
	}

	targetPath := uri[1:] // Remove leading slash
	targetPathParts := strings.Split(targetPath, "/")
	targetFileName := targetPathParts[len(targetPathParts)-1]
	targetPathParts = targetPathParts[:len(targetPathParts)-1] // Remove filename

	currentPath = strings.Replace(currentPath, basePath, "", 1)
	currentPathParts := strings.Split(currentPath, "/")
	currentPathParts = currentPathParts[:len(currentPathParts)-1] // Remove filename

	resultPathParts := []string{}

	forkIndex := 0 // Index where targetPath diverges from currentPath

	// Strategy:
	// 1.  Iterate over both paths from the start, popping the parent-most directory of each
	// 2.  Stop when we find the first directory that doesn't match. This is the point where we either need to start
	//     backtracking, or looking forward.
	// 2a. If the first directory that doesn't match is the directory of the current file, then the target is "ahead",
	//     inside another descendent.
	// 2b. Otherwise we need to start backtrack a number of folders equal to the number of directories between the
	//     current file and the last match.

	foundCommonDir := false
	currentPathLen := len(currentPathParts)
	targetPathLen := len(targetPathParts)
	for i := 0; i < targetPathLen; i++ {
		if i <= currentPathLen-1 || targetPathParts[i] == currentPathParts[i] {
			foundCommonDir = true
			if i == currentPathLen-1 {
				// We've reached the directory of the current file. Target is in a descendent directory.
				resultPathParts = targetPathParts[i+1:]
				forkIndex = i + 1
				goto exit
			} else if i == targetPathLen-1 {
				// We've reached the end of the target, and it's not the directory of the current file. The target is in
				// a parent directory.

				forkIndex = i + 1
				for len(resultPathParts) < currentPathLen-targetPathLen {
					resultPathParts = append(resultPathParts, "..")
				}
				goto exit
			}

			continue
		} else {
			// We've reached a fork in the paths. Need to start backtracking.
			forkIndex = i + 1
			for k := 0; k < currentPathLen; k++ {
				for len(resultPathParts) < currentPathLen-len(resultPathParts) {
					resultPathParts = append(resultPathParts, "..")
				}
				resultPathParts = append(resultPathParts, targetPathParts[i:]...)
				goto exit
			}
		}
	}

	// No common directory found: need to go all the way up to the base directory
	if !foundCommonDir {
		for k := 0; k < len(currentPathParts)-len(targetPathParts); k++ {
			resultPathParts = append(resultPathParts, "..")
		}
	}

exit:

	var startIndex int
	if foundCommonDir {
		startIndex = forkIndex + 1
	} else {
		startIndex = forkIndex
	}

	for i := startIndex; i < len(targetPathParts); i++ {
		resultPathParts = append(resultPathParts, targetPathParts[i])
	}
	resultPathParts = append(resultPathParts, targetFileName)

	resultPath := strings.Join(resultPathParts, "/")
	resultPath = strings.ReplaceAll(resultPath, "?", "%3F")
	return resultPath
}

func MissingExtension(uri string) bool {
	var i int
	if strings.Contains(uri, "/") {
		i = strings.LastIndex(uri, "/")
	} else {
		i = 0
	}

	return !strings.Contains(uri[i:], ".")
}

func RestoreExtension(uri string) string {
	return uri + ".html"
}
