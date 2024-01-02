package series

import "time"

const dateFomrat = "2006-01-02"

func DateToday() string {
	return time.Now().Format(dateFomrat)
}

func DateYearAgo() string {
	return time.Now().AddDate(-1, 0, 0).Format(dateFomrat)
}
