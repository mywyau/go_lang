package errorhandling

import "os"

func read() error {
	f, err := os.Open("file.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
