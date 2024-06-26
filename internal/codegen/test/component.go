package test

import (
	"context"
	"fmt"

	"github.com/kanengo/akasar"
)

type ID int64

type Reverser interface {
	Reverse(context.Context, string, *arg1, ID) (string, ID, res1, error)
	ZNoRetryMethod(ctx context.Context) error
}

var _ akasar.NotRetriable = Reverser.ZNoRetryMethod

type Foo struct {
}

type reverser struct {
	akasar.Components[Reverser]
	foo *Foo
}

type arg1 struct {
	akasar.AutoMarshal
	Id     int64
	Name   string
	Skills []int32
	a2     *arg2
}

type arg2 struct {
	akasar.AutoMarshal
	Age int
}

type res1 struct {
	akasar.AutoMarshal
	Id int
}

func (r *reverser) ZNoRetryMethod(ctx context.Context) error {
	return nil
}

func (*reverser) Reverse(_ context.Context, s string, arg1 *arg1, id ID) (string, ID, res1, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes), 0, res1{}, nil
}

type Error struct {
	akasar.AutoMarshal
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d-%s", e.Code, e.Message)
}
