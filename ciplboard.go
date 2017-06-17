package Clipboard-Replacer
import cb "github.com/atotto/clipboard"
import "regexp"
import "time"

type regExPair struct {
	Exp *regexp.Regexp
	replaceTo string
}
var regExs []regExPair

func wait(s time.Duration) {
	time.Sleep(s * time.Second)
}

func AddRegEx(exp string, replaceTo string) {
	r, _ := regexp.Compile(exp)
	pair := regExPair {
		Exp: r,
		replaceTo: replaceTo,
	}
	regExs = append(regExs, pair)
}

func replace(text string) {
	txt := text
	for _, pair := range regExs {
		replaced := pair.Exp.ReplaceAllString(txt, pair.replaceTo)
		write(replaced)
		txt = replaced
	}
}

func Watch(delay time.Duration) {
	for {
		wait(delay)
		str := read()
		if str != "" {
			replace(str)
		}
	}
}

func read() string {
	str, _ := cb.ReadAll()
	return str
}

func write(text string) {
	cb.WriteAll(text)
}
