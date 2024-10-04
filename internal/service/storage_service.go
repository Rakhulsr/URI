package service

type StorageService interface {
	SaveUrlMap(originalUrl, userId string) string
	RetrieveRealUrl(shortUrl string) string
}
