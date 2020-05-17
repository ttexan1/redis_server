package usecase

import (
	"redis_app/command"
	"redis_app/domain"
	"strconv"
	"strings"
)

// Single is the interface for single key value command
type Single interface {
	Get(string) domain.RespString
	Set(string, string, []string) domain.RespString
	Del([]string) domain.RespString
	IncrBy(string, string) domain.RespString
}

// NewSingle returns the Single interface
func (uc *UseCase) NewSingle() Single {
	return &single{
		SingleStore: uc.SingleStore,
	}
}

// SingleStore is the interface for single structure data store process
type SingleStore interface {
	GetValue(string) (domain.Single, *domain.Error)
	SetValue(string, domain.Single) *domain.Error
	Delete(string)
}

type single struct {
	SingleStore
}

func (sg *single) Get(key string) domain.RespString {
	data, err := sg.GetValue(key)
	if err != nil {
		return err.RespError()
	}
	return data.RespString()
}

func (sg *single) Set(key, value string, options []string) domain.RespString {
	data, err := sg.GetValue(key)
	if err != nil && err.Message != domain.RespErrorNilValue {
		return err.RespError()
	}
	// if NX or  XX option exist
	for _, option := range options {
		switch strings.ToLower(option) {
		case command.SetOptionNX:
			if err == nil {
				return domain.RespErrorNilValue
			}
			if err := sg.SetValue(key, domain.Single{
				Value: value,
			}); err != nil {
				return err.RespError()
			}
			return domain.RespOK
		case command.SetOptionXX:
			if err != nil {
				return err.RespError()
			}
			data.Value = value
			if err := sg.SetValue(key, data); err != nil {
				return err.RespError()
			}
			return domain.RespOK
		}
	}
	if err := sg.SetValue(key, domain.Single{Value: value}); err != nil {
		return err.RespError()
	}
	return domain.RespOK
}

func (sg *single) IncrBy(key, value string) domain.RespString {
	data, err := sg.GetValue(key)
	if err != nil && err.Message != domain.RespErrorNilValue {
		return err.RespError()
	}

	// check if value is int
	val, er := strconv.Atoi(value)
	if er != nil {
		return domain.RespErr(er.Error())
	}

	oldVal, err := data.IntValue()
	if err != nil {
		return err.RespError()
	}
	newVal := val + oldVal
	data.Value = string(newVal)
	if err := sg.SetValue(key, data); err != nil {
		return err.RespError()
	}
	return domain.RespInteger(newVal)
}

// func (sg *single) Incr(key string) domain.RespString {
// 	return sg.IncrBy(key, "1")
// }

// func (sg *single) Decr(key string) domain.RespString {
// 	return sg.IncrBy(key, "-1")
// }

// func (sg *single) DecrBy(key, value string) domain.RespString {
// 	return sg.IncrBy(key, "-"+value)
// }

func (sg *single) Del(keys []string) domain.RespString {
	count := 0
	for _, key := range keys {
		if _, err := sg.GetValue(key); err == nil {
			sg.Delete(key)
			count++
		}
	}
	return domain.RespInteger(count)
}
