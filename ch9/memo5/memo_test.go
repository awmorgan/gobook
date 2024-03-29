package memo_test

import (
	"testing"

	memo "github.com/awmorgan/gobook/ch9/memo5"
	"github.com/awmorgan/gobook/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T)  {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T)  {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}