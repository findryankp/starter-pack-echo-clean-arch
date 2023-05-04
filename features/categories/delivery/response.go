package delivery

import "cleanarc/features/categories"

type Response struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
}

func CoreToResponse(dataCore categories.Core) Response {
	return Response{
		Id:       dataCore.Id,
		Category: dataCore.Category,
	}
}

func ListCoreToResponse(dataCore []categories.Core) []Response {
	var dataResponse []Response
	for _, v := range dataCore {
		dataResponse = append(dataResponse, CoreToResponse(v))
	}
	return dataResponse
}
