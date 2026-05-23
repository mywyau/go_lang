package main

import (
	"fmt"
	"go_lang/file_processor/internal/processor"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// --- Router
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// --- Background worker pool
	go startFileProcessor()

	// --- Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func startFileProcessor() {
	p := processor.NewPool(4)
	p.Start()

	go func() {
		filepath.Walk("./file_processor/files", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if !info.IsDir() {
				p.Submit(processor.Job{Path: path})
			}
			return nil
		})
		p.Stop()
	}()

	for res := range p.Results() {
		if res.Err != nil {
			fmt.Println("ERROR:", res.Path, "→", res.Err)
			continue
		}
		fmt.Printf("[OK] %s → %d bytes\n", res.Path, res.Size)
	}

	fmt.Println("Finished processing all files.")
}
