package utils

import (
	"errors"
	"math/rand"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("year can't be negative or equal to 0")
	}
	if year >= 9999 {
		return errors.New("input year is too large")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month > 12 || month < 1 {
		return errors.New("wrong month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day > 31 || day < 1 {
		return errors.New("wrong day")
	}
	if d.month >= 8 && d.month%2 != 0 {
		if day > 30 {
			return errors.New("wrong day.Can't larger than 30")
		}
	}
	if d.month < 8 && d.month%2 == 0 {
		if day > 30 {
			return errors.New("wrong day.Can't larger than 30")
		}
	}
	if !((d.year%4 == 0 && d.year%100 != 0) || d.year%400 == 0) { //not leap year
		if d.month == 2 && day > 28 {
			return errors.New("wrong day.Can't larger than 29")
		}
	}
	d.day = day
	return nil
}

func RandBirthday() Date {
	rand.Seed(time.Now().UnixNano())
	d := Date{}
	for {
		randYear := rand.Intn(120) + 1900
		if d.SetYear(randYear) == nil {
			break
		}
	}
	for {
		randMonth := rand.Intn(12) + 1
		if d.SetMonth(randMonth) == nil {
			break
		}
	}
	for {
		randDay := rand.Intn(31) + 1
		if d.SetDay(randDay) == nil {
			break
		}
	}
	return d
}
