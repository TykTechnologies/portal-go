package tdp

type Products struct{}

func (u Products) CreateProduct(input ProductCreateInput) {}

func (u Products) GetProduct(id uint64) {}

func (u Products) ListProducts(options ProductListOptions) {}

func (u Products) UpdateProducts(id uint64, input ProductUpdateInput) {}

type ProductUpdateInput struct {
	Catalogues []uint64
}

type ProductCreateInput struct{}

type ProductListOptions struct{}
