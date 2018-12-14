package response

import "net/http"

type Transformer struct {
	http.ResponseWriter

	Transform func(oldBody []byte) (newBody []byte, status int)
}

func (t *Transformer) Write(b []byte) (int, error) {
	return t.ResponseWriter.Write(t.transform(b))
}

func (t *Transformer) transform(b []byte) []byte {
	if t.Transform == nil {
		return b
	}
	out, status := t.Transform(b)
	if status != 0 {
		t.WriteHeader(status)
	}
	return out
}
