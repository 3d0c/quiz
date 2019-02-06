package sessions

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
)

const (
	KeyLength = 8
)

var (
	once     sync.Once
	instance *sessions
)

type session struct {
	score    int
	answered map[int]bool
}

type sessions struct {
	data map[string]*session
	sync.Mutex
}

func Instance() *sessions {
	once.Do(func() {
		instance = &sessions{
			data: make(map[string]*session),
		}
	})

	return instance
}

func (s *sessions) Create() (string, error) {
	var (
		ok  bool
		key string
		err error
	)

	s.Lock()
	defer s.Unlock()

	if key, err = s.generateKey(); err != nil {
		return "", err
	}

	if _, ok = s.data[key]; ok {
		return "", errors.New("non uniq key generated")
	}

	s.data[key] = &session{score: 0, answered: make(map[int]bool)}

	return key, nil
}

func (s *sessions) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.data[key]; !ok {
		return errors.New("Session not found")
	}

	delete(s.data, key)

	return nil
}

func (s *sessions) getSession(key string) (*session, error) {
	var (
		ok     bool
		result *session
	)

	if result, ok = s.data[key]; !ok {
		return nil, errors.New("session not found")
	}
	if result == nil {
		return nil, errors.New("unexpected error")
	}

	return result, nil
}

func (s *sessions) Contains(key string) bool {
	s.Lock()
	defer s.Unlock()

	_, ok := s.data[key]

	return ok
}

func (s *sessions) GetScore(key string) (int, error) {
	var (
		sn  *session
		err error
	)

	s.Lock()
	defer s.Unlock()

	if sn, err = s.getSession(key); err != nil {
		return 0, err
	}

	return sn.score, nil
}

func (s *sessions) IncrementScore(key string) (int, error) {
	var (
		sn  *session
		err error
	)

	s.Lock()
	defer s.Unlock()

	if sn, err = s.getSession(key); err != nil {
		return 0, err
	}

	sn.score += 1

	return sn.score, nil
}

func (s *sessions) CompareScore(key string) (int, error) {
	var (
		sn  *session
		err error
	)

	s.Lock()
	defer s.Unlock()

	if sn, err = s.getSession(key); err != nil {
		return 0, err
	}

	total := len(s.data) - 1
	betterThan := 0

	if total == 0 {
		return 100, nil
	}

	for k, v := range s.data {
		if k == key {
			continue
		}

		if sn.score > v.score {
			betterThan += 1
		}
	}

	if betterThan == 0 {
		return 0, nil
	}

	return betterThan * 100 / total, nil
}

func (s *sessions) IsAnswered(key string, qid int) (bool, error) {
	var (
		sn  *session
		err error
	)

	s.Lock()
	defer s.Unlock()

	if sn, err = s.getSession(key); err != nil {
		return false, err
	}

	return sn.answered[qid], nil
}

func (s *sessions) SetAnswered(key string, qid int) error {
	var (
		sn  *session
		err error
	)

	s.Lock()
	defer s.Unlock()

	if sn, err = s.getSession(key); err != nil {
		return err
	}

	sn.answered[qid] = true

	return nil
}

func (s *sessions) generateKey() (string, error) {
	b := make([]byte, KeyLength)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
