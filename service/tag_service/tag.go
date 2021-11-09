package tag_service

import "fmt"

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) GetAll() (string, error) {
	fmt.Println(" ---- getall tags-----")
	// var (
	// 	tags, cacheTags []models.Tag
	// )

	// cache := cache_service.Tag{
	// 	State: t.State,

	// 	PageNum:  t.PageNum,
	// 	PageSize: t.PageSize,
	// }
	// key := cache.GetTagsKey()
	// if gredis.Exists(key) {
	// 	data, err := gredis.Get(key)
	// 	if err != nil {
	// 		logging.Info(err)
	// 	} else {
	// 		json.Unmarshal(data, &cacheTags)
	// 		return cacheTags, nil
	// 	}
	// }

	// tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	// if err != nil {
	// 	return nil, err
	// }

	// gredis.Set(key, tags, 3600)
	return t.Name, nil
}
