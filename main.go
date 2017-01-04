package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
)

func main() {
	res := map[*regexp.Regexp]func(string, ...interface{}) string{
		regexp.MustCompile("(PASS)"): color.GreenString,
		regexp.MustCompile("(FAIL)"): color.RedString,
		regexp.MustCompile("(RUN)"):  color.BlueString,
		regexp.MustCompile("(SKIP)"): color.YellowString,
		regexp.MustCompile("^(ok)"):  color.CyanString,
		regexp.MustCompile("^(\\?)"): color.MagentaString,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		for re, co := range res {
			t = re.ReplaceAllString(t, co("$1"))
		}
		fmt.Println(t)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
