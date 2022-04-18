package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"time"
)

func main() {
	dir := "."
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		date, err := extractDate(f.Name())
		if err == nil {
			fmt.Printf("mkdir -p %s\n", renderDirPath(dir, date))
			fmt.Printf("touch -m -t %s \"%s/%s\"\n", date.Format("200601021504"), dir, f.Name())
			fmt.Printf("mv \"%s/%s\" \"%s/%s\"\n", dir, f.Name(), renderDirPath(dir, date), f.Name())
		}
	}
}

func extractDate(filename string) (time.Time, error) {
	date, err := tryExtractYYYYMMDD(filename)
	if err == nil {
		return date, nil
	} else {
		date, err := tryExtractYYMMDD(filename)
		if err == nil {
			return date, nil
		} else {
			date, err := tryExtractYYYYMM(filename)
			if err == nil {
				return date, nil
			} else {
				date, err := tryExtractYY_MM(filename)
				if err == nil {
					return date, nil
				} else {
					return time.Time{}, fmt.Errorf("no date found")
				}
			}
		}
	}
}

func tryExtractYYYYMMDD(filename string) (time.Time, error) {
	re2 := regexp.MustCompile(`(202\d)(\d\d)(\d\d)`)
	matches := re2.FindStringSubmatch(filename)

	if len(matches) == 4 {
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day, _ := strconv.Atoi(matches[3])
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, fmt.Errorf("no date found")
}

func tryExtractYYMMDD(filename string) (time.Time, error) {
	re2 := regexp.MustCompile(`^(\d\d)(\d\d)(\d\d)\D`)
	matches := re2.FindStringSubmatch(filename)

	if len(matches) == 4 {
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day, _ := strconv.Atoi(matches[3])
		return time.Date(year+2000, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, fmt.Errorf("no date found")
}

func tryExtractYYYYMM(filename string) (time.Time, error) {
	re2 := regexp.MustCompile(`\D(202\d)(\d\d)\D`)
	matches := re2.FindStringSubmatch(filename)

	if len(matches) == 3 {
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day := 1 // fallback: first day of month
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, fmt.Errorf("no date found")
}

func tryExtractYY_MM(filename string) (time.Time, error) {
	re2 := regexp.MustCompile(`\D(\d\d)\.(\d\d)\D`)
	matches := re2.FindStringSubmatch(filename)

	if len(matches) == 3 {
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day := 1 // fallback: first day of month
		return time.Date(year+2000, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, fmt.Errorf("no date found")
}

func renderDirPath(parentDir string, date time.Time) string {
	return fmt.Sprintf("%s/%d/%02d", parentDir, date.Year(), date.Month())
}
