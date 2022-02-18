package main

import (
	"fmt"
)

var (
	ErrRestreamers = func(msg string) error {
		return fmt.Errorf("error during restreamers operation, %s", msg)
	}
	ErrCustomers = func(msg string) error {
		return fmt.Errorf("error during customers operation, %s", msg)
	}
	ErrVideoUpload = func(msg string) error {
		return fmt.Errorf("error during videos operation, %s", msg)
	}
	ErrEvents = func(msg string) error {
		return fmt.Errorf("error during event operation, %s", msg)
	}
)
