package handlers

import "strings"

func buildMessage(key string, data interface{}) map[string]interface{} {
	return map[string]interface{}{key: data}
}

func getIdsParam(idsParam string) []string {
	var ids []string

	if idsParam != "" {
		for _, id := range strings.Split(idsParam, ",") {
			ids = append(ids, id)
		}
	}

	return ids
}
