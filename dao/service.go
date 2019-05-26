package dao

import (
	"github.com/arangodb/go-driver"
	"triton/models"
)

/*ServiceDAO - прослойка для работы с DAO сервиса*/
type ServiceDAO struct {
	database driver.Database
}

const NAME = "HISTORY"
const REQUEST = "FOR d in HISTORY LIMIT @offset, @limit return d"

/*CreateHistory - добавление новой истории*/
func (dao *ServiceDAO) CreateHistory(model models.History) (meta driver.DocumentMeta, err error) {
	col, err := dao.database.Collection(nil, NAME)
	meta, err = col.CreateDocument(nil, model)
	return
}

/*FindAllHistories - получить все даннные историй*/
func (dao *ServiceDAO) FindAllHistories(skip, limit int) (result []models.History, err error) {
	vars := map[string]interface{}{}
	vars["offset"] = skip
	vars["limit"] = limit
	cur, err := dao.database.Query(nil, REQUEST, vars)
	defer cur.Close()
	result = []models.History{}
	for {
		var doc models.History
		_, err := cur.ReadDocument(nil, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {

		}
		result = append(result, doc)
	}
	return
}
