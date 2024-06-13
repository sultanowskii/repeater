package repeater

type resultChecker func(err error) bool

func IsError(err error) bool {
	return err != nil
}

func IsNoError(err error) bool {
	return err == nil
}
