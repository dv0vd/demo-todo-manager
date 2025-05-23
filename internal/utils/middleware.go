package utils

func MiddlewareContentTypeCheck(header string) bool {
	return header == "application/json"
}
