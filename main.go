package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const srcDir = "_content"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	return copyDir(srcDir, os.Args[1])
}

func copyDir(from string, to string) error {
	if err := os.RemoveAll(to); err != nil {
		return err
	}

	err := filepath.Walk(from, func(src string, info os.FileInfo, err error) error {
		base := filepath.Base(from)
		rel, err := filepath.Rel(base, src)
		if err != nil {
			return err
		}

		dst := filepath.Join(to, rel)

		if info.IsDir() {
			if err := os.Mkdir(dst, info.Mode()); err != nil {
				return err
			}

			return nil
		}

		in, err := os.Open(src)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, info.Mode())
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err = io.Copy(out, in); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
