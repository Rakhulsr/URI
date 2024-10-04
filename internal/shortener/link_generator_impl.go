package shortener

import (
	"fmt"
	"math/big"

	"github.com/Rakhulsr/go-url-shortener/internal/utils"
)

type LinkGenaratorImpl struct{}

func NewLinkGeneratorImpl() *LinkGenaratorImpl {
	return &LinkGenaratorImpl{}
}

// func GenerateId() string {
// 	id := uuid.NewString()

// 	return id

// }

func (lg *LinkGenaratorImpl) GenerateShortLink(originalUrl, userId string) string {

	hash := utils.Sha256Of(originalUrl + userId)
	genNumber := new(big.Int).SetBytes(hash).Uint64()
	shortLink := utils.Base56Encode([]byte(fmt.Sprint(genNumber)))
	return shortLink[:5]

}
