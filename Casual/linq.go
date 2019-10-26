package main

func selectChapters(chapters []LhReadChapter) []string {
	results := make([]string, len(chapters))
	for i, x := range chapters {
		results[i] = x.MangaChapter
	}

	return results
}

func selectMangaNames(mangas []LhMangaChapterData) []string {
	results := make([]string, len(mangas))
	for i, x := range mangas {
		results[i] = x.MangaName
	}

	return results
}

func selectMangaChapters(mangas []LhMangaChapterData) []string {
	results := make([]string, len(mangas))
	for i, x := range mangas {
		results[i] = x.Chapter
	}

	return results
}
