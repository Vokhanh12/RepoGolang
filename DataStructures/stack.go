package DataStructures

type Stack struct {
	data []interface{}
	top  int
}

func (s *Stack) Push(element interface{}) {

	s.top++
	s.data = append(s.data, element)

}

func (s *Stack) Pop() interface{} {

	if len(s.data) > 0 {

		s.top--

		last := s.data[s.top]

		s.data = s.data[:s.top]

		return last

	}

	return nil

}
