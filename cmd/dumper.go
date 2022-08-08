package main

import (
	"bufio"
	"fmt"
	"os"
	"phone-script/cmd/parser"
	"phone-script/domain"
	"phone-script/domain/validator"
	"strings"
)

func main() {
	areaCodes, err := getAreaCodes("area_codes.txt")
	if err != nil {
		panic(err)
	}

	phones, err := getNumbers("phones.txt")
	if err != nil {
		panic(err)
	}

	result := make(map[parser.AreaCodeDump]int)
	for i := range phones {
		var area parser.AreaCodeDump

		for c := range areaCodes {
			if strings.Contains(phones[i].Number, string(areaCodes[c])) {
				area = areaCodes[c]
				break
			}
		}

		result[area]++
	}

	for key, value := range result {
		fmt.Printf("[%s]: %d\n", key, value)
	}
}

func getNumbers(filePath string) ([]domain.Phone, error) {
	var phones []domain.Phone

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var item parser.PhoneDump
		var phone domain.Phone

		item, err = parser.ParseNumbers(scanner.Text())
		if err != nil {
			continue
		}

		phone, err = convertDump(item)
		if err != nil {
			continue
		}

		err = validator.Validate(phone)
		if err != nil {
			continue
		}

		phones = append(phones, phone)
	}

	return phones, nil
}

func getAreaCodes(filePath string) ([]parser.AreaCodeDump, error) {
	var codes []parser.AreaCodeDump

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var code parser.AreaCodeDump

		code, err = parser.ParseAreaCodes(scanner.Text())
		if err != nil {
			return nil, err
		}

		codes = append(codes, code)
	}

	return codes, nil
}

func convertDump(dump parser.PhoneDump) (phone domain.Phone, err error) {
	strNum := strings.ReplaceAll(dump.Number, " ", "")

	phone.OptionalCode = dump.OptionalCode
	phone.Number = strNum

	return
}
