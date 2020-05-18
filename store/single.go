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
	// ss.mutex.Lock()
	// defer ss.mutex.Unlock()
	ss.singleDB[key] = elem
	return nil // いったんエラーは想定されていない
}

// IncrSet is for INCR function to execute atomic
func (ss *single) IncrSet(key string, val int) (int, *domain.Error) {
	// oldVal := 0
	// ss.mutex.Lock()
	// defer ss.mutex.Unlock()
	// if data, ok := ss.singleDB[key]; ok {
	// 	var err error
	// 	if oldVal, err = strconv.Atoi(data.Value); err != nil {
	// 		return 0, &domain.Error{Message: domain.RespErrorWrongArgumentType}
	// 	}
	// }
	// ss.singleDB[key] = domain.Single{
	// 	Value: strconv.Itoa(oldVal + val),
	// }
	// return oldVal + val, nil
	return 0, nil
}

// Delete delte value
func (ss *single) Delete(key string) {
	delete(ss.singleDB, key)
}
