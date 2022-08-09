package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// article represents data about a record article.
type article struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Authors string  `json:"authors"`
    PublishDate  string `json:"date"`
}

// articles slice to seed record article data.
var articles = []article{
    {
		ID: "1", 
		Title: "The Effect of Listening to Holy Quran Recitation on Anxiety: A Systematic Review",
		Authors: "Ashraf Ghiasi, Afsaneh Keramat",
		PublishDate: "2018 Nov-Dec",
	},
	{
		ID: "2", 
		Title: "Food addiction and obesity: unnecessary medicalization of hedonic overeating.",
		Authors: "Finlayson G",
		PublishDate: "2017 Aug",
	},
	{
		ID: "3", 
		Title: "Aging as disease.",
		Authors: "De Winter G.",
		PublishDate: "2015 May",
	},
	{
		ID: "4", 
		Title: "Emerging technologies and their potential for generating new assistive technologies",
		Authors: "Sarah Abdi, Irene Kitsara, Mark S Hawley, L P de Witte",
		PublishDate: "2021 Dec 1",
	},
	{
		ID: "5", 
		Title: "Overmedicalization in Children with Medical Complexity.",
		Authors: "Pordes E, Goodpasture M, Bordini BJ",
		PublishDate: "2020 Nov 1",
	},
	{
		ID: "6", 
		Title: "Test Article for ...",
		Authors: "Muhammad Rahman, Md Abd Ar Rahman",
		PublishDate: "2022 August 8",
	},
}

// CreateArticle godoc
// @Summary Creates a article
// @Description Create Article detail using POST request
// @Tags Articles
// @Param Body body article true "The body to create a article"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /article/create [post]
func CreateArticle(c *gin.Context) {
    var newArticle article
    if err := c.BindJSON(&newArticle); err != nil {
        return
    }
    articles = append(articles, newArticle)
    c.IndentedJSON(http.StatusCreated, newArticle)
}

// GetDetail godoc
// @Summary Fetch all article
// @Description Get all articles
// @Tags Articles
// @Accept  json
// @Param   id     path    string     true        "Article ID"
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /article/detail/{id} [GET]
func GetDetail(c *gin.Context) {
    id := c.Param("id")
    for _, a := range articles {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}

// ListArticles godoc
// @Summary Fetch all article
// @Description Get all articles
// @Tags Articles
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /article/list [GET]
func ListArticles(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, articles)
}
