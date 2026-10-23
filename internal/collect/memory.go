package collect

type Memory struct {
	Total     uint64
	Available uint64
}

func parseMemory(content string) (Memory, error) {

	return Memory{}, nil
}

func ReadMemory() (Memory, error) {

	return Memory{}, nil
}

func (m Memory) Used() uint64 {
	return m.Total - m.Available
}
