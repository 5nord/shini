package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Print(v interface{}) (err error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		if s.Len() == 0 {
			return errors.New("empty list")
		}
		for i := 0; i < s.Len(); i++ {
			if err2 := print(s.Index(i).Interface()); err2 != nil {
				err = err2
			}
		}
	default:
		err = print(v)
	}

	return err
}

func print(v interface{}) (err error) {
	if v == nil {
		fmt.Println()
		return errors.New("key not found")
	}

	fmt.Println(v)
	return nil
}
