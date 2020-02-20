package shikimorigo

import "testing"

func TestGetAnimeID(t *testing.T) {
	for _, pair := range []struct {
		in  int
		out string
	}{
		{38000, "Kimetsu no Yaiba"},
		{19815, "No Game No Life"},
	} {
		v, err := GetAnimeID(pair.in)
		if err != nil {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"got", "Request error: "+err.Error(),
			)
		}
		if v.Name != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"got", v.Name,
			)
		}
	}
}

func TestGetMangaID(t *testing.T) {
	for _, pair := range []struct {
		in  int
		out string
	}{
		{33327, "Tokyo Ghoul"},
		{23390, "Shingeki no Kyojin"},
	} {
		v, err := GetMangaID(pair.in)
		if err != nil {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"got", "Request error: "+err.Error(),
			)
		}
		if v.Name != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"got", v.Name,
			)
		}
	}
}

func TestGetRanobeID(t *testing.T) {
	for _, pair := range []struct {
		in  int
		out string
	}{
		{12854, "Toaru Majutsu no Index"},
		{73631, "Haikyuu!! Shousetsu-ban!!"},
	} {
		v, err := GetRanobeID(pair.in)
		if err != nil {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"got", "Request error: "+err.Error(),
			)
		}
		if v.Name != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"got", v.Name,
			)
		}
	}
}
