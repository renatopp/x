package logx

import "fmt"

// Print is an alias to fmt.Printf.
func Print(msg string, v ...any) {
	fmt.Printf(msg, v...)
}

// Println is an alias to fmt.Printf with a newline at the end.
func Println(msg string, v ...any) {
	fmt.Printf(msg+"\n", v...)
}

// Sprintf is an alias to fmt.Sprintf.
func Sprint(msg string, v ...any) string {
	return fmt.Sprintf(msg, v...)
}

// Sprintln is an alias to fmt.Sprintf with a newline at the end.
func Sprintln(msg string, v ...any) string {
	return fmt.Sprintf(msg+"\n", v...)
}

// Fatal is an alias to panic with a formatted message.
func Fatal(msg string, v ...any) {
	panic(fmt.Sprintf(msg, v...))
}
