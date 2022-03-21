package errors

import "github.com/streamwest-1629/go-utilpkgs/kv"

type (
	Instance struct {
		message string
		reason  error
		kvs     []kv.StrStr
	}
)

func NewKV(key, value string) kv.StrStr {
	return KV(key, value)
}

func KV(key, value string) kv.StrStr {
	return kv.StrStr{
		Key:   key,
		Value: value,
	}
}

func New(msg string, kvs ...kv.StrStr) error {
	return NewWith(msg, nil, kvs...)
}

func NewWith(msg string, reason error, kvs ...kv.StrStr) error {
	return Instance{
		message: msg,
		reason:  reason,
		kvs:     kvs,
	}
}

func (err Instance) Error() string {

	builder := err.message

	// join key/value pairs
	if length := len(err.kvs); length > 0 {
		builder += " ("
		for i := range err.kvs {
			builder += err.kvs[i].Key + ": " + err.kvs[i].Value
			if i+1 < length {
				builder += ", "
			} else {
				builder += ")"
			}
		}
	}

	// join reason error
	if err.reason != nil {
		builder += ": " + err.reason.Error()
	}

	return builder
}
