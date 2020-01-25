package Models

type Configuration struct {
	MongoServer struct {
		URI string `json:"URI,omitempty"`
		DataBase string `json:"DataBase,omitempty"`
		InventoryCollection string `json:"InventoryCollection,omitempty"`
		ProviderCollection string `json:"ProviderCollection,omitempty"`
		OrderCollection string `json:"OrderCollection,omitempty"`
	}
	Server struct {
		Port string `json:"Port,omitempty"`
		Prefix string `json:"Prefix,omitempty"`
	}
}
