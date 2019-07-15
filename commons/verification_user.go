package commons

import "regexp"

func VerificationEmail(email string) (v bool) {
	pattern := "^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"
	m := regexp.MustCompile(pattern)
	res := m.MatchString(email)
	return res
}

func VerificationPassword(pwd string) (v bool, err string) {
	
}

func VerificationMobile(mobile string) (v bool, err string) {

}