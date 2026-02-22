package domain

type Prediction struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	Confidence float64 `json:"confidence"`
	Class      string  `json:"class"`
	ClassID    int     `json:"class_id"`
}

type RoboflowResponse struct {
	Predictions []Prediction `json:"predictions"`
}