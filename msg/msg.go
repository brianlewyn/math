package msg

import "fmt"

// ! Error
type myError struct {
	happen string
}

func Error(msg string) *myError {
	return &myError{happen: msg}
}

func Errorf(msg string, a ...any) *myError {
	msg = fmt.Sprintf(msg, a...)
	return &myError{happen: msg}
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.happen)
}

// ! My Sprint
func Sprint(a ...any) string {
	return fmt.Sprintf("%v", a)
}
func Sprintf(msg string, a ...any) string {
	return fmt.Sprintf(msg, a...)
}
