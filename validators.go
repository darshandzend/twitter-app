package main

import (
	"fmt"
)

func validateDraft(d draft) error {

	if len(d.Text) > 140 {
		err := fmt.Errorf("a tweet cannot be than 140 chars")
		return err
	}

	return nil
}
