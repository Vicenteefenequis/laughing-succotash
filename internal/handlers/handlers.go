package handlers


func buildMessage(key string,data interface{}) map[string]interface{} {
	return map[string]interface{}{key: data}
}