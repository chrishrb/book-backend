package response

type Error struct {
	StatusCode int    `json:"statusCode"`
	Details    string `json:"details"`
}
