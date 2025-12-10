package main

import (
    "fmt"
    "os"
    "go_lang/file_processor/internal/processor"
    "path/filepath"
)

func main() {

    // 1. Create worker pool with N workers
    p := processor.NewPool(4)
    p.Start()

    // 2. Scan directory for files
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

    // 3. Read & print results
    for res := range p.Results() {
        if res.Err != nil {
            fmt.Println("ERROR:", res.Path, "→", res.Err)
            continue
        }
        fmt.Printf("[OK] %s → %d bytes\n", res.Path, res.Size)
    }

    fmt.Println("Finished processing all files.")
}
