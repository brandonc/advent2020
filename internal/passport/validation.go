package passport

import (
	"regexp"
	"strconv"
	"strings"
)

var hclPattern = *regexp.MustCompile(`^#[0-9a-f]{6}$`)
var pidPattern = *regexp.MustCompile(`^[0-9]{9}$`)
var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

// IsValidByr validates byr field
func IsValidByr(val string) bool {
	year, err := strconv.Atoi(val)
	return err == nil && year >= 1920 && year <= 2002
}

// IsValidIyr validates iyr field
func IsValidIyr(val string) bool {
	year, err := strconv.Atoi(val)
	return err == nil && year >= 2010 && year <= 2020
}

// IsValidEyr validates eyr field
func IsValidEyr(val string) bool {
	year, err := strconv.Atoi(val)
	return err == nil && year >= 2020 && year <= 2030
}

// IsValidHgt validates hgt field
func IsValidHgt(val string) bool {
	if strings.HasSuffix(val, "cm") {
		cm, err := strconv.Atoi(strings.TrimSuffix(val, "cm"))
		return err == nil && cm >= 150 && cm <= 193
	}

	if strings.HasSuffix(val, "in") {
		in, err := strconv.Atoi(strings.TrimSuffix(val, "in"))
		return err == nil && in >= 59 && in <= 76
	}

	return false
}

// IsValidHcl validates hcl field
func IsValidHcl(val string) bool {
	return hclPattern.MatchString(val)
}

// IsValidEcl validates ecl field
func IsValidEcl(val string) bool {
	switch val {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}

	return false
}

// IsValidPid validates pid field
func IsValidPid(val string) bool {
	return pidPattern.MatchString(val)
}

// AllFieldsPresent validates that document fields exist
func AllFieldsPresent(document *map[string]string) bool {
	for _, f := range requiredFields {
		if _, ok := (*document)[f]; !ok {
			return false
		}
	}

	return true
}

// IsValid validates that document fields exist and each field is valid
func IsValid(document *map[string]string) bool {
	return AllFieldsPresent(document) && 
		IsValidByr((*document)["byr"]) &&
		IsValidEcl((*document)["ecl"]) &&
		IsValidEyr((*document)["eyr"]) &&
		IsValidHcl((*document)["hcl"]) &&
		IsValidHgt((*document)["hgt"]) &&
		IsValidIyr((*document)["iyr"]) &&
		IsValidPid((*document)["pid"])
}