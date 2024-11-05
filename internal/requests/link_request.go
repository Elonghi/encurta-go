package requests

type LinkRequest struct {
	URL            string `json:"url" binding:"required"`
	CustomShortURL string `json:"custom_short_url"`
	QrCode         bool   `json:"qr_code"`
}
