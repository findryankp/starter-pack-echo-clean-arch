package delivery

import "cleanarc/features/categories"

type Request struct {
	Category string `json:"category" form:"category"`
}

func RequestToCore(request *Request) categories.Core {
	return categories.Core{
		Category: request.Category,
	}
}
