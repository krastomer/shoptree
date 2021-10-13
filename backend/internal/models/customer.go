package models

import (
	"regexp"
	"strings"
	"time"
)

type CustomerProfile struct {
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type Customer struct {
	CustomerProfile
	ID        int
	CreatedAt time.Time
}

func (c *CustomerProfile) TrimSpace() {
	compile := regexp.MustCompile(`\s+`)
	c.Name = compile.ReplaceAllString(c.Name, " ")
	c.Name = strings.TrimLeft(c.Name, " ")
	c.Name = strings.TrimRight(c.Name, " ")
}
