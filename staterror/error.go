package staterror

type StatError struct {
	s string
}

func (s *StatError) Error() string {
	return s.s
}

func New(text string) error {
	return &StatError{text}
}
