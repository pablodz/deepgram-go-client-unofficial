package models

type Response struct {
	ChannelIndex []int        `json:"channel_index"`
	Duration     float64      `json:"duration"`
	Start        float64      `json:"start"`
	IsFinal      bool         `json:"is_final"`
	SpeechFinal  bool         `json:"speech_final"`
	Channel      Alternatives `json:"channel"`
	Metadata     Metadata     `json:"metadata"`
}

type Alternatives struct {
	Alternatives []struct {
		Transcript string  `json:"transcript"`
		Confidence float64 `json:"confidence"`
		Words      []struct {
			Word       string  `json:"word"`
			Start      float64 `json:"start"`
			End        float64 `json:"end"`
			Confidence float64 `json:"confidence"`
		} `json:"words"`
	} `json:"alternatives"`
}
type Metadata struct {
	RequestId string `json:"request_id"`
	ModelUUID string `json:"model_uuid"`
}
