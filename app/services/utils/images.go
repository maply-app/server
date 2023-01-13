package utils

//import "strings"
//
//// image formats and magic numbers
//var magicTable = map[string]string{
//	"\xff\xd8\xff":      "image/jpeg",
//	"\x89PNG\r\n\x1a\n": "image/png",
//	"GIF87a":            "image/gif",
//	"GIF89a":            "image/gif",
//}
//
//// MimeFromIncipit returns the mime type of image file from its first few
//// bytes or the empty string if the file does not look like a known file type
//func MimeFromIncipit(i []byte) string {
//	//iStr := []byte(i)
//	for magic, mime := range magicTable {
//		if strings.HasPrefix(string(i), magic) {
//			return mime
//		}
//	}
//	return ""
//}
