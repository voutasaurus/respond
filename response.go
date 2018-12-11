package response

import "net/http"

type Transformer struct {
	http.ResponseWriter

	Transform func(b []byte) []byte
}

func (r *Transformer) Write(b []byte) (int, error) {
	return r.ResponseWriter.Write(r.transform(b))
}

func (r *Transformer) transform(b []byte) []byte {
	if r.Transform == nil {
		return b
	}
	return r.Transform(b)
}
