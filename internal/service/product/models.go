package product

var allProducts = []Product{ // TODO мок данных - надо сделать изменение данных - команды добавления, удаления
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Product struct {
	Title string
}
