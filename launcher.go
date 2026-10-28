
package main

import (
    "bufio"
    "fmt"
    "go/ast" // Add this line
    "go/parser"
    "go/token"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "strconv"
    "strings"
)

type LabFile struct {
	Path string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Find the folder containing this launcher.
	root, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot access working folder:", err)
		return
	}

	for {
		labs, err := scanLabs(root)
		if err != nil {
			fmt.Println("Scan error:", err)
			return
		}

		fmt.Println("\n====================================")
		fmt.Println("       GO ASSIGNMENT LAUNCHER")
		fmt.Println("====================================")
		fmt.Println("Folder:", root)
		fmt.Println()

		if len(labs) == 0 {
			fmt.Println("No valid Go lab files found.")
		}

		for i, lab := range labs {
			fmt.Printf("%2d. %s\n", i+1, lab.Path)
		}

		fmt.Println("\nR = Rescan | 0 = Exit")
		fmt.Print("Choose file number: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "r") {
			continue
		}

		if input == "0" {
			fmt.Println("Launcher closed.")
			return
		}

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(labs) {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		runLab(root, labs[choice-1].Path)
	}
}

// Scan all subfolders for valid Go files with a main function.
func scanLabs(root string) ([]LabFile, error) {
	var labs []LabFile

	err := filepath.WalkDir(root, func(
		path string, entry os.DirEntry, walkErr error,
	) error {
		if walkErr != nil {
			fmt.Println("Skipping inaccessible path:", path)
			return nil
		}

		if entry.IsDir() && path != root {
			name := entry.Name()
			if strings.HasPrefix(name, ".") ||
				name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		if entry.IsDir() ||
			!strings.HasSuffix(path, ".go") ||
			strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Parse the file to check that it is valid Go source.
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil || file.Name.Name != "main" {
			return nil
		}

		// Only list standalone programs with func main().
		hasMain := false
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok &&
				fn.Recv == nil && fn.Name.Name == "main" {
				hasMain = true
				break
			}
		}

		if hasMain {
			rel, err := filepath.Rel(root, path)
			if err == nil && rel != "launcher.go" {
				labs = append(labs, LabFile{Path: rel})
			}
		}

		return nil
	})

	sort.Slice(labs, func(i, j int) bool {
		return strings.ToLower(labs[i].Path) <
			strings.ToLower(labs[j].Path)
	})

	return labs, err
}

// Run the selected Go source file.
func runLab(root, file string) {
	fmt.Println("\nRunning:", file)
	fmt.Println("------------------------------------")

	cmd := exec.Command("go", "run", file)
	cmd.Dir = root
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("\nExecution failed:", err)
	}

	fmt.Println("\nPress Enter to return to the menu.")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
