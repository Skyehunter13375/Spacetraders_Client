package Model

type ApiError struct {
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
	ReqId   string   `json:"requestId"`
}

//	{
//		"error": {
//			"code": 4103,
//			"message": "Missing Bearer token in the request. Did you confirm sending the 'Bearer {token}' as the authorization header?",
//			"data": {},
//			"requestId": "019b188c-b106-74cd-a86b-ecafdcbb2e8d"
//		}
//	}

//	{
//		"error": {
//			"code": 4111,
//			"message": "Cannot register agent. Agent symbol NULLSKY has already been claimed.",
//			"data": {
//				"agentSymbol": "NULLSKY"
//			},
//			"requestId": "019bd743-87b6-760a-af4d-94a503d1f31c"
//		}
//	}

