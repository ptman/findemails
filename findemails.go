package findemails

import "regexp"

// Yes, I know about http://emailregex.com/ , but let's focus on practical stuff
// TODO: list of all TLDs?
// TODO: more valid characters?
// TODO: stricter limits on domain names?

// https://www.w3.org/TR/html5/forms.html#valid-e-mail-address
var emailRE = regexp.MustCompile(
	"[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")

// FindEmails returns a slice of email addresses found in the input
func FindEmails(src string) []string {
	return emailRE.FindAllString(src, -1)
}
