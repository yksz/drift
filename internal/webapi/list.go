package webapi

import (
	"encoding/json"
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

func ListAPI(root string) http.Handler {
	return &listAPI{root: filepath.Clean(root)}
}

type listAPI struct {
	root string
}

func (a *listAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := path.Clean(a.root + query.Get("path"))
	if err := a.serve(filename, w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *listAPI) serve(filename string, w http.ResponseWriter, r *http.Request) error {
	file, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return nil
	}
	dirname := filename

	dir, err := os.Open(dirname)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	sort.Sort(byName(files))

	rel, err := filepath.Rel(a.root, dirname)
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

type byName []os.FileInfo

func (n byName) Len() int           { return len(n) }
func (n byName) Less(i, j int) bool { return n[i].Name() < n[j].Name() }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
