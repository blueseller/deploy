package main

type InputParser interface {
	Read() ([]byte, error)
}

const maxReadBytes = 1024

type ShellInput struct {
	fd          int
}

func (*s ShellInput) Read() ([]byte,error) {
	buf := make([]byte, maxReadBytes)
	n, err := syscall.Read(t.fd, buf)
	if err != nil {
		    return []byte{}, err
	}
	return buf[:n], nil
}

// NewStandardInputParser returns ConsoleParser object to read from stdin.
func NewStandardInputParser() *ShellInput{
	in, err := syscall.Open("/dev/tty", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	return &ShellInput{
		fd: in,
	}
}
