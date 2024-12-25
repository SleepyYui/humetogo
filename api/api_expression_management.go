package api

import "fmt"

func (h *HumeApi) ExprStartInferenceJob(request InferenceJobRequest) (*InferenceJobResponse, error) {
	h.printDebug("ExprStartInferenceJob")
	var inferenceJobResponse InferenceJobResponse

	url := fmt.Sprintf("%s/batch/jobs", h.BaseUrl)
	h.printDebug(url)

	err := h.post(url, request, &inferenceJobResponse)
	if err != nil {
		return nil, err
	}

	return &inferenceJobResponse, nil
}

func (h *HumeApi) ExprGetJobDetails(jobID string) (*JobDetailsResponse, error) {
	h.printDebug("ExprGetJobDetails")
	var jobDetailsResponse JobDetailsResponse

	url := fmt.Sprintf("%s/batch/jobs/%s", h.BaseUrl, jobID)
	h.printDebug(url)

	err := h.get(url, &jobDetailsResponse)
	if err != nil {
		return nil, err
	}

	return &jobDetailsResponse, nil
}
