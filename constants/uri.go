package constants

import "fmt"

const (
	ImageCDN = "https://cdn.depocket.com/images"
)

func GetTokenIcon(chain string, address string) string {
	return fmt.Sprintf("%s/%s/%s.png", ImageCDN, chain, address)
}

func GetProjectIcon(chain string, projectCode string) string {
	return fmt.Sprintf("%s/%s/%s.png", ImageCDN, chain, projectCode)
}
