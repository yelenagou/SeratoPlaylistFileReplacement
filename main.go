package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePath string
	fmt.Print("Enter the path to the .m3u file: (no quotes around file name) ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}
	filePath = strings.TrimSpace(input) // Trim extra spaces or newlines

	// Open the .m3u file for reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Replace .flac with .mp3
		line = strings.ReplaceAll(line, ".flac", ".mp3")
		// Replace D:\\HiQualMusicLibrary\\ with D:\\SeratoMusicLibrary\\SERATO_MUSIC\\
		line = strings.ReplaceAll(line, "D:\\HiQualMusicLibrary\\", "D:\\SeratoMusicLibrary\\SERATO_MUSIC\\")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Open the .m3u file for writing
	outputFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Printf("Error flushing to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File updated successfully!")
}
