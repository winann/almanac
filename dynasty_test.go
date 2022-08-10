package almanac

import "testing"

func TestNewDynastyInfo(t *testing.T) {
	var year = 9999
	var d = NewDynastyInfo(year)
	var e = Dynasty{1949, 9999, 1948, "当代", "中国", "", "公历纪元"}

	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 1948
	d = NewDynastyInfo(year)
	e = Dynasty{1912, 37, 0, "近、现代", "中华民国", "", "民国"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 1949
	d = NewDynastyInfo(year)
	e = Dynasty{1949, 9999, 1948, "当代", "中国", "", "公历纪元"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 1582
	d = NewDynastyInfo(year)
	e = Dynasty{1573, 48, 0, "明", "神宗", "朱翊钧", "万历"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 888
	d = NewDynastyInfo(year)
	e = Dynasty{888, 1, 0, "唐", "僖宗", "李儇", "文德"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 777
	d = NewDynastyInfo(year)
	e = Dynasty{766, 14, 0, "唐", "肃宗", "李亨", "大历"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 60
	d = NewDynastyInfo(year)
	e = Dynasty{58, 18, 0, "东汉", "明帝", "刘庄", "永平"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = 0
	d = NewDynastyInfo(year)
	e = Dynasty{-1, 2, 0, "西汉", "哀帝", "刘欣", "元寿"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = -1
	d = NewDynastyInfo(year)
	e = Dynasty{-1, 2, 0, "西汉", "哀帝", "刘欣", "元寿"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}

	year = -1000
	d = NewDynastyInfo(year)
	e = Dynasty{-1019, 25, 0, "西周", "康王", "姬钊", "康王"}
	if e != *d {
		t.Error("NewDynastyInfo ", year, "Error:", d, ",But Expect:", e)
	}
}
