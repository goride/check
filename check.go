package check

import "log"

func Check(arg ...interface{}) (value interface{}) {
	argLen := len(arg)
	var err interface{}
	switch argLen {
	case 1:
		err = arg[0]
	case 2:
		value = arg[0]
		err = arg[1]
	}
	if err != nil {
		log.Fatal(err)
	}
	return value
}