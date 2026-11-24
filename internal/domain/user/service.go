package user

import (
	"regexp"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/domainerr"
)

var (
	usernameRe = regexp.MustCompile(`^[a-zA-Z0-9_]{3,64}$`)
	emailRe    = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	mobileRe    = regexp.MustCompile(`^09[0-9]{9}$`)
)

func validateUserName(name string) error{
	if !usernameRe.MatchString(name){
		return domainerr.New(domainerr.InvalidUsername , "invalid username")
	}
	return nil
}

func validateEmail(email string) error{
	if !emailRe.MatchString(email){
		return domainerr.New(domainerr.InvalidEmail , "invalid email")
	}
	return nil
}

func validateMobile(mobile string) error{
	if !mobileRe.MatchString(mobile){
		return domainerr.New(domainerr.InvalidMobile , "invalid mobile")
	}
	return nil
}