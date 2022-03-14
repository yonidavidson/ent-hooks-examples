package ent

import "strings"

// FullConnectionString returns the connection string with an unmasked password.
func (u *User) FullConnectionString() string {
	return strings.ReplaceAll(u.ConnectionString, "****", u.Password)
}
