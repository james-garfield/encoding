package main

import (
	"archive/zip"
	"os"
	"strings"
)

// "C:\Users\jorda\Desktop\Light Novels\Classroom of the Elite\COTE3.zip"

func main() {
	// Get the path to the EPUB file
	epubPath := os.Args[1]
	// Convert to a ZIP file.
	zipPath := convertToZip(epubPath)

	reader, writer, files := getZipFiles(zipPath)

	// Loop through the files
	for _, file := range files {
		// Encode the file as UTF-8
		encodeHtml(reader, writer, file)
	}

	// Close the zip file
	reader.Close()
	writer.Close() 

	// Convert back to an EPUB
	convertToEpub(zipPath)
}

// Get the files from a zip file without extracting it
func getZipFiles(zipPath string) (*zip.ReadCloser, *zip.Writer, []string) {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		panic(err)
	}

	files := []string{}

	// Loop through the files in the zip file
	for _, file := range zipReader.File {
		extension := strings.Split(file.Name, ".")
		if len(extension) == 1 {
			continue
		}
		if extension[1] != "html" {
			continue
		}

		files = append(files, file.Name)
	}

	writer, err := os.Open(zipPath)
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(writer)

	return zipReader, zipWriter, files
}


// Encode a HTML file as UTF-8
func encodeHtml(zipReader *zip.ReadCloser, zipWriter *zip.Writer, path string) {
	// Get contents of file
	file, err := zipReader.Open(path)
	if err != nil {
		panic(err)
	}

	var contents []byte 
	// Get bytes of file
	number, err := file.Read(contents)
	if err != nil {
		panic(err)
	}
	print(number)
	file.Close()

	// Open a writer
	fileWriter, err := zipWriter.Create(path)
	if err != nil {
		panic(err)
	}

	// Write the bytes to the file
	fileWriter.Write(contents)
}

func convertToZip(epubPath string) string {
	return _convertFileTo(epubPath, "zip", false)
}

func convertToEpub(zipPath string) string {
	return _convertFileTo(zipPath, "epub", true)
}

// Convert the zip file back to a EPUB
func _convertFileTo(path string, new string, delete bool) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	// Get the extension
	extension := strings.Split(path, ".")[1]
	// Create a new path!
	newPath	:= strings.Replace(path, "." + extension, "." + new, 1)

	// Write the data to the new path
	err = os.WriteFile(newPath, data, 0644)

	if err != nil {
		panic(err)
	}

	// Remove the old file
	if delete {
		os.Remove(path)
	}

	return newPath
}