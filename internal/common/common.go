package common

type ctxKey string

const CtxOperationKey string = "operation_id"

func MakeCtxKey(key string) ctxKey {
	return ctxKey(key)
}
