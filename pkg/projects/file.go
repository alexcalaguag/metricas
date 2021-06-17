package projects

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func ReadingFile() []string {
	repositories := []string{}
	fptr := flag.String("fpath", "../pkg/projects/repositories.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		repositories = append(repositories, s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return repositories
}
