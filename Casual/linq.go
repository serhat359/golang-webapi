package main

func selectChapters(chapters []LhReadChapter) []string {
	results := make([]string, len(chapters))
	for i, x := range chapters {
		results[i] = x.MangaChapter
	}

	return results
}