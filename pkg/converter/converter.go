package converter

import (
	"http-to-postman/pkg/models"
)

func ConvertToPostmanCollection(requests []models.Request) models.PostmanCollection {
	items := httpRequestsToPostmanItems(requests)

	return models.PostmanCollection{
		Info: models.CollectionInfo{
			Name:   "Converted Collection",
			Schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Item: items,
	}
}

func httpRequestsToPostmanItems(requests []models.Request) []models.PostmanItem {
	var items []models.PostmanItem
	for _, request := range requests {
		items = append(items, models.PostmanItem{
			Name:    request.Description,
			Request: httpRequestToPostmanRequest(request),
		})
	}

	return items
}

func httpRequestToPostmanRequest(request models.Request) models.PostmanRequest {

	postmanUrl := models.PostmanURL{
		Raw: request.Url,
	}

	var postmanHeaders []models.PostmanHeader
	for _, header := range request.Headers {
		postmanHeaders = append(postmanHeaders, models.PostmanHeader{
			Key:   header.Key,
			Value: header.Value,
		})
	}

	postmanBody := models.PostmanBody{
		Mode: "raw", // TODO: Support other modes
		Raw:  request.Body,
		Options: models.PostmanBodyOptions{
			Raw: models.PostmanBodyOptionsRaw{
				Language: "json",
			},
		},
	}

	return models.PostmanRequest{
		Method: request.Method,
		Url:    postmanUrl,
		Header: postmanHeaders,
		Body:   postmanBody,
	}

}
