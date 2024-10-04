package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const userId = "1fe15ae2-9702-4f17-aab0-61088f15c57b"

// func TestGenerateId(t *testing.T) {
// 	result := GenerateId()

// 	fmt.Println(result)
// }

func TestShortLinkGenerate(t *testing.T) {
	shortUrlGen := NewLinkGeneratorImpl()

	testCases := []struct {
		originalUrl   string
		expectedShort string
	}{
		{
			originalUrl:   "https://www.apple.com/id/newsroom/2024/09/apple-debuts-iphone-16-pro-iphone-16-pro-max/",
			expectedShort: "L7nDK",
		},
		{
			originalUrl:   "https://www.samsung.com/id/smartphones/galaxy-a/galaxy-a54-5g-green-256gb-sm-a546elgdxid/buy/",
			expectedShort: "NZyTJ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.originalUrl, func(t *testing.T) {
			shortUrl := shortUrlGen.GenerateShortLink(tc.originalUrl, userId)
			assert.Equal(t, tc.expectedShort, shortUrl)
		})
	}
}
