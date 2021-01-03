package structure

type ChapterStruct struct {
	Code int64       `json:"code"`
	Data ChapterData `json:"data"`
}

type ChapterData struct {
	ChapterList     []ChapterList `json:"chapter_list"`
	MaxUpdateTime   int64         `json:"max_update_time"`
	MaxChapterIndex int64         `json:"max_chapter_index"`
}

type ChapterList struct {
	ChapterID      string `json:"chapter_id"`
	ChapterIndex   string `json:"chapter_index"`
	ChapterTitle   string `json:"chapter_title"`
	WordCount      string `json:"word_count"`
	TsukkomiAmount string `json:"tsukkomi_amount"`
	IsPaid         string `json:"is_paid"`
	Mtime          string `json:"mtime"`
	IsValid        string `json:"is_valid"`
	AuthAccess     int64  `json:"auth_access"`
}
