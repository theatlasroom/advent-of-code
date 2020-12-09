package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/2020/utils"
)

/**
--- Day 4: Passport Processing ---

You arrive at the airport only to realize that you grabbed your North Pole Credentials instead of your passport. While these documents are extremely similar,
North Pole Credentials aren't issued by a country and therefore aren't actually valid documentation for travel in most of the world.

It seems like you're not the only one having problems, though; a very long line has formed for the automatic passport scanners,
and the delay could upset your travel itinerary.

Due to some questionable network security, you realize you might be able to solve both of these problems at the same time.

The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields.
The expected fields are as follows:

    byr (Birth Year)
    iyr (Issue Year)
    eyr (Expiration Year)
    hgt (Height)
    hcl (Hair Color)
    ecl (Eye Color)
    pid (Passport ID)
    cid (Country ID)

Passport data is validated in batch files (your puzzle input).
Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. Passports are separated by blank lines.

Here is an example batch file containing four passports:

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

The first passport is valid - all eight fields are present. The second passport is invalid - it is missing hgt (the Height field).

The third passport is interesting; the only missing field is cid, so it looks like data from North Pole Credentials, not a passport at all!
Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.

The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.

According to the above rules, your improved system would report 2 valid passports.

Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?

*/

type passport struct {
	// we will default ints to -1 for easy validation
	// we will default strings to "" for easy validation
	PassportID                           string
	CountryID                            int
	BirthYear, IssueYear, ExpirationYear int
	Height, EyeColour, HairColour        string
}

var (
	pid = regexp.MustCompile(`pid:(\S*)`)
	cid = regexp.MustCompile(`cid:(\S*)`)
	byr = regexp.MustCompile(`byr:(\S*)`)
	iyr = regexp.MustCompile(`iyr:(\S*)`)
	eyr = regexp.MustCompile(`eyr:(\S*)`)
	hgt = regexp.MustCompile(`hgt:(\S*)`)
	ecl = regexp.MustCompile(`ecl:(\S*)`)
	hcl = regexp.MustCompile(`hcl:(\S*)`)
)

// validation rules for each field
var (
	cidValid = regexp.MustCompile(`cid:(\S*)`)                // ignore
	pidValid = regexp.MustCompile(`pid:(\d{9})\b`)            // a nine-digit number, including leading zeroes.
	byrValid = regexp.MustCompile(`byr:(19[2-9]\d|200[0-2])`) // four digits; at least 1920 and at most 2002.
	iyrValid = regexp.MustCompile(`iyr:(201\d|2020)`)         // four digits; at least 2010 and at most 2020.
	eyrValid = regexp.MustCompile(`eyr:(202\d|2030)`)         // four digits; at least 2020 and at most 2030.
	// Height:
	// - a number followed by either cm or in:
	// - If cm, the number must be at least 150 and at most 193.
	// - If in, the number must be at least 59 and at most 76.
	hgtValid = regexp.MustCompile(`hgt:((59|6\d|7[0-6])in|(1[5-8]\d|19[0-3])cm)`)
	eclValid = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`) // exactly one of: amb blu brn gry grn hzl oth.
	hclValid = regexp.MustCompile(`hcl:(#[\da-f]{6})`)                 // a # followed by exactly six characters 0-9 or a-f
)

func (p passport) toString() string {
	return fmt.Sprintf(
		"pid: %v, cid: %v, byr: %v, iyr: %v, eyr: %v, hgt: %v, ecl: %v, hcl: %v",
		p.PassportID, p.CountryID, p.BirthYear, p.IssueYear, p.ExpirationYear, p.Height, p.EyeColour, p.HairColour,
	)
}

func (p passport) hasAllRequiredFields() (bool, string) {
	e := reflect.ValueOf(&p).Elem()
	// iterate through all the fields and validate them
	for i := 0; i < e.NumField(); i++ {
		v := e.Field(i).Interface()
		field := e.Type().Field(i).Name
		if field == "CountryID" {
			continue
		}

		switch v.(type) {
		case int:
			vi, ok := v.(int)
			if !ok || vi < 0 {
				return false, ""
				// return false, fmt.Sprintf("invalid field %v value %v", field, vi)
			}
		case string:
			vs, ok := v.(string)
			if !ok || len(vs) < 1 {
				if field == "EyeColour" {
					return false, fmt.Sprintf("invalid field %v value %v", field, vs)
				}
				return false, ""
			}
		default:
			log.Fatal("Unknown type for field %v", v)
		}
	}
	return true, ""
}

type passportRule struct {
	input string
	re    *regexp.Regexp
}

func (v passportRule) extractData() string {
	res := v.re.FindStringSubmatch(v.input)
	if len(res) > 1 {
		return res[1]
	}
	if len(res) > 2 {
		fmt.Printf("input %v matches %v\n", v.input, res)
	}
	return ""
}

func extractIntOrDefault(v passportRule) int {
	str, err := strconv.Atoi(v.extractData())
	if err == nil {
		return str
	}
	return -1
}

func extractStringOrDefault(v passportRule) string {
	str := v.extractData()
	return str
}

func newPassport(input string) passport {
	pidValue := extractStringOrDefault(passportRule{input, pid})
	cidValue := extractIntOrDefault(passportRule{input, cid})
	byrValue := extractIntOrDefault(passportRule{input, byr})
	iyrValue := extractIntOrDefault(passportRule{input, iyr})
	eyrValue := extractIntOrDefault(passportRule{input, eyr})
	hgtValue := extractStringOrDefault(passportRule{input, hgt})
	eclValue := extractStringOrDefault(passportRule{input, ecl})
	hclValue := extractStringOrDefault(passportRule{input, hcl})

	return passport{
		PassportID:     pidValue,
		CountryID:      cidValue,
		BirthYear:      byrValue,
		IssueYear:      iyrValue,
		ExpirationYear: eyrValue,
		Height:         hgtValue,
		EyeColour:      eclValue,
		HairColour:     hclValue,
	}
}

func newValidPassport(input string) (passport, string) {
	pidValue := extractStringOrDefault(passportRule{input, pidValid})
	cidValue := extractIntOrDefault(passportRule{input, cidValid})
	byrValue := extractIntOrDefault(passportRule{input, byrValid})
	iyrValue := extractIntOrDefault(passportRule{input, iyrValid})
	eyrValue := extractIntOrDefault(passportRule{input, eyrValid})
	hgtValue := extractStringOrDefault(passportRule{input, hgtValid})
	eclValue := extractStringOrDefault(passportRule{input, eclValid})
	hclValue := extractStringOrDefault(passportRule{input, hclValid})

	debugInfo := fmt.Sprintln("pidValue", pidValue, "cidValue", cidValue, "byrValue", byrValue, "eyrValue", eyrValue, "iyrValue", iyrValue, "hgtValue", hgtValue, "eclValue", eclValue, "hclValue", hclValue)

	return passport{
		PassportID:     pidValue,
		CountryID:      cidValue,
		BirthYear:      byrValue,
		IssueYear:      iyrValue,
		ExpirationYear: eyrValue,
		Height:         hgtValue,
		EyeColour:      eclValue,
		HairColour:     hclValue,
	}, debugInfo
}

func clean(data string) []string {
	passportBlocks := strings.Split(data, "\n\n")

	var passports []string
	for _, p := range passportBlocks {
		passports = append(passports, strings.Join(strings.Split(p, "\n"), "\n"))
	}
	return passports
}

func part01(passports []string) int {
	count := 0
	for _, input := range passports {
		p := newPassport(input)
		ok, _ := p.hasAllRequiredFields()
		if ok {
			count++
		}
	}
	return count
}

func part02(passports []string) int {
	count := 0
	for _, input := range passports {
		p, _ := newValidPassport(input)
		ok, _ := p.hasAllRequiredFields()
		if ok {
			count++
		}
	}
	return count
}

// // TODO: redo with concurrency
func main() {
	utils.Banner(4)
	passports := clean(utils.LoadDataAsString("4.txt"))
	fmt.Printf("%v passports => %v candidates\n", len(passports), part01(passports))
	fmt.Printf("%v passports => %v valid\n", len(passports), part02(passports))
}
