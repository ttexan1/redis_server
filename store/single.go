package store

import (
	"redis_app/domain"
	"sync"
)

type single struct {
	mutex    sync.Mutex
	singleDB map[string]domain.Single
}

// GetValue is a getvalue
func (ss *single) GetValue(key string) (domain.Single, *domain.Error) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	info, ok := ss.singleDB[key]
	if !ok {
		return domain.Single{}, domain.NewError(
			domain.RespErrorNilValue, nil,
		)
	}
	return info, nil
}

// SetValue set the data from given args
func (ss *single) SetValue(key string, elem domain.Single) *domain.Error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.singleDB[key] = elem
	return nil // いったんエラーは想定されていない
}

// Delete delte value
func (ss *single) Delete(key string) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	delete(ss.singleDB, key)
}
