package huffman

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Encode(filename string, outputFilename string) error {

	frequency, err := createFrequencyMap(filename)
	if err != nil {
		return err
	}

	prefixTable := buildPrefixTable(frequency)

	compressedData, err := compressData(filename, prefixTable)
	if err != nil {
		return err
	}

	if err := outputToFile(outputFilename, prefixTable, compressedData); err != nil {
		return err
	}

	return nil
}

func createFrequencyMap(filename string) (map[rune]int, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file")
	}

	frequency := make(map[rune]int)

	for _, character := range fileContent {
		frequency[rune(character)]++
	}

	return frequency, nil
}

func buildPrefixTable(frequency map[rune]int) map[rune]string {

	huffmanTree := buildTree(frequency)

	prefixTable := constructTable(huffmanTree)

	return prefixTable
}

func compressData(filename string, prefixTable map[rune]string) ([]byte, error) {

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file")
	}

	var compressedData []byte
	var currentByte byte

	bitIndex := 0

	for _, character := range fileContent {
		code, exists := prefixTable[rune(character)]
		if !exists {
			return nil, fmt.Errorf("huffman code not found for character %c", rune(character))
		}

		for _, codeBit := range code {

			if codeBit == '1' {
				currentByte |= (1 << (7 - bitIndex))
			}
			bitIndex++

			if bitIndex == 8 {
				compressedData = append(compressedData, currentByte)
				currentByte = 0
				bitIndex = 0
			}

		}

	}

	if bitIndex > 0 {
		compressedData = append(compressedData, currentByte)
	}

	return compressedData, nil
}

func outputToFile(outputFilename string, prefixTable map[rune]string, compressedData []byte) error {

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}

	if err := writeHeader(outputFile, prefixTable); err != nil {
		return err
	}

	if err := binary.Write(outputFile, binary.BigEndian, compressedData); err != nil {
		return err
	}

	return nil
}

func writeHeader(outputFile *os.File, prefixTable map[rune]string) error {

	numCharacters := uint32(len(prefixTable))
	if err := binary.Write(outputFile, binary.BigEndian, numCharacters); err != nil {
		return err
	}

	for character, code := range prefixTable {

		if err := binary.Write(outputFile, binary.BigEndian, uint8(character)); err != nil {
			return err
		}

		codeLength := uint8(len(code))
		if err := binary.Write(outputFile, binary.BigEndian, codeLength); err != nil {
			return err
		}

		var currentByte byte

		bitIndex := 0

		for _, codeBit := range code {

			if codeBit == '1' {
				currentByte |= (1 << (7 - bitIndex))
			}
			bitIndex++

			if bitIndex == 8 {
				if err := binary.Write(outputFile, binary.BigEndian, currentByte); err != nil {
					return err
				}
				currentByte = 0
				bitIndex = 0
			}

		}

		if bitIndex > 0 {
			if err := binary.Write(outputFile, binary.BigEndian, currentByte); err != nil {
				return err
			}
		}

	}

	endMarker := uint16(0xFFFF)
	if err := binary.Write(outputFile, binary.BigEndian, endMarker); err != nil {
		return err
	}

	return nil
}
