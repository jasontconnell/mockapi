package process

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	BasePath    string
	ContentType string
	Mappings    map[string]string
}

func NewServer(basePath string, mappings map[string]string, contentType string) *Server {
	s := new(Server)
	s.Mappings = mappings
	s.ContentType = contentType

	if filepath.IsAbs(basePath) {
		s.BasePath = basePath
	} else {
		d, _ := os.Getwd()
		s.BasePath = filepath.Join(d, basePath)
	}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	p := req.URL.Path

	filename, ok := s.Mappings[p]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	b, err := os.ReadFile(filepath.Join(s.BasePath, filename))
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	log.Println("request", req.URL.Path, "response", filename)
	w.Header().Add("Content-Type", s.ContentType)
	w.Write(b)
}
