package helpers

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func ReadInput(filename string) ([]string, error) {
    fmt.Println("ReadInput")

    file := path.Join("./inputs", filename)
    fmt.Printf("Reading file from -> %s\n", file)

    raw, err := os.ReadFile(file)

    if err != nil {
        return nil, err
    }

    contents := strings.Split(strings.TrimSpace(string(raw)), "\n")
    return contents, nil
}
