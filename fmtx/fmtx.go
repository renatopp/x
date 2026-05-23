package fmtx

import (
	"fmt"
	"io"
)

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

func FPrint(w io.Writer, msg string, v ...any) (n int, err error) {
	return fmt.Fprintf(w, msg, v...)
}

func FPrintln(w io.Writer, msg string, v ...any) (n int, err error) {
	return fmt.Fprintf(w, msg+"\n", v...)
}

// Fatal is an alias to panic with a formatted message.
func Fatal(msg string, v ...any) {
	panic(fmt.Sprintf(msg, v...))
}

func Error(msg string, v ...any) error {
	return fmt.Errorf(msg, v...)
}

func Scan(format string, v ...any) (n int, err error) {
	return fmt.Scanf(format, v...)
}

func Scanln(v ...any) (n int, err error) {
	return fmt.Scanln(v...)
}

func Sscan(str, format string, v ...any) (n int, err error) {
	return fmt.Sscanf(str, format, v...)
}

func Sscanln(str string, v ...any) (n int, err error) {
	return fmt.Sscanln(str, v...)
}

func Fscan(r io.Reader, format string, v ...any) (n int, err error) {
	return fmt.Fscanf(r, format, v...)
}

func Fscanln(r io.Reader, v ...any) (n int, err error) {
	return fmt.Fscanln(r, v...)
}
