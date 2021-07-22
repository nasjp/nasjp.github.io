package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

const srcDir = "_content"

const postDir = "_posts"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	if err := copyDir(srcDir, os.Args[1]); err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return nil
	}

	if err := genHTML(postDir, os.Args[1]); err != nil {
		return err
	}

	return nil
}

func copyDir(from string, to string) error {
	if err := os.RemoveAll(to); err != nil {
		return err
	}

	err := filepath.Walk(from, func(src string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

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

func genHTML(from string, to string) error {
	err := filepath.Walk(from, func(src string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		base := filepath.Base(from)
		rel, err := filepath.Rel(base, src)
		if err != nil {
			return err
		}

		dst := filepath.Join(to, strings.TrimSuffix(rel, "md")+"html")

		in, err := os.Open(src)
		if err != nil {
			return err
		}
		defer in.Close()

		source, err := io.ReadAll(in)
		if err != nil {
			return err
		}

		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer out.Close()

		if err := goldmark.Convert(source, out); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
