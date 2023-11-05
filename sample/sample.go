package sample

import (
	"github.com/mongo-tut/internal/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Products = []database.Product{
	{
		ID:          primitive.NewObjectID(),
		ProductId:   primitive.NewObjectID(),
		Name:        "Product 1",
		Description: "This is product 1",
		Weight:      "1kg",
		Ratings:     5,
		Warranty:    2,
		Price:       3000,
	},
	{
		ID: primitive.NewObjectID(),

		ProductId:   primitive.NewObjectID(),
		Name:        "Product 2",
		Description: "This is product 2",
		Weight:      "2kg",
		Ratings:     4,
		Warranty:    1,
		Price:       6000,
	},
	{
		ID: primitive.NewObjectID(),

		ProductId:   primitive.NewObjectID(),
		Name:        "Product 3",
		Description: "This is product 3",
		Weight:      "5kg",
		Ratings:     4,
		Warranty:    0,
		Price:       8000,
	},
	{
		ID: primitive.NewObjectID(),

		ProductId:   primitive.NewObjectID(),
		Name:        "Product 4",
		Description: "This is product 4",
		Weight:      "8kg",
		Ratings:     4,
		Warranty:    6,
		Price:       9000,
	},
}

var Shops = []database.Shop{
	{
		ID:   primitive.NewObjectID(),
		Name: "Shop 1",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{12.9715987, 77.5945627},
		},
		ProductsInveontroy: []database.ProductInventory{
			{
				ProductID: Products[0].ProductId,
				Quantity:  70,
			},
			{
				ProductID: Products[1].ProductId,
				Quantity:  80,
			},
			{
				ProductID: Products[2].ProductId,
				Quantity:  800,
			},
		},
	},
	{
		ID:   primitive.NewObjectID(),
		Name: "Shop 2",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{28.6139391, 77.2090212},
		},
		ProductsInveontroy: []database.ProductInventory{
			{
				ProductID: Products[2].ProductId,
				Quantity:  700,
			},
			{
				ProductID: Products[3].ProductId,
				Quantity:  10,
			},
		},
	},
	{
		ID:   primitive.NewObjectID(),
		Name: "Shop 3",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{18.51957, 73.85535},
		},
		ProductsInveontroy: []database.ProductInventory{
			{
				ProductID: Products[0].ProductId,
				Quantity:  90,
			},
			{
				ProductID: Products[1].ProductId,
				Quantity:  80,
			},
			{
				ProductID: Products[3].ProductId,
				Quantity:  70,
			},
		},
	},
	{
		ID:   primitive.NewObjectID(),
		Name: "Shop 4",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{21.19594, 72.83023},
		},
		ProductsInveontroy: []database.ProductInventory{
			{
				ProductID: Products[0].ProductId,
				Quantity:  771,
			},
			{
				ProductID: Products[1].ProductId,
				Quantity:  113,
			},
		},
	},
}

var Users = []database.Users{
	{
		ShopID: Shops[0].ID,
		Name:   "Laukik",
		Email:  "lavquik@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-74.0060, 40.7128},
		},
	},
	{
		ShopID: Shops[3].ID,
		Name:   "shubham",
		Email:  "shubham@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-118.2437, 34.0522},
		},
	},
	{
		ShopID: Shops[0].ID,
		Name:   "maonoj",
		Email:  "maonoj@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-0.1278, 51.5074},
		},
	},
	{
		ShopID: Shops[1].ID,
		Name:   "Ram",
		Email:  "Ram@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{2.3522, 48.8566},
		},
	},
	{
		ShopID: Shops[1].ID,
		Name:   "Sunny",
		Email:  "Sunny@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{-80.1918, 25.7617},
		},
	},
	{
		ShopID: Shops[2].ID,
		Name:   "Riya",
		Email:  "Riya@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{37.6176, 55.7558},
		},
	},
	{
		ShopID: Shops[2].ID,
		Name:   "Saloni",
		Email:  "Saloni@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{37.6176, 55.7558},
		},
	},
	{
		ShopID: Shops[3].ID,
		Name:   "Manvi",
		Email:  "Manvi@gamil.com",
		Location: database.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{151.2093, -33.8688},
		},
	},
}
