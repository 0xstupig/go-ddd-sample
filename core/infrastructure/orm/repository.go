package orm

type Repository interface {
	GetAll(target interface{}, limit, offset int, preloads ...string) error
	GetBy(target interface{}, filters map[string]interface{}, limit, offset int, preloads ...string) (interface{}, error)
	GetOne(target interface{}, filters map[string]interface{}, preloads ...string) error

	Create(target interface{}) error
	Update(target interface{}) error
	Delete(target interface{}) error
}
