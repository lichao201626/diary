package redis

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type RedisError string

func (err RedisError) Error() string { return "Redis Error: " + string(err) }

var doesNotExist = RedisError("Key does not exist ")

// reads a bulk reply (i.e $5\r\nhello)
func ReadBulk(reader *bufio.Reader, head string) ([]byte, error) {
	var err error
	var data []byte

	if head == "" {
		head, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
	}
	switch head[0] {
	case ':':
		data = []byte(strings.TrimSpace(head[1:]))

	case '$':
		size, err := strconv.Atoi(strings.TrimSpace(head[1:]))
		if err != nil {
			return nil, err
		}
		if size == -1 {
			return nil, doesNotExist
		}
		lr := io.LimitReader(reader, int64(size))
		data, err = ioutil.ReadAll(lr)
		if err == nil {
			// read end of line
			_, err = reader.ReadString('\n')
		}
	default:
		return nil, RedisError("Expecting Prefix '$' or ':'")
	}

	return data, err
}

func WriteRequest(writer io.Writer, cmd string, args ...string) error {
	b := CommandBytes(cmd, args...)
	_, err := writer.Write(b)
	return err
}

func CommandBytes(cmd string, args ...string) []byte {
	var cmdbuf bytes.Buffer
	fmt.Fprintf(&cmdbuf, "*%d\r\n$%d\r\n%s\r\n", len(args)+1, len(cmd), cmd)
	for _, s := range args {
		fmt.Fprintf(&cmdbuf, "$%d\r\n%s\r\n", len(s), s)
	}
	return cmdbuf.Bytes()
}

func ReadResponse(reader *bufio.Reader) (interface{}, error) {

	var line string
	var err error

	//read until the first non-whitespace line
	for {
		line, err = reader.ReadString('\n')
		if len(line) == 0 || err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			break
		}
	}

	if line[0] == '+' {
		return strings.TrimSpace(line[1:]), nil
	}

	if strings.HasPrefix(line, "-ERR ") {
		errmesg := strings.TrimSpace(line[5:])
		return nil, RedisError(errmesg)
	}

	if line[0] == ':' {
		n, err := strconv.ParseInt(strings.TrimSpace(line[1:]), 10, 64)
		if err != nil {
			return nil, RedisError("Int reply is not a number")
		}
		return n, nil
	}

	if line[0] == '*' {
		size, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			return nil, RedisError("MultiBulk reply expected a number")
		}
		if size <= 0 {
			return make([][]byte, 0), nil
		}
		res := make([][]byte, size)
		for i := 0; i < size; i++ {
			res[i], err = ReadBulk(reader, "")
			if err == doesNotExist {
				continue
			}
			if err != nil {
				return nil, err
			}
			// dont read end of line as might not have been bulk
		}
		return res, nil
	}
	return ReadBulk(reader, line)
}
