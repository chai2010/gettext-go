// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func gopkgDir(pkgpath string) string {
	cmd := exec.Command("go", "list", "-f", `"{{.Dir}}"`, pkgpath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	output = bytes.TrimSpace(output)
	output = bytes.Trim(output, `"`)
	return string(output)
}

func gopkgName(pkgpath string) string {
	cmd := exec.Command("go", "list", "-f", `"{{.Name}}"`, pkgpath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	output = bytes.TrimSpace(output)
	output = bytes.Trim(output, `"`)
	return string(output)
}

func gopkgPath(pkgpath string) string {
	cmd := exec.Command("go", "list", pkgpath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	output = bytes.TrimSpace(output)
	output = bytes.Trim(output, `"`)
	return string(output)
}

func gopkgFiles(pkgpath string) []string {
	cmd := exec.Command("go", "list", "-f", `"{{.GoFiles}}"`, pkgpath)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// [a.go b.go ...]
	output = bytes.TrimSpace(output)
	output = bytes.Trim(output, `["]`)

	files := strings.Split(string(output), " ")
	for i, s := range files {
		files[i] = s
	}
	sort.Strings(files)
	return files
}

func gopkgFilesAbspath(pkgpath string) []string {
	dir := gopkgDir(pkgpath)
	files := gopkgFiles(pkgpath)

	for i, s := range files {
		files[i] = filepath.Join(dir, s)
	}

	sort.Strings(files)
	return files
}
