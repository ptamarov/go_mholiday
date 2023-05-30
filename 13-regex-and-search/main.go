package main

import (
	"fmt"
	"log"
	"regexp"
)

var test = []string{
	"072665ae-a034-4cc3-a2e8-9f1822c4ebbb",
	"072665ae-a034-6cc3-a2e8-9f1822c4ebbb",  // Invalid version bit (1-5).
	"072665ae-a034-4cc3-72e8-9f1822c4ebbb",  // Invalid type bit (89abBB).
	"072665ae-a034-4cc3-a2e8-9f1822c4ebb",   // Too short.
	"072665ae-a034-4cc3-a2e8-9f1822c4ebbcb", // Too long.
	"072665ae-a034-3cc3-82e8-9f1822c4ebbb",
}

func main() {

	for i, t := range test {
		if !isUUID(t) {
			fmt.Println(i, t, "\tfails.")
		}
	}

	phone := "(214) 514-9548"
	new := rewritePhoneNumber(phone)
	fmt.Println(new)

	phrase := "Call us at (214) 514-9548 for more information."
	new = rewritePhoneNumberInsideString(phrase)
	fmt.Println(new)

}

func isUUID(s string) bool {
	var pu = `^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`
	var uu = regexp.MustCompile(pu) // Must compile panics if Regex does not compile. Good for constants, but use Compile for runtime inputs.
	return uu.MatchString(s)
}

func rewritePhoneNumber(s string) string {
	var pu = `^\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})$`
	var uu = regexp.MustCompile(pu)

	match := uu.FindStringSubmatch(s)

	fmt.Printf("%q\n", match) // First entry is the whole text that matched, then come the submatches.

	if len(match) == 4 {
		return fmt.Sprintf("+1 %s-%s-%s", match[1], match[2], match[3])
	} else {
		log.Println("Failed to find a valid phone number.")
		return ""
	}
}

func rewritePhoneNumberInsideString(s string) string {
	var pu = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`
	var uu = regexp.MustCompile(pu)

	return uu.ReplaceAllString(s, "+1 ${1}-${2}-${3}") // Capture groups in replacement.

}
