package api

type Context struct {
	TraceId string
	Vars    map[string]string
	Body    []byte
	With    []string
}
