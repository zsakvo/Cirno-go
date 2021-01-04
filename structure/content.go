package structure

type ContentStruct struct {
	Code int64       `json:"code"`
	Data ContentData `json:"data"`
}

type ContentData struct {
	ChapterInfo ChapterInfo `json:"chapter_info"`
}

type ChapterInfo struct {
	ChapterID         string `json:"chapter_id"`
	BookID            string `json:"book_id"`
	DivisionID        string `json:"division_id"`
	UnitHlb           int64  `json:"unit_hlb"`
	ChapterIndex      string `json:"chapter_index"`
	ChapterTitle      string `json:"chapter_title"`
	AuthorSay         string `json:"author_say"`
	WordCount         string `json:"word_count"`
	Discount          string `json:"discount"`
	IsPaid            string `json:"is_paid"`
	AuthAccess        int64  `json:"auth_access"`
	BuyAmount         string `json:"buy_amount"`
	TsukkomiAmount    string `json:"tsukkomi_amount"`
	TotalHlb          string `json:"total_hlb"`
	Uptime            string `json:"uptime"`
	Mtime             string `json:"mtime"`
	RecommendBookInfo string `json:"recommend_book_info"`
	TxtContent        string `json:"txt_content"`
}
