package controllers

import (
	"blog/app/models"
	"fmt"
	"strconv"

	"github.com/revel/revel"
)

/**
 * Add a blog for admin user
 */

// User for User Controller
type Post struct {
	Admin
}

func (p *Post) Index() revel.Result {
	return p.RenderTemplate("Admin/Post/Index.html")
}

// NewPostHandler to Add new article.
func (a *Post) NewPostHandler() revel.Result {
	data := new(PostData)
	a.Params.Bind(&data, "data")
	fmt.Println(data)
	a.Validation.Required(data.Title).Message("title can't be null.")
	a.Validation.Required(data.Content).Message("context can't be null.")
	a.Validation.Required(data.Date).Message("date can't be null.")

	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		// TODO Redirect new post page.
	}

	blog := new(models.Blogger)
	blog.Title = data.Title
	blog.Content = data.Content
	blog.CreateTime = data.Date

	uid := a.Session["UID"]
	id, _ := strconv.Atoi(uid)

	blog.CreateBy = id

	if data.passwd != "" {
		blog.Passwd = data.passwd
	}

	has, err := blog.New()

	if err != nil || has <= 0 {
		a.Flash.Error("msg", "create new blogger post error.")
		// TODO Redirect new post page.
	}

	return a.RenderHtml("ok")
}
