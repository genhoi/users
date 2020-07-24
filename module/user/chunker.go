package user

func Chunk(entities <-chan User, size int) <-chan []User {
	out := make(chan []User)

	i := 0
	go func() {
		chunk := make([]User, 0, size)
		for e := range entities {
			chunk = append(chunk, e)
			i++

			if i == size {
				out <- chunk

				i = 0
				chunk = make([]User, 0, size)
			}
		}

		if i > 0 {
			out <- chunk
		}

		close(out)
	}()

	return out
}

func ChunkGenerate(entities <-chan GenerateUser, size int) <-chan []GenerateUser {
	out := make(chan []GenerateUser)

	i := 0
	go func() {
		chunk := make([]GenerateUser, 0, size)
		for e := range entities {
			chunk = append(chunk, e)
			i++

			if i == size {
				out <- chunk

				i = 0
				chunk = make([]GenerateUser, 0, size)
			}
		}

		if i > 0 {
			out <- chunk
		}

		close(out)
	}()

	return out
}
