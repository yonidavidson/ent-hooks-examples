package ent

import "strings"

// FullPhoneNumber returns the phone number unmasked.
func (u *User) FullPhoneNumber() string {
	return strings.ReplaceAll(u.PhoneNumber, "****", u.LastDigits)
}
