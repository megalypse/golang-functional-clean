package common

type ctxKey string

func MakeCtxKey(key string) ctxKey {
	return ctxKey(key)
}
