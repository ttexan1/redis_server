package store

import (
	"redis_app/domain"
)

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
	ss.singleDB[key] = elem
	return nil
}

// Delete delte value
func (ss *single) Delete(key string) {
	delete(ss.singleDB, key)
}
