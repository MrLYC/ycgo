package ycgo

import (
	errs "errors"
	"fmt"
	"strings"
	"testing"
	"ycgo/errors"
)

func TestErrorNew(t *testing.T) {
	var (
		err     = errors.ErrorNew("mrlyc", 1)
		stack   = err.Stack()
		message = err.Message()
		info    = err.Error()
	)
	if strings.Index(message, "mrlyc") == -1 {
		t.Errorf("Message not found: %v", "mrlyc")
	}
	if strings.Index(info, "mrlyc") == -1 {
		t.Errorf("Info not found: %v", "mrlyc")
	}
	if strings.Index(stack, "TestErrorNew") == -1 {
		t.Errorf("Stack not found: %v", "TestErrorNew")
	}
	if strings.Index(info, "TestErrorNew") == -1 {
		t.Errorf("Info not found: %v", "TestErrorNew")
	}
}

func TestErrorf(t *testing.T) {
	var (
		name1 = "mrlyc1"
		name2 = "mrlyc2"
	)
	if strings.Index(errors.Errorf(name1).Message(), name1) == -1 {
		t.Errorf("Message not found: %v", name1)
	}
	if strings.Index(errors.Errorf("This is %v", name2).Message(), name2) == -1 {
		t.Errorf("Message not found: %v", name2)
	}
}

func TestErrorfInGoroutine(t *testing.T) {
	var (
		ch  = make(chan *errors.TraceableError)
		err *errors.TraceableError
	)
	go func() {
		e := errors.Errorf("mrlyc")
		ch <- e
	}()

	err = <-ch
	close(ch)
	if strings.Index(err.Stack(), "TestErrorfInGoroutine.func1") == -1 {
		t.Error(err)
	}
}

func TestErrorWrap(t *testing.T) {
	var (
		str  = "string"
		err  = errs.New("errors")
		terr = errors.Errorf("tracable")
		e    *errors.TraceableError
	)
	e = errors.ErrorWrap(str)
	if e.Message() != str || strings.Index(e.Stack(), "TestErrorWrap") == -1 {
		t.Error(e)
	}
	e = errors.ErrorWrap(err)
	if e.Message() != err.Error() || strings.Index(e.Stack(), "TestErrorWrap") == -1 {
		t.Error(e)
	}
	e = errors.ErrorWrap(terr)
	if *e != *terr {
		t.Error(e)
	}
	e = errors.ErrorWrap(*terr)
	if *e != *terr {
		t.Error(e)
	}
}

func TestTraceableErrorClone(t *testing.T) {
	var (
		err1 = errors.Errorf("mrlyc")
		err2 = err1.Clone()
	)
	if err1.Message() != err2.Message() || err1.Stack() == err2.Stack() {
		t.Errorf("error clone failed")
	}
}

func TestTraceableErrorPanic(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			err := errors.ErrorWrap((e))
			if strings.Index(err.Stack(), "TestTraceableErrorPanic") == -1 {
				t.Error(err)
			}
			if strings.Index(err.Message(), "mrlyc") == -1 {
				t.Error(err)
			}
		}
	}()
	errors.Errorf("mrlyc").Panic()
}

func TestErrorRecoverCall1(t *testing.T) {
	defer errors.ErrorRecoverCall(func(err *errors.TraceableError) {
		if strings.Index(err.Stack(), "TestErrorRecoverCall1") == -1 {
			t.Error(err)
		}
		if strings.Index(err.Message(), "mrlyc") == -1 {
			t.Error(err)
		}
	})
	errors.Errorf("mrlyc").Panic()
}

func TestErrorRecoverCall2(t *testing.T) {
	defer errors.ErrorRecoverCall(func(err *errors.TraceableError) {
		if strings.Index(err.Stack(), "TestErrorRecoverCall2") == -1 {
			t.Error(err)
		}
		if strings.Index(err.Message(), "mrlyc") == -1 {
			t.Error(err)
		}
	})
	panic("mrlyc")
}

func TestErrorRecoverCall3(t *testing.T) {
	defer errors.ErrorRecoverCall(func(err *errors.TraceableError) {
		if strings.Index(err.Stack(), "TestErrorRecoverCall3") == -1 {
			t.Error(err)
		}
		if strings.Index(err.Message(), "mrlyc") == -1 {
			t.Error(err)
		}
	})
	panic(fmt.Errorf("mrlyc"))
}

func TestErrorRecoverCall4(t *testing.T) {
	defer errors.ErrorRecoverCall(func(err *errors.TraceableError) {
		t.Errorf("will not happend")
	})
}

func TestErrorPrintf(t *testing.T) {
	err := errors.Errorf("mrlyc")
	errInfo := fmt.Sprintf("%v", err)
	if strings.Index(errInfo, "mrlyc") == -1 || strings.Index(errInfo, "TestErrorPrintf") == -1 {
		t.Errorf("Error message is unacceptable")
	}
}
