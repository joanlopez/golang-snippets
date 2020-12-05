package read

import (
	"bufio"
	"io"
	"os"
)

func Lines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Technically incorrect

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func Words(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Technically incorrect

	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func Runes(filename string) ([]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Technically incorrect

	var runes []rune
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err == nil {
			runes = append(runes, r)
		}

		if err == io.EOF {
			return runes, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func Bytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Technically incorrect

	var bytes []byte
	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err == nil {
			bytes = append(bytes, b)
		}

		if err == io.EOF {
			return bytes, nil
		}

		if err != nil {
			return nil, err
		}
	}
}
