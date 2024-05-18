package apimanage

import "io"

type apiReadCloser struct {
	index   int
	content []byte
}

func NewAPIReadCloser(content []byte) *apiReadCloser {
	return &apiReadCloser{
		index:   0,
		content: content,
	}
}

// func (r *apiReadCloser) Len() int {
// 	return len(r.content)
// }

func (r *apiReadCloser) Read(p []byte) (n int, err error) {
	for r.index < len(r.content) {
		if n < len(p) {
			p[n] = r.content[r.index]
		} else {
			break
		}
		n++
		r.index++
	}
	if r.index == len(r.content) {
		err = io.EOF
	}
	return
}

func (r *apiReadCloser) Close() error {
	return nil
}
