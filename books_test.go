package main


import (
	"fmt"
	"testing"
)


var bookStrings = []struct{
	s string
	b Book
	e error
} {
	{
		"Стихотворения (1908-1937) (Осип Мандельштам)",
		Book{"Осип Мандельштам", "Стихотворения (1908-1937)"},
		nil,
	},
	{
		"Learning PHP, MySQL, JavaScript, CSS & HTML5 (Robin Nixon)",
		Book{"Robin Nixon", "Learning PHP, MySQL, JavaScript, CSS & HTML5"},
		nil,
	},
	{
		"Getting Things Done: How to achieve stress-free productivity (Allen, David)",
		Book{"Allen, David", "Getting Things Done: How to achieve stress-free productivity"},
		nil,
	},
	{
		"FDJKLSJDFJJJFDSJSLKDJFKLSJ ||||| sdjlkf jlsdkjf lkj((((",
		Book{},
		fmt.Errorf("Unknown string format: %s", "FDJKLSJDFJJJFDSJSLKDJFKLSJ ||||| sdjlkf jlsdkjf lkj(((("),
	},	
}


var books = []Book{
	{"Iain Banks", "Excession"},
	{"Iain Banks", "State of the art"},
	{"Bertrand Russell", "A History of Western Philosophy"},
	{"Robert Martin", "Clean Code"},			
}


func TestBookStorage(t *testing.T) {

	bs := NewBookStorage()

	fmt.Println(bs)

	bs.Add(&books[0])
	checkLen(&bs, t, 1)

	bs.AddIfMissing(&books[0])
	checkLen(&bs, t, 1)

	ob := Book{"Iain Banks", "Excession"}
	bs.Remove(&ob)
	checkLen(&bs, t, 0)

	for _, book := range books {
		bs.Add(&book)
	}
	checkLen(&bs, t, len(books))
	

	books[0].Title = "Use of Weapons"

	fmt.Println(bs, books)	

}

func checkLen(bs *BookStorage, t *testing.T, expected int) {
	if len := bs.Len(); len != expected {
		t.Errorf("Wrong Len of BookStorage. Actual: %d. Expected: %", len, expected)
	}
}


func TestBook(t *testing.T) {
	for _, d := range bookStrings {
		res, err := CreateBook(d.s)
		if res != d.b {
			t.Errorf("Wrong book created! String: %s | Actual: %v | Expected: %v", d.s, res, d.b)
		}
		if err != d.e && err.Error() != d.e.Error() {

			fmt.Println("Fucking error: ", len(err.Error()), err)
			fmt.Println("Fucking error: ", len(d.e.Error()), d.e)			
			t.Errorf("Wrong error message! String: %s | Actual: %v | Expected: %v", d.s, err, d.e)
		}			
	}
}
