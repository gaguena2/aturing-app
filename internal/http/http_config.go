package http

func Loader() map[string]string {
	return map[string]string{
		"HOST_SECRET": "http://127.0.0.1:8200/v1/secret/data/aturing-app",
		"HOST_TOKEN":  "giropops",
	}
}
