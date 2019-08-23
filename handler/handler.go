package handler

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/YuukiRen/GO_Api/driver"
	models "github.com/YuukiRen/GO_Api/models"
	repository "github.com/YuukiRen/GO_Api/repository"
	article "github.com/YuukiRen/GO_Api/repository/article"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Article{
	return &Article{
		repo: article.NewSQLArticleRepo(db.SQL),
	}
}

// Post ...
type Post struct{
	repo repository.ArticleRepo
}

// Fetch all data
func (p *Article) Fetch(w http.ResponseWriter, r *http.Request){
	payload, _ := p.repo.Fetch(r.Context(),5)

}

// Create a new post
func (p *Article) Create(w http.ResponseWriter, r *http.Request){
	article := models.Article{}
	json.NewDecoder(r.Body).Decode(&article)

	id,err := p.repo.Create(r.Context(), &article)
	fmt.Println(id)
	if err != nil{
		respondWithError(w,http.StatusInternalServerError,"Server Error")
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{"message":"successfully create new article"})
}

// Update a post by id
func (p *Article) Update(w http.ResponseWriter, r *http.Request){
	id,_:=strconv.Atoi(chi.URLParam(r,"id"))\
	data := models.Article{ID:int64(id)}

	json.NewDecoder(r.Body).Decoder(&data)
	payload, err:= p.repo.Update(r.Context(),&data)

	if err != nil{
		respondWithError(w, http.StatusInternalServerError,"Server Error")
	}
	respondWithJSON(w, http.StatusOK, payload)
}

// get data by id
func (p *Article) GetByID(w http.ResponseWriter, r *http.Request){
	id,_ :=strconv.Atoi(chi.URLParam(r,"id"))
	payload,err:=p.repo.GetByID(r.Context(),int64(id))
	if err!=nil{
		respondWithError(w,http.StatusNoContent ,"Content not found")
	}
	respondWithJSON(w, http.StatusOK,payload)
}

// Delete a post
func (p *Article) Delete(w http.ResponseWriter, r *http.Request){
	id,_ := strconv.Atoi(chi.URLParam(r,"id"))
	_, err := p.repo.Delete(r.Context(),int64(id))

	if err!=nil{
		respondWithError(w, http.StatusInternalServerError,"Server Error")
	}
}

