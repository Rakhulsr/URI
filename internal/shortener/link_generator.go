package shortener

type LinkGenarator interface {
	GenerateShortLink(originalUrl, userId string) string
}
