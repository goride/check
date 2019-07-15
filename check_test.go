package check

import (
	"fmt"
	"testing"
)


func TestBasic(t *testing.T) {
	//Check(func() error {
	//	return errors.New("abc")
	//}())
	Check(func() error {
		return nil
	}())
	value := Check(func() (v interface{}, err error) {
		return 1 , nil
	}())
	fmt.Print("test get value: ", value, "\r\n")
	//Check(func() (v interface{}, err error) {
	//	return 1 , errors.New("error message")
	//}())
}
