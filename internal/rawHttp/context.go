package rawHttp

type Context struct{
	headers map[string]interface{}
	body []byte
}