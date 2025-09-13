package jsonparser

import (
	"errors"
	"strconv"
)

func trimWhiteSpace(json string, currentPosition int) int {
	count := 0

	for i := currentPosition; i < len(json); i++ {
		if json[i] != ' ' && json[i] != '\n' && json[i] != '\t' {
			break
		}
		count++
	}

	return count
}

/*
 * TODO: Use functional or object oriented approach
 */

func Parse(stringToParse string) (interface{}, error) {
	parsedJson, err, _ := internalParse(stringToParse, 0)
	return parsedJson, err
}

func internalParse(json string, currentPosition int) (interface{}, error, int) {
	newPosition := trimWhiteSpace(json, currentPosition) + currentPosition

	if newPosition >= len(json) {
		return nil, errors.New("expecting token but reached end"), len(json) - currentPosition
	}

	switch json[newPosition] {
	case '[':
		return parseArray(json, newPosition)
	case '{':
		return parseObject(json, newPosition)
	default:
		return parseValue(json, newPosition)
	}
}

func parseArray(json string, currentPosition int) ([]interface{}, error, int) {
	arr := make([]interface{}, 0)
	var needValueEnd bool

	i := currentPosition + 1
	for i < len(json) {
		i += trimWhiteSpace(json, i)

		if json[i] == ']' {
			return arr, nil, i - currentPosition + 1
		}

		if needValueEnd {
			if json[i] == ',' {
				i++
				needValueEnd = false
				continue
			} else {
				return arr, errors.New("unexpected token"), i - currentPosition + 1
			}
		}

		val, err, count := internalParse(json, i)
		i += count - 1
		if err != nil {
			return arr, err, i - currentPosition
		}
		arr = append(arr, val)
		needValueEnd = true

		i++
	}

	return arr, errors.New("unexpected token"), i - currentPosition
}

func parseObject(json string, currentPosition int) (map[string]interface{}, error, int) {
	obj := make(map[string]interface{})
	var needValueEnd bool

	i := currentPosition + 1
	for i < len(json) {
		i += trimWhiteSpace(json, i)

		if json[i] == '}' {
			return obj, nil, i - currentPosition + 1
		}

		if needValueEnd {
			if json[i] == ',' {
				i++
				needValueEnd = false
				continue
			} else {
				return obj, errors.New("Unexpected token"), i - currentPosition + 1
			}
		}

		key, val, err, count := parseObjectProperty(json, i)
		i += count - 1
		if err != nil {
			return obj, err, i - currentPosition + 1
		}
		obj[key] = val
		needValueEnd = true

		i++
	}

	return obj, errors.New("Unexpected token"), i - currentPosition + 1
}

func parseObjectProperty(json string, currentPosition int) (string, interface{}, error, int) {
	i := currentPosition

	key, err, count := parseString(json, i)
	i += count
	if err != nil {
		return "", nil, err, i - currentPosition + 1
	}

	i += trimWhiteSpace(json, i)
	if json[i] == ':' {
		i++
	} else {
		return "", nil, errors.New("Unexpected token"), i - currentPosition + 1
	}
	i += trimWhiteSpace(json, i)

	val, err, count := internalParse(json, i)
	i += count - 1
	if err != nil {
		return "", nil, err, i - currentPosition + 1
	}

	return key, val, nil, i - currentPosition + 1
}

func parseValue(json string, currentPosition int) (interface{}, error, int) {
	newPosition := trimWhiteSpace(json, currentPosition) + currentPosition

	if newPosition >= len(json) {
		return nil, errors.New("Expecting token but reached end"), len(json) - currentPosition
	}

	switch json[newPosition] {
	case 'n':
		return nil, parseNull(json, newPosition), 4
	case 'f', 't':
		return parseBool(json, newPosition)
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		return parseNumber(json, newPosition)
	case '"':
		return parseString(json, newPosition)
	default:
		panic("Unexpected token")
	}
}

func parseNull(json string, currentPosition int) error {
	if currentPosition+4 > len(json) {
		return errors.New("Expected null")
	}

	if json[currentPosition:currentPosition+4] != "null" {
		return errors.New("Unexpected token, expected null")
	}

	return nil
}

func parseBool(json string, currentPosition int) (bool, error, int) {
	switch json[currentPosition] {
	case 't':
		if currentPosition+4 > len(json) {
			return false, errors.New("Expected true"), len(json) - currentPosition
		}

		if json[currentPosition:currentPosition+4] != "true" {
			return false, errors.New("Unexpected token, expected true"), len(json) - currentPosition
		}

		return true, nil, 4
	case 'f':
		if currentPosition+5 > len(json) {
			return false, errors.New("Expected false"), len(json) - currentPosition
		}

		if json[currentPosition:currentPosition+5] != "false" {
			return false, errors.New("Unexpected token, expected false"), len(json) - currentPosition
		}

		return false, nil, 5
	default:
		return false, errors.New("Unextpected token, expected boolean"), len(json) - currentPosition
	}
}

func parseNumber(json string, currentPosition int) (int, error, int) {
	var numberString string

	i := currentPosition
	for i < len(json) {
		if json[i] < '0' || json[i] > '9' {
			break
		}
		numberString += string(json[i])
		i++
	}

	number, err := strconv.ParseInt(numberString, 10, 64)
	return int(number), err, i - currentPosition
}

func parseString(json string, currentPosition int) (string, error, int) {
	var str string

	i := currentPosition + 1
	for i < len(json) {
		if json[i] == '"' {
			return str, nil, i - currentPosition + 1
		} else if json[i] == '\\' {
			i++
			str += string(json[i])
		} else {
			str += string(json[i])
		}
		i++
	}

	return "", errors.New("Expected end of string(\")"), i - currentPosition + 1
}
