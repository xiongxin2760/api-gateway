package apimanage

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

func (r *apiReadCloser) Read(p []byte) (n int, err error) {
	for i := r.index; i < len(r.content); i++ {
		n++
		p = append(p, r.content[i])
		if n == len(p) {
			break
		}
	}
	return
}

func (r *apiReadCloser) Close() error {
	return nil
}
