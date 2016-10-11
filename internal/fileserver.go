package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

var listTemplate *template.Template

func init() {
	listTemplate = template.Must(template.ParseFiles("templates/list.html"))
}

type FileServer struct {
	root string
}

func NewFileServer(root string) *FileServer {
	return &FileServer{root: filepath.Clean(root)}
}

func (fs *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filepath := fs.root + path.Clean(r.URL.Path)
	if err := serveFile(filepath, w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveFile(filepath string, w http.ResponseWriter, r *http.Request) error {
	file, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	// redirect to canonical path
	if file.IsDir() {
		url := r.URL.Path
		if !strings.HasSuffix(url, "/") {
			RedirectLocal(path.Base(url)+"/", w, r)
			return nil
		}
	} else {
		url := r.URL.Path
		if strings.HasSuffix(url, "/") {
			RedirectLocal("../"+path.Base(url), w, r)
			return nil
		}
	}

	if file.IsDir() {
		return listFiles(filepath, w, r)
	}
	return serveContent(filepath, w, r)
}

type fileAttribute struct {
	Name     string
	IsDir    bool
	IsHidden bool
}

func listFiles(filepath string, w http.ResponseWriter, r *http.Request) error {
	dir, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	sort.Sort(byName(files))

	return listTemplate.Execute(w, Map(files, func(fi os.FileInfo) *fileAttribute {
		return &fileAttribute{
			Name:     fi.Name(),
			IsDir:    fi.IsDir(),
			IsHidden: isHidden(fi),
		}
	}))
}

func Map(src []os.FileInfo, f func(os.FileInfo) *fileAttribute) []*fileAttribute {
	dst := make([]*fileAttribute, len(src))
	for i, v := range src {
		dst[i] = f(v)
	}
	return dst
}

func isHidden(fi os.FileInfo) bool {
	return fi.Name()[0] == '.'
}

func serveContent(filepath string, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "FILE: %s\n", filepath)
	return nil
}

type byName []os.FileInfo

func (s byName) Len() int           { return len(s) }
func (s byName) Less(i, j int) bool { return s[i].Name() < s[j].Name() }
func (s byName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
