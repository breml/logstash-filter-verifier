// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package logging

import (
	"sync"
)

// Ensure, that LoggerMock does implement Logger.
// If this is not the case, regenerate this file with moq.
var _ Logger = &LoggerMock{}

// LoggerMock is a mock implementation of Logger.
//
//	func TestSomethingThatUsesLogger(t *testing.T) {
//
//		// make and configure a mocked Logger
//		mockedLogger := &LoggerMock{
//			DebugFunc: func(args ...interface{})  {
//				panic("mock out the Debug method")
//			},
//			DebugfFunc: func(format string, args ...interface{})  {
//				panic("mock out the Debugf method")
//			},
//			ErrorFunc: func(args ...interface{})  {
//				panic("mock out the Error method")
//			},
//			ErrorfFunc: func(format string, args ...interface{})  {
//				panic("mock out the Errorf method")
//			},
//			FatalFunc: func(args ...interface{})  {
//				panic("mock out the Fatal method")
//			},
//			FatalfFunc: func(format string, args ...interface{})  {
//				panic("mock out the Fatalf method")
//			},
//			InfoFunc: func(args ...interface{})  {
//				panic("mock out the Info method")
//			},
//			InfofFunc: func(format string, args ...interface{})  {
//				panic("mock out the Infof method")
//			},
//			WarningFunc: func(args ...interface{})  {
//				panic("mock out the Warning method")
//			},
//			WarningfFunc: func(format string, args ...interface{})  {
//				panic("mock out the Warningf method")
//			},
//		}
//
//		// use mockedLogger in code that requires Logger
//		// and then make assertions.
//
//	}
type LoggerMock struct {
	// DebugFunc mocks the Debug method.
	DebugFunc func(args ...interface{})

	// DebugfFunc mocks the Debugf method.
	DebugfFunc func(format string, args ...interface{})

	// ErrorFunc mocks the Error method.
	ErrorFunc func(args ...interface{})

	// ErrorfFunc mocks the Errorf method.
	ErrorfFunc func(format string, args ...interface{})

	// FatalFunc mocks the Fatal method.
	FatalFunc func(args ...interface{})

	// FatalfFunc mocks the Fatalf method.
	FatalfFunc func(format string, args ...interface{})

	// InfoFunc mocks the Info method.
	InfoFunc func(args ...interface{})

	// InfofFunc mocks the Infof method.
	InfofFunc func(format string, args ...interface{})

	// WarningFunc mocks the Warning method.
	WarningFunc func(args ...interface{})

	// WarningfFunc mocks the Warningf method.
	WarningfFunc func(format string, args ...interface{})

	// calls tracks calls to the methods.
	calls struct {
		// Debug holds details about calls to the Debug method.
		Debug []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Debugf holds details about calls to the Debugf method.
		Debugf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Error holds details about calls to the Error method.
		Error []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Errorf holds details about calls to the Errorf method.
		Errorf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Fatal holds details about calls to the Fatal method.
		Fatal []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Fatalf holds details about calls to the Fatalf method.
		Fatalf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Info holds details about calls to the Info method.
		Info []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Infof holds details about calls to the Infof method.
		Infof []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Warning holds details about calls to the Warning method.
		Warning []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Warningf holds details about calls to the Warningf method.
		Warningf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
	}
	lockDebug    sync.RWMutex
	lockDebugf   sync.RWMutex
	lockError    sync.RWMutex
	lockErrorf   sync.RWMutex
	lockFatal    sync.RWMutex
	lockFatalf   sync.RWMutex
	lockInfo     sync.RWMutex
	lockInfof    sync.RWMutex
	lockWarning  sync.RWMutex
	lockWarningf sync.RWMutex
}

// Debug calls DebugFunc.
func (mock *LoggerMock) Debug(args ...interface{}) {
	if mock.DebugFunc == nil {
		panic("LoggerMock.DebugFunc: method is nil but Logger.Debug was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockDebug.Lock()
	mock.calls.Debug = append(mock.calls.Debug, callInfo)
	mock.lockDebug.Unlock()
	mock.DebugFunc(args...)
}

// DebugCalls gets all the calls that were made to Debug.
// Check the length with:
//
//	len(mockedLogger.DebugCalls())
func (mock *LoggerMock) DebugCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockDebug.RLock()
	calls = mock.calls.Debug
	mock.lockDebug.RUnlock()
	return calls
}

// Debugf calls DebugfFunc.
func (mock *LoggerMock) Debugf(format string, args ...interface{}) {
	if mock.DebugfFunc == nil {
		panic("LoggerMock.DebugfFunc: method is nil but Logger.Debugf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockDebugf.Lock()
	mock.calls.Debugf = append(mock.calls.Debugf, callInfo)
	mock.lockDebugf.Unlock()
	mock.DebugfFunc(format, args...)
}

// DebugfCalls gets all the calls that were made to Debugf.
// Check the length with:
//
//	len(mockedLogger.DebugfCalls())
func (mock *LoggerMock) DebugfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockDebugf.RLock()
	calls = mock.calls.Debugf
	mock.lockDebugf.RUnlock()
	return calls
}

// Error calls ErrorFunc.
func (mock *LoggerMock) Error(args ...interface{}) {
	if mock.ErrorFunc == nil {
		panic("LoggerMock.ErrorFunc: method is nil but Logger.Error was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	mock.lockError.Unlock()
	mock.ErrorFunc(args...)
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//
//	len(mockedLogger.ErrorCalls())
func (mock *LoggerMock) ErrorCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockError.RLock()
	calls = mock.calls.Error
	mock.lockError.RUnlock()
	return calls
}

// Errorf calls ErrorfFunc.
func (mock *LoggerMock) Errorf(format string, args ...interface{}) {
	if mock.ErrorfFunc == nil {
		panic("LoggerMock.ErrorfFunc: method is nil but Logger.Errorf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockErrorf.Lock()
	mock.calls.Errorf = append(mock.calls.Errorf, callInfo)
	mock.lockErrorf.Unlock()
	mock.ErrorfFunc(format, args...)
}

// ErrorfCalls gets all the calls that were made to Errorf.
// Check the length with:
//
//	len(mockedLogger.ErrorfCalls())
func (mock *LoggerMock) ErrorfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockErrorf.RLock()
	calls = mock.calls.Errorf
	mock.lockErrorf.RUnlock()
	return calls
}

// Fatal calls FatalFunc.
func (mock *LoggerMock) Fatal(args ...interface{}) {
	if mock.FatalFunc == nil {
		panic("LoggerMock.FatalFunc: method is nil but Logger.Fatal was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockFatal.Lock()
	mock.calls.Fatal = append(mock.calls.Fatal, callInfo)
	mock.lockFatal.Unlock()
	mock.FatalFunc(args...)
}

// FatalCalls gets all the calls that were made to Fatal.
// Check the length with:
//
//	len(mockedLogger.FatalCalls())
func (mock *LoggerMock) FatalCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockFatal.RLock()
	calls = mock.calls.Fatal
	mock.lockFatal.RUnlock()
	return calls
}

// Fatalf calls FatalfFunc.
func (mock *LoggerMock) Fatalf(format string, args ...interface{}) {
	if mock.FatalfFunc == nil {
		panic("LoggerMock.FatalfFunc: method is nil but Logger.Fatalf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockFatalf.Lock()
	mock.calls.Fatalf = append(mock.calls.Fatalf, callInfo)
	mock.lockFatalf.Unlock()
	mock.FatalfFunc(format, args...)
}

// FatalfCalls gets all the calls that were made to Fatalf.
// Check the length with:
//
//	len(mockedLogger.FatalfCalls())
func (mock *LoggerMock) FatalfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockFatalf.RLock()
	calls = mock.calls.Fatalf
	mock.lockFatalf.RUnlock()
	return calls
}

// Info calls InfoFunc.
func (mock *LoggerMock) Info(args ...interface{}) {
	if mock.InfoFunc == nil {
		panic("LoggerMock.InfoFunc: method is nil but Logger.Info was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockInfo.Lock()
	mock.calls.Info = append(mock.calls.Info, callInfo)
	mock.lockInfo.Unlock()
	mock.InfoFunc(args...)
}

// InfoCalls gets all the calls that were made to Info.
// Check the length with:
//
//	len(mockedLogger.InfoCalls())
func (mock *LoggerMock) InfoCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockInfo.RLock()
	calls = mock.calls.Info
	mock.lockInfo.RUnlock()
	return calls
}

// Infof calls InfofFunc.
func (mock *LoggerMock) Infof(format string, args ...interface{}) {
	if mock.InfofFunc == nil {
		panic("LoggerMock.InfofFunc: method is nil but Logger.Infof was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockInfof.Lock()
	mock.calls.Infof = append(mock.calls.Infof, callInfo)
	mock.lockInfof.Unlock()
	mock.InfofFunc(format, args...)
}

// InfofCalls gets all the calls that were made to Infof.
// Check the length with:
//
//	len(mockedLogger.InfofCalls())
func (mock *LoggerMock) InfofCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockInfof.RLock()
	calls = mock.calls.Infof
	mock.lockInfof.RUnlock()
	return calls
}

// Warning calls WarningFunc.
func (mock *LoggerMock) Warning(args ...interface{}) {
	if mock.WarningFunc == nil {
		panic("LoggerMock.WarningFunc: method is nil but Logger.Warning was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockWarning.Lock()
	mock.calls.Warning = append(mock.calls.Warning, callInfo)
	mock.lockWarning.Unlock()
	mock.WarningFunc(args...)
}

// WarningCalls gets all the calls that were made to Warning.
// Check the length with:
//
//	len(mockedLogger.WarningCalls())
func (mock *LoggerMock) WarningCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockWarning.RLock()
	calls = mock.calls.Warning
	mock.lockWarning.RUnlock()
	return calls
}

// Warningf calls WarningfFunc.
func (mock *LoggerMock) Warningf(format string, args ...interface{}) {
	if mock.WarningfFunc == nil {
		panic("LoggerMock.WarningfFunc: method is nil but Logger.Warningf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockWarningf.Lock()
	mock.calls.Warningf = append(mock.calls.Warningf, callInfo)
	mock.lockWarningf.Unlock()
	mock.WarningfFunc(format, args...)
}

// WarningfCalls gets all the calls that were made to Warningf.
// Check the length with:
//
//	len(mockedLogger.WarningfCalls())
func (mock *LoggerMock) WarningfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockWarningf.RLock()
	calls = mock.calls.Warningf
	mock.lockWarningf.RUnlock()
	return calls
}
