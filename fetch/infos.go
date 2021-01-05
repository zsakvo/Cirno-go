package fetch

// import (
// 	"../ciweimao"
// 	"../structure"
// 	"github.com/cheggaaa/pb"
// )

// func GetInfos(bid string) (string, [][]structure.ChapterList) {
// 	detail := ciweimao.GetDetail(bid)
// 	chapters := ciweimao.GetCatalog(bid)
// 	totalCount := len(chapters)
// 	txtBar = pb.StartNew(totalCount)
// 	txtContainer := make(map[string]string)
// 	errs := []structure.ChapterList{}
// 	txt := make(chan textStruct, 409600)
// 	err := make(chan structure.ChapterList, 102400)
// 	chaptersArr := splitArray(chapters, 16)
// 	return detail.BookName + " / " + detail.AuthorName, chaptersArr
// }
