package responses

import "github.com/elonghi/encurtago/internal/domain"

type LinkResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
	QrCode   string `json:"qr_code"`
}

func SetLinkResponse(link domain.Link) LinkResponse {
	return LinkResponse{
		URL:      link.URL,
		ShortURL: link.ShortURL,
		QrCode:   link.QrCode,
	}
}
