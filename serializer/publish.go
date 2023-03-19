package serializer

type PublishResponse struct {
	Response
	VideoUrl string `json:"video_url,omitempty"`
}
