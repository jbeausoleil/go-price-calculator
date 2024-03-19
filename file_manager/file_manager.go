package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

// FileManager represents a type that provides file management functionalities. It has two properties, InputFilePath and OutputFilePath,
// which store the paths of the input and output files respectively.
// ReadLines is a method defined on the FileManager type that reads the content of the input file line by line and returns it as a slice
// of strings. It returns an error if an error occurs during file handling or reading.
// WriteResult is a method defined on the FileManager type that takes data as an argument and writes it to the output file in JSON format.
// It returns an error if an error occurs while creating the file or converting the data to JSON.
// New is a function that creates and returns a new instance of FileManager with the provided input and output file paths.
type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// ReadLines reads the contents of the file specified by the InputFilePath field of the FileManager struct.
// It returns a slice of strings where each string represents a line in the file.
// If there is an error opening or reading the file, it returns an error.
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file) // os.Open provides *os.File which implements the Reader interface

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		//file.Close()
		return nil, errors.New("failed to read file content")
	}

	//file.Close()
	return lines, nil
}

// WriteResult writes the given data to the output file specified in FileManager.
// It returns an error if there is any issue creating the file or encoding the data to JSON.
func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

// New creates a new instance of FileManager with the given input and output file paths.
//
// Usage:
//
//	fm := New("input.txt", "output.txt")
//
// Parameters:
//
//	inputPath  - string representing the file path to read from
//	outputPath - string representing the file path to write to
//
// Returns:
//
//	FileManager - a new instance of FileManager
func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
