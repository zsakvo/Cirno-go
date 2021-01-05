package fetch

import (
	"os"

	"../structure"
)

type chapterStruct struct {
	Text string
	Cid  string
}

func splitArray(arr []structure.ChapterList, num int64) [][]structure.ChapterList {
	var segmens = make([][]structure.ChapterList, 0)
	max := int64(len(arr))
	if max < num {
		segmens = append(segmens, arr)
		return segmens
	}
	quantity := max / num
	end := int64(0)
	for i := int64(1); i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, arr[i-1+end:qu])
		} else {
			segmens = append(segmens, arr[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
