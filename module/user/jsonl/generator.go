package jsonl

import (
	"bufio"
	"encoding/json"
	"github.com/genhoi/users/module/user"
	"io"
)

func Generator(r io.Reader, chanSize int) <-chan user.GenerateUser {
	users := make(chan user.GenerateUser, chanSize)

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			result := user.GenerateUser{}

			err := scanner.Err()
			if err != nil {
				result.Err = err
				users <- result
				continue
			}

			result.Err = json.Unmarshal(scanner.Bytes(), &result.User)
			users <- result
		}

		close(users)
	}()

	return users
}
