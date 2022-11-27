package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrUnsupportedInput = errors.New("input path isn't a directory")

func (env Environment) GetEnvs() (toSet map[string]string, toDelete []string) {
	toSet = make(map[string]string)
	for key, value := range env {
		if value.NeedRemove {
			toDelete = append(toDelete, key)
		} else {
			toSet[key] = value.Value
		}
	}
	return toSet, toDelete
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrUnsupportedInput
	}
	fileInfos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		fileInfos = append(fileInfos, info)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		if strings.Contains(fileInfo.Name(), "=") {
			continue
		}
		if fileInfo.Size() == 0 {
			env[fileInfo.Name()] = EnvValue{NeedRemove: true}
			continue
		}
		err := func() error {
			file, err := os.OpenFile(filepath.Join(dir, fileInfo.Name()), os.O_RDONLY, os.ModePerm)
			if err != nil {
				return err
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			if scanner.Scan() {
				lineBytes := bytes.Replace(scanner.Bytes(), []byte{0}, []byte{10}, 1)
				env[fileInfo.Name()] = EnvValue{Value: strings.TrimRightFunc(string(lineBytes), isSpace)}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}
	return env, nil
}

// unicode package fails golangci-lint for some reason
func isSpace(r rune) bool {
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

func (e Environment) Strings() []string {
	kvs := make([]string, 0, len(e))
	for key, value := range e {
		kvs = append(kvs, fmt.Sprintf("%v=%v", key, value))
	}
	return kvs
}
