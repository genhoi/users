package jsonl

import (
	"encoding/json"
	"github.com/genhoi/users/module/user"
	"io"
)

func Generator(r io.Reader, chanSize int) <-chan user.GenerateUser {
	users := make(chan user.GenerateUser, chanSize)

	go func() {
		decoder := json.NewDecoder(r)
		for {
			result := user.GenerateUser{}
			err := decoder.Decode(&result.User)
			if err == io.EOF {
				break
			}
			result.Err = err
			users <- result
		}

		close(users)
	}()

	return users
}
