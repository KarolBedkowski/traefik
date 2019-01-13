package accesslog

import "io"

type captureRequestReader struct {
	source   io.ReadCloser
	count    int64
	body     []byte
	saveBody bool
}

func (r *captureRequestReader) Read(p []byte) (int, error) {
	n, err := r.source.Read(p)
	r.body = append(r.body, p...)
	if r.saveBody {
		r.count += int64(n)
	}
	return n, err
}

func (r *captureRequestReader) Close() error {
	return r.source.Close()
}
