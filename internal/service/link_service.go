package service

import (
	"errors"
	"log"

	"github.com/elonghi/encurtago/internal/domain"
	"github.com/elonghi/encurtago/internal/helpers"
	"github.com/elonghi/encurtago/internal/repository"
	"github.com/elonghi/encurtago/internal/requests"
)

// LinkService fornece métodos para encurtar e expandir URLs.
type LinkService struct {
	linkRepo repository.LinkRepository
}

// NewLinkService cria uma nova instância de LinkService.
func NewLinkService(linkRepository repository.LinkRepository) *LinkService {
	return &LinkService{linkRepo: linkRepository}
}

// ShortenURL encurta uma URL fornecida ou retorna uma existente se já encurtada.
func (s *LinkService) ShortenURL(req requests.LinkRequest) (domain.Link, error) {
	// Valida a URL fornecida.
	if !helpers.CheckURL(req.URL) {
		return domain.Link{}, errors.New("URL inválida")
	}

	// Verifica se a URL já foi encurtada anteriormente.
	existingLink, err := s.linkRepo.FindByUrl(req.URL)
	if req.CustomShortURL != "" {
		_, err := s.linkRepo.FindByShortURL(req.CustomShortURL)
		if err == nil {
			return domain.Link{}, errors.New("URL curta personalizada já existe")
		}

		shortURL := req.CustomShortURL
		existingLink.ShortURL = shortURL
		log.Print(existingLink)
		s.linkRepo.Update(existingLink)
		return existingLink, nil
	}
	if err == nil && existingLink.ShortURL != "" {
		return existingLink, nil
	}
	// Gera a URL curta (pode ser personalizada).
	shortURL := helpers.GenerateShortURL(req.URL)
	if req.CustomShortURL != "" {
		shortURL = req.CustomShortURL
	}

	// Cria um novo domínio Link.
	newLink := domain.Link{
		URL:      req.URL,
		ShortURL: shortURL,
	}

	// Gera o QR Code se solicitado.
	if req.QrCode {
		qrCode, err := helpers.GenerateQrCode(shortURL)
		if err != nil {
			log.Printf("Erro ao gerar QR Code: %v", err)
		} else {
			newLink.QrCode = qrCode
		}
	}

	// Salva o novo Link no repositório.
	createdLink, err := s.linkRepo.Create(newLink)
	if err != nil {
		return domain.Link{}, err
	}

	log.Printf("Novo link criado: %+v", createdLink)

	return createdLink, nil
}

// UnshortenURL expande uma URL curta para sua URL original.
func (s *LinkService) UnshortenURL(shortURL string) (string, error) {

	// Se não estiver no cache, busca no repositório.
	link, err := s.linkRepo.FindByShortURL(shortURL)
	if err != nil {
		return "", err
	}

	return link.URL, nil
}
