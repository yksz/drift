package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
)

type FileAttribute struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsDir    bool   `json:"is_dir"`
	IsHidden bool   `json:"is_hidden"`
}

func FileHandler(root string) http.Handler {
	return &fileHandler{root: filepath.Clean(root)}
}

type fileHandler struct {
	root string
}

func (h *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filepath := path.Clean(h.root + query.Get("path"))
	if err := h.serveFile(filepath, w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *fileHandler) serveFile(filepath string, w http.ResponseWriter, r *http.Request) error {
	file, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	if file.IsDir() {
		return h.listFiles(filepath, w, r)
	}
	return h.serveContent(filepath, w, r)
}

func (h *fileHandler) listFiles(dirpath string, w http.ResponseWriter, r *http.Request) error {
	dir, err := os.Open(dirpath)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	sort.Sort(byName(files))

	rel, err := filepath.Rel(h.root, dirpath)
	if err != nil {
		return err
	}

	attrs := Map(files, func(file os.FileInfo) *FileAttribute {
		return &FileAttribute{
			Name:     file.Name(),
			Path:     filepath.Join(rel, file.Name()),
			IsDir:    file.IsDir(),
			IsHidden: isHidden(file),
		}
	})
	return json.NewEncoder(w).Encode(attrs)
}

func Map(src []os.FileInfo, fn func(os.FileInfo) *FileAttribute) []*FileAttribute {
	dst := make([]*FileAttribute, len(src))
	for i, v := range src {
		dst[i] = fn(v)
	}
	return dst
}

func isHidden(fi os.FileInfo) bool {
	return fi.Name()[0] == '.'
}

func (h *fileHandler) serveContent(filepath string, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "FILE: %s\n", filepath)
	return nil
}

type byName []os.FileInfo

func (s byName) Len() int           { return len(s) }
func (s byName) Less(i, j int) bool { return s[i].Name() < s[j].Name() }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
