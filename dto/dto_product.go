package dto

type ProductDto struct {
	Id_Product   string  `json:"id_product"`
	Name_product string  `json:"name_product"`
	Cost         float64 `json:"cost"`
	Stock        int     `json:"stock"`
	Category     string  `json:"category"`

	/*Name string `json:"name"`
	UniversalCode string `json:"universal_code"`
	Picture string `json:"picture_url"`
	Price float32 `json:"base_price"`

	*/

}

type ProductsDto []ProductDto
