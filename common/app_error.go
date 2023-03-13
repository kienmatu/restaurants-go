package common

type AppError struct {
	StatusCode int    `json:"statusCode"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Key        string `json:"errorKey"`
	Log        string `json:"log"`
}
