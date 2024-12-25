package api

type InferenceJobRequest struct {
	Models        InferenceJobModels        `json:"models"`
	Transcription InferenceJobTranscription `json:"transcription"`
	Urls          []string                  `json:"urls"`
	Text          []string                  `json:"text"`
	CallbackUrl   string                    `json:"callback_url"`
	Notify        bool                      `json:"notify"`
}

type InferenceJobResponse struct {
	JobID string `json:"job_id"`
}

type InferenceJobTranscription struct {
	// BCP-47 tag
	Language            string  `json:"language"`
	IdentifySpeakers    bool    `json:"identify_speakers"`
	ConfidenceThreshold float64 `json:"confidence_threshold"`
}

type InferenceJobModels struct {
	Face     InferenceJobModelFace     `json:"face"`
	Burst    []string                  `json:"burst"`
	Prosody  InferenceJobModelProsody  `json:"prosody"`
	Language InferenceJobModelLanguage `json:"language"`
	Ner      InferenceJobModelNer      `json:"ner"`
	Facemesh string                    `json:"facemesh"`
}

type InferenceJobModelFace struct {
	FpsPred       float64  `json:"fps_pred"`
	ProbThreshold float64  `json:"prob_threshold"`
	IdentifyFaces bool     `json:"identify_faces"`
	MinFaceSize   uint64   `json:"min_face_size"`
	Facs          []string `json:"facs"`
	Descriptions  []string `json:"descriptions"`
	SaveFaces     bool     `json:"save_faces"`
}

type InferenceJobModelProsody struct {
	// Allowed values: word, sentence, utterance, conversational_turn
	Granularity      string                         `json:"granularity"`
	Window           InferenceJobModelProsodyWindow `json:"window"`
	IdentifySpeakers bool                           `json:"identify_speakers"`
}

type InferenceJobModelProsodyWindow struct {
	Length float64 `json:"length"`
	Step   float64 `json:"step"`
}

type InferenceJobModelLanguage struct {
	// Allowed values: word, sentence, utterance, conversational_turn
	Granularity      string `json:"granularity"`
	Sentiment        string `json:"sentiment"`
	Toxicity         string `json:"toxicity"`
	IdentifySpeakers bool   `json:"identify_speakers"`
}

type InferenceJobModelNer struct {
	IdentifySpeakers bool `json:"identify_speakers"`
}

type JobDetailsResponse struct {
	JobID   string              `json:"job_id"`
	Request InferenceJobRequest `json:"request"`
	State   JobState            `json:"state"`
	Type    string              `json:"type"`
}

type JobState struct {
	Status             string `json:"status"`
	CreatedTimestampMs int64  `json:"created_timestamp_ms"`
	StartedTimestampMs int64  `json:"started_timestamp_ms"`
	EndedTimestampMs   int64  `json:"ended_timestamp_ms"`
	NumErrors          uint64 `json:"num_errors"`
	NumPredictions     uint64 `json:"num_predictions"`
	Message            string `json:"message"`
}
