package stock

type Symbol struct {
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	DisplaySymbol  string `json:"displaySymbol"`
	Figi           string `json:"figi"`
	Isin           string `json:"isin,omitempty"`
	ShareClassFigi string `json:"shareClassFIGI"`
	Symbol         string `json:"symbol"`
	Symbol2        string `json:"symbol2"`
	Type           string `json:"type"`
}
