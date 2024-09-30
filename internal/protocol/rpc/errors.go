package rpc

// Helper structure for listeners (responders, receivers) boilerplate-less initialization.
type ErrorsCatcher struct {
	errors []error
}

func (ec *ErrorsCatcher) Catch(err error) {
	if err != nil {
		ec.errors = append(ec.errors, err)
	}
}

func (ec *ErrorsCatcher) FirstError() error {
	if len(ec.errors) > 0 {
		return ec.errors[0]
	}

	return nil
}
