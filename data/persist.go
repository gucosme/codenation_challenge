package data

import "os"

// UpdateFile overwrites the given file with new data
func UpdateFile(file *os.File, data []byte) error {
	_, err := file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
