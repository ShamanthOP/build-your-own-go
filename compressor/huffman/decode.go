package huffman

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Decode(filename string) error {

	numCharacters, prefixTable, err := readHeader(filename)
	if err != nil {
		return err
	}

	fmt.Println(numCharacters)
	fmt.Println(prefixTable)

	return nil
}

func readHeader(filename string) (int, map[rune]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, err
	}

	var numCharacters uint32
	if err := binary.Read(file, binary.BigEndian, &numCharacters); err != nil {
		return 0, nil, err
	}

	fmt.Println(numCharacters)

	prefixTable := make(map[rune]string)

	for i := uint32(0); i < numCharacters; i++ {
		var char uint8
		if err := binary.Read(file, binary.BigEndian, &char); err != nil {
			return 0, nil, err
		}

		fmt.Println(char)

		var codeLength uint8
		if err := binary.Read(file, binary.BigEndian, &codeLength); err != nil {
			return 0, nil, err
		}

		fmt.Println(codeLength)

		var code string
		for j := uint8(0); j < codeLength-6; j++ {
			var bit byte
			if err := binary.Read(file, binary.BigEndian, &bit); err != nil {
				fmt.Println(filename)
				return 0, nil, err
			}
			if bit == 1 {
				code += "1"
			} else {
				code += "0"
			}
		}

		prefixTable[rune(char)] = code
	}

	fmt.Println(filename, prefixTable)

	var endMarker uint16
	if err := binary.Read(file, binary.BigEndian, &endMarker); err != nil {
		return 0, nil, err
	}
	if endMarker != 0xFFFF {
		return 0, nil, fmt.Errorf("invalid end-of-header marker %v %v", endMarker, 0xFFFF)
	}

	fmt.Println(filename)

	return int(numCharacters), prefixTable, nil
}
