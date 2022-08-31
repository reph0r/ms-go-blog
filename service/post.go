package service

import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func SavePost(post *models.Post)  {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post)  {
	dao.UpdatePost(post)
}
func GetPostDetail(pid int) (*models.PostRes,error)  {
	post,err := dao.GetPostById(pid)
	if err != nil {
		return nil,err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes,nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}

func SearchPost(condition string) []models.SearchResp  {
	posts,_ := dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _,post := range posts{
		searchResps = append(searchResps,models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResps
}