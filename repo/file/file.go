package file

import "os"

type File struct{}

func New() *File {
	return &File{}
}

func (file *File) WriteToFile(path, text string) error {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		return err
	}

	return nil
}
