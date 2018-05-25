package main // import "github.com/dim13/ct"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
)

func main() {
	res := map[*regexp.Regexp]string{
		regexp.MustCompile(`(PASS)`): color.GreenString(`$1`),
		regexp.MustCompile(`(FAIL)`): color.RedString(`$1`),
		regexp.MustCompile(`(RUN)`):  color.BlueString(`$1`),
		regexp.MustCompile(`(SKIP)`): color.YellowString(`$1`),
		regexp.MustCompile(`^(\?)`):  color.MagentaString(`$1`),
		regexp.MustCompile(`^(ok)`):  color.CyanString(`$1`),
		regexp.MustCompile(`(\S+\.go):(\d+)`): fmt.Sprintf("%v:%v",
			color.CyanString(`$1`), color.RedString(`$2`)),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		for re, s := range res {
			t = re.ReplaceAllString(t, s)
		}
		fmt.Println(t)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
