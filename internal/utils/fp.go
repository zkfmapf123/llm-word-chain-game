package utils

type CheckFuncParmas[T any] struct {
	Fn  func(v T) bool
	Err error
}

func CheckParallel[T any](params T, fns ...CheckFuncParmas[T]) error {
	for _, fn := range fns {
		if fn.Fn(params) {
			return fn.Err
		}
	}

	return nil
}

// userWords 사용자가 입력한 단어
// loadWords 불러온 단어
// 불러온단어의 끝글자가 사용자가 입력단어의 시작이어야 함
func ComparisonVarsPrefixSuffix(userWord, loadWord string) bool {
	u, l := []rune(userWord), []rune(loadWord)

	if len(u) == 0 || len(l) == 0 {
		return false
	}

	return l[len(l)-1] == u[0]
}
