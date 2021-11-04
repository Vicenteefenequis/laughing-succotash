package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomnomnom/linkheader"
	"strconv"
	"strings"
)

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

func getPaginationParam(c echo.Context) (int, int) {
	limit, _ := strconv.Atoi(c.QueryParam("per_page"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	var offset int

	if page != 0 && limit != 0 {
		offset = (limit * page) - limit
	}

	return limit, offset
}

func setPagination(c echo.Context, isLast bool) {
	limit := c.QueryParam("per_page")
	page := c.QueryParam("page")

	c.Response().Header().Set("X-Page-Num", limit)
	c.Response().Header().Set("X-Page-Size", page)

	link := linkheader.Links{
		{URL: "http://localhost:8080" + "?page=" + page + "&per_page=" + limit, Rel: "self"},
	}

	pageSize, _ := strconv.Atoi(page)

	if !isLast {
		link = append(link, linkheader.Link{URL: "http://localhost:8080" + "?page=" + strconv.Itoa(pageSize+1) + "&per_page=" + limit, Rel: "next"})
	}

	c.Response().Header().Set("Link", link.String())
}
