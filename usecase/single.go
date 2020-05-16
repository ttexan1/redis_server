package usecase

import (
	"fmt"
	"redis_app/command"
	"redis_app/domain"
	"strconv"
)

// Single is the interface for single key value command
type Single interface {
	Get(string) string
	Set(string, string, []string) string
	Del([]string) string
	IncrBy(string, string) string
}

// SingleStore is the interface
type SingleStore interface {
	GetValue(string) (domain.Single, error)
	SetValue(string, domain.Single) error
	Delete(string)
}

type single struct {
	SingleStore
}

// NewSingle returns the Single interface
func (uc *UseCase) NewSingle() Single {
	return &single{
		SingleStore: uc.SingleStore,
	}
}

func (sg *single) Get(key string) string {
	res, err := sg.GetValue(key)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("$%d\r\n%s", res.Length, res.Value)
}

func (sg *single) Set(key, value string, options []string) string {
	res, err := sg.GetValue(key)
	if err != nil && err.Error() != domain.ErrorTypeNilValue {
		return err.Error() // todo フォーマットを整形
	}
	for _, option := range options {
		switch option {
		case command.SetOptionNX:
			if err.Error() != domain.ErrorTypeNilValue {
				return domain.ErrorTypeNilValue
			}
			res = domain.Single{
				Value:  value,
				Length: len(value),
			}
			if err := sg.SetValue(key, res); err != nil {
				return err.Error()
			}
			return domain.ResponseOK
		case command.SetOptionXX:
			if err.Error() == domain.ErrorTypeNilValue {
				return domain.ErrorTypeNilValue
			}
			res.Value = value
			res.Length = len(value)
			if err := sg.SetValue(key, res); err != nil {
				return err.Error()
			}
			return domain.ResponseOK
		}
	}
	if err := sg.SetValue(key, domain.Single{
		Value:  value,
		Length: len(value),
	}); err != nil {
		return err.Error()
	}
	return domain.ResponseOK
}

func (sg *single) Del(keys []string) string {
	count := 0
	for _, key := range keys {
		if _, err := sg.GetValue(key); err == nil {
			sg.Delete(key)
			count++
		}
	}
	return fmt.Sprintf(":%d", count)
}

func (sg *single) IncrBy(key, value string) string {
	val, err := strconv.Atoi(value)
	if err != nil {
		return "-ERR" + err.Error()
	}

	res, err := sg.GetValue(key)
	if err != nil && err.Error() != domain.ErrorTypeNilValue {
		return err.Error() // todo フォーマットを整形
	}

	if res.Value == "" {
		res.Value = value
		if err := sg.SetValue(key, res); err != nil {
			return err.Error() // todo フォーマットを整形
		}
		return fmt.Sprintf("$%d\r\n%s\r\n", res.Length, res.Value)
	}

	oldVal, err := res.IntValue()
	if err != nil {
		return err.Error() // todo フォーマットを整形
	}
	newVal := string(val + oldVal)
	res.Value = newVal
	res.Length = len(newVal)
	if err := sg.SetValue(key, res); err != nil {
		return err.Error() // todo フォーマットを整形
	}
	return fmt.Sprintf(":%s\r\n", res.Value)
}
