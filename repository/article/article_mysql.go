package article

import(
	"context"
	"database/sql"

	models "github.com/YuukiRen/GO_Api/models"
	artRepo "github.com/YuukiRen/GO_Api/repository"
)

type mysqlArticleRepo struct{
	Conn *sql.DB
}

func NewSQLArticleRepo(Conn *sql.DB) artRepo.ArticleRepo{
	return &mysqlArticleRepo{
		Conn:Conn,
	}
}

func (m *mysqlArticleRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Article,error){
	rows,err := m.Conn.QueryContext(ctx,query,args...)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	payload := make([]*models.Article,0)
	for rows.Next(){
		data := new(models.Article)
		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Description,
		)
		if err!=nil{
			return nil,err
		}
		payload = append(payload,data)
	}
	return payload,nil
}

func (m *mysqlArticleRepo) Fetch(ctx context.Context, num int64) ([]*models.Article,error){
	query := "SELECT id,title,description FROM articles limit ?"
	return m.fetch(ctx,query,num)
}

func (m *mysqlArticleRepo) GetByID(ctx context.Context, id int64) (*models.Article, error){
	query := "SELECT id,title,description FROM articles WHERE id=?"

	rows,err := m.fetch(ctx,query,id)
	if err!=nil{
		return nil,err
	}

	payload := &models.Article{}
	if len(rows)>0{
		payload = rows[0]
	} else{
		return nil,models.ErrNotFound
	}
	return payload,nil
}

func (m *mysqlArticleRepo) Create(ctx context.Context, p *models.Article) (int64, error){
	query:="INSERT articles SET title=?, description=?"
	stmt,err := m.Conn.PrepareContext(ctx,query)
	if err != nil{
		return -1,err
	}

	res,err :=stmt.ExecContext(ctx,p.Title,p.Description)
	if err != nil{
		return -1,err
	}
	return res.LastInsertId()
}

func (m *mysqlArticleRepo) Update(ctx context.Context,p *models.Article) (*models.Article,error){
	query:="UPDATE articles SET title=?, description=? WHERE id=?"
	stmt,err := m.Conn.PrepareContext(ctx,query)
	if err!=nil{
		return nil,err
	}

	_,err:= stmt.ExecContext(ctx,p.Title,p.Description,p.ID)
	if err!=nil{
		return nil,err
	}
	defer stmt.Close()
	return p,nil
}

func (m *mysqlArticleRepo) Delete(ctx context.Context,id int64) (bool, error){
	query := "DELETE FROM articles WHERE id = ?"
	stmt, err := m.Conn.PrepareContext(ctx,query)
	if err!= nil{
		return false, err
	}

	_,err = m.Conn.ExecContext(ctx,id)
	if err != nil{
		return false,err
	}
	return true,nil
}

