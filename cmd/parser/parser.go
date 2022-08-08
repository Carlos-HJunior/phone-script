package parser

import (
	"errors"
	"regexp"
)

const numberRegex = `^(?P<init>\+|0{2}|)(?P<number>\d{1,}[\d ]*)$`
const AreaRegex = `^(?P<number>\d*)$`

type PhoneDump struct {
	OptionalCode string
	Number       string
}

type AreaCodeDump string

func ParseNumbers(line string) (dump PhoneDump, err error) {
	r, err := regexp.Compile(numberRegex)
	if err != nil {
		return
	}

	items := r.FindStringSubmatch(line)
	if len(items) < 2 {
		err = errors.New("line doesn't match")
		return
	}

	dump = PhoneDump{
		OptionalCode: items[1],
		Number:       items[2],
	}

	return
}

func ParseAreaCodes(line string) (code AreaCodeDump, err error) {
	r, err := regexp.Compile(AreaRegex)
	if err != nil {
		return
	}

	items := r.FindStringSubmatch(line)
	if len(items) < 2 {
		err = errors.New("line doesn't match")
		return
	}

	code = AreaCodeDump(items[1])
	return
}
