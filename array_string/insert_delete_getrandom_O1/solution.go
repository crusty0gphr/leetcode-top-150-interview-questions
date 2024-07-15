package insert_delete_getrandom_O1

import "math/rand"

type RandomizedSet struct {
	data map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data: make(map[int]struct{}),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.data[val]; ok {
		return false
	}

	s.data[val] = struct{}{}
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.data[val]; !ok {
		return false
	}

	delete(s.data, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	randLen := rand.Intn(len(s.data))
	for key := range s.data {
		randLen--
		if randLen < 0 {
			return key
		}
	}
	return 0
}
