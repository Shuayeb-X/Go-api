package database

var Productlist []Product

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imgUrl"`
}

func init() {
	prd1 := Product{
		Id:          1,
		Title:       "Dragon",
		Description: "Dragon is red",
		Price:       "200 tk",
		ImgUrl:      "https://www.tastingtable.com/img/gallery/how-to-eat-dragon-fruit-for-the-uninitiated/intro-1682966430.jpg",
	}

	Productlist = append(Productlist, prd1)
}
