package services
import "go-redis-k6-fiber/repositories"


type catalogService struct {
	productRepo repositories.ProductRepository
}

func NewCatalogService(productRepo repositories.ProductRepository) CatalogService {
	return catalogService{productRepo}
}

func (s catalogService) GetProduct() (products []Product,err error){
	productsDB,err := s.productRepo.GetProduct()
	if err != nil {
		return nil,err
	}
	for _,p := range productsDB{
		products = append(products, Product{ID: int(p.ID),Name: p.Name,Quantity: p.Quantity})
	}
	return products,nil
}