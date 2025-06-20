package utils

func If[T any](cond bool, onTrue, onFalse T) T {
	if cond {
		return onTrue
	} else {
		return onFalse
	}
}
