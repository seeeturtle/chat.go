package chatgo

import "errors"

func gracefulExit(f func(...interface{}), args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case string:
				err = errors.New(v)
			default:
				err = errors.New("Panic")
			}
		}
	}()
	f(args...)
	return nil
}
