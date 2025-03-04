// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"
)

func newPostTag(db *gorm.DB, opts ...gen.DOOption) postTag {
	_postTag := postTag{}

	_postTag.postTagDo.UseDB(db, opts...)
	_postTag.postTagDo.UseModel(&model.PostTag{})

	tableName := _postTag.postTagDo.TableName()
	_postTag.ALL = field.NewAsterisk(tableName)
	_postTag.ID = field.NewUint(tableName, "id")
	_postTag.CreatedAt = field.NewTime(tableName, "created_at")
	_postTag.UpdatedAt = field.NewTime(tableName, "updated_at")
	_postTag.DeletedAt = field.NewField(tableName, "deleted_at")
	_postTag.PostID = field.NewInt64(tableName, "post_id")
	_postTag.TagID = field.NewInt64(tableName, "tag_id")

	_postTag.fillFieldMap()

	return _postTag
}

type postTag struct {
	postTagDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	PostID    field.Int64
	TagID     field.Int64

	fieldMap map[string]field.Expr
}

func (p postTag) Table(newTableName string) *postTag {
	p.postTagDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p postTag) As(alias string) *postTag {
	p.postTagDo.DO = *(p.postTagDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *postTag) updateTableName(table string) *postTag {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.PostID = field.NewInt64(table, "post_id")
	p.TagID = field.NewInt64(table, "tag_id")

	p.fillFieldMap()

	return p
}

func (p *postTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *postTag) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["post_id"] = p.PostID
	p.fieldMap["tag_id"] = p.TagID
}

func (p postTag) clone(db *gorm.DB) postTag {
	p.postTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p postTag) replaceDB(db *gorm.DB) postTag {
	p.postTagDo.ReplaceDB(db)
	return p
}

type postTagDo struct{ gen.DO }

type IPostTagDo interface {
	gen.SubQuery
	Debug() IPostTagDo
	WithContext(ctx context.Context) IPostTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPostTagDo
	WriteDB() IPostTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPostTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPostTagDo
	Not(conds ...gen.Condition) IPostTagDo
	Or(conds ...gen.Condition) IPostTagDo
	Select(conds ...field.Expr) IPostTagDo
	Where(conds ...gen.Condition) IPostTagDo
	Order(conds ...field.Expr) IPostTagDo
	Distinct(cols ...field.Expr) IPostTagDo
	Omit(cols ...field.Expr) IPostTagDo
	Join(table schema.Tabler, on ...field.Expr) IPostTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPostTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPostTagDo
	Group(cols ...field.Expr) IPostTagDo
	Having(conds ...gen.Condition) IPostTagDo
	Limit(limit int) IPostTagDo
	Offset(offset int) IPostTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPostTagDo
	Unscoped() IPostTagDo
	Create(values ...*model.PostTag) error
	CreateInBatches(values []*model.PostTag, batchSize int) error
	Save(values ...*model.PostTag) error
	First() (*model.PostTag, error)
	Take() (*model.PostTag, error)
	Last() (*model.PostTag, error)
	Find() ([]*model.PostTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PostTag, err error)
	FindInBatches(result *[]*model.PostTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PostTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPostTagDo
	Assign(attrs ...field.AssignExpr) IPostTagDo
	Joins(fields ...field.RelationField) IPostTagDo
	Preload(fields ...field.RelationField) IPostTagDo
	FirstOrInit() (*model.PostTag, error)
	FirstOrCreate() (*model.PostTag, error)
	FindByPage(offset int, limit int) (result []*model.PostTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPostTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p postTagDo) Debug() IPostTagDo {
	return p.withDO(p.DO.Debug())
}

func (p postTagDo) WithContext(ctx context.Context) IPostTagDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postTagDo) ReadDB() IPostTagDo {
	return p.Clauses(dbresolver.Read)
}

func (p postTagDo) WriteDB() IPostTagDo {
	return p.Clauses(dbresolver.Write)
}

func (p postTagDo) Session(config *gorm.Session) IPostTagDo {
	return p.withDO(p.DO.Session(config))
}

func (p postTagDo) Clauses(conds ...clause.Expression) IPostTagDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postTagDo) Returning(value interface{}, columns ...string) IPostTagDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postTagDo) Not(conds ...gen.Condition) IPostTagDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postTagDo) Or(conds ...gen.Condition) IPostTagDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postTagDo) Select(conds ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postTagDo) Where(conds ...gen.Condition) IPostTagDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postTagDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPostTagDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p postTagDo) Order(conds ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postTagDo) Distinct(cols ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postTagDo) Omit(cols ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postTagDo) Join(table schema.Tabler, on ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postTagDo) RightJoin(table schema.Tabler, on ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postTagDo) Group(cols ...field.Expr) IPostTagDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postTagDo) Having(conds ...gen.Condition) IPostTagDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postTagDo) Limit(limit int) IPostTagDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postTagDo) Offset(offset int) IPostTagDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPostTagDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postTagDo) Unscoped() IPostTagDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postTagDo) Create(values ...*model.PostTag) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postTagDo) CreateInBatches(values []*model.PostTag, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postTagDo) Save(values ...*model.PostTag) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postTagDo) First() (*model.PostTag, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostTag), nil
	}
}

func (p postTagDo) Take() (*model.PostTag, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostTag), nil
	}
}

func (p postTagDo) Last() (*model.PostTag, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostTag), nil
	}
}

func (p postTagDo) Find() ([]*model.PostTag, error) {
	result, err := p.DO.Find()
	return result.([]*model.PostTag), err
}

func (p postTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PostTag, err error) {
	buf := make([]*model.PostTag, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postTagDo) FindInBatches(result *[]*model.PostTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postTagDo) Attrs(attrs ...field.AssignExpr) IPostTagDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postTagDo) Assign(attrs ...field.AssignExpr) IPostTagDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postTagDo) Joins(fields ...field.RelationField) IPostTagDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postTagDo) Preload(fields ...field.RelationField) IPostTagDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postTagDo) FirstOrInit() (*model.PostTag, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostTag), nil
	}
}

func (p postTagDo) FirstOrCreate() (*model.PostTag, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostTag), nil
	}
}

func (p postTagDo) FindByPage(offset int, limit int) (result []*model.PostTag, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postTagDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postTagDo) Delete(models ...*model.PostTag) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postTagDo) withDO(do gen.Dao) *postTagDo {
	p.DO = *do.(*gen.DO)
	return p
}
