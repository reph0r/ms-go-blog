package views

import (
	"errors"
	"log"
	"ms-go-blog/common"
	"ms-go-blog/context"
	"ms-go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) CategoryNew(ctx *context.MsContext)  {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1  1参数 分类的id
	cIdStr := ctx.GetPathVariable("id")
	cId,_ := strconv.Atoi(cIdStr)
	pageStr,_ := ctx.GetForm("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page,_ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse,err := service.GetPostsByCategoryId(cId,page,pageSize);
	if err != nil {
		categoryTemplate.WriteError(ctx.W,err)
		return
	}
	categoryTemplate.WriteData(ctx.W,categoryResponse)
}

func (*HTMLApi) Category(w http.ResponseWriter,r *http.Request)  {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1  1参数 分类的id
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path,"/c/")
	cId,err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w,errors.New("不识别此请求路径"))
		return
	}
	if err := r.ParseForm();err != nil{
		log.Println("表单获取失败：",err)
		categoryTemplate.WriteError(w,errors.New("系统错误，请联系管理员!!"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page,_ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse,err := service.GetPostsByCategoryId(cId,page,pageSize);
	if err != nil {
		categoryTemplate.WriteError(w,err)
		return
	}
	categoryTemplate.WriteData(w,categoryResponse)
}
