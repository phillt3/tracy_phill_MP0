package message

import (
	"time"
)

type Message struct {
	To      string
	From    string
	Title   string
	Content string
	Date    time.Time
}

func NewMessage(to string, from string, title string, content string, date time.Time) Message {
	var m Message
	m.To = to
	m.From = from
	m.Title = title
	m.Content = content
	m.Date = date
	return m
}
