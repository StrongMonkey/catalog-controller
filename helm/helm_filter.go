package helm

import (
	"encoding/base64"
	"fmt"
	"strings"

	catalogv1 "github.com/rancher/types/apis/catalog.cattle.io/v1"
)

var (
	allowedFileTypes = map[string]bool{
		"yaml":       true,
		"tpl":        true,
		"md":         true,
		"txt":        true,
		"yml":        true,
		"helmignore": true,
	}
)

func filterFile(f catalogv1.File) catalogv1.File {
	extPos := strings.LastIndex(f.Name, ".")
	if extPos == -1 {
		// file type undetermined, so base64 encode it
		return encodedFile(f)
	}
	ext := f.Name[extPos+1:]
	if _, ok := allowedFileTypes[strings.ToLower(ext)]; ok {
		return f
	}
	return encodedFile(f)
}

func encodedFile(f catalogv1.File) catalogv1.File {
	return catalogv1.File{
		Name:     fmt.Sprintf("%s.base64", f.Name),
		Contents: base64.StdEncoding.EncodeToString([]byte(f.Contents)),
	}
}
