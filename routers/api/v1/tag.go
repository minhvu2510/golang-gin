package v1

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/minhvu2510/golang-gin/pkg/app"
	"github.com/minhvu2510/golang-gin/pkg/e"
	"github.com/minhvu2510/golang-gin/pkg/setting"
	"github.com/minhvu2510/golang-gin/pkg/util"
	"github.com/minhvu2510/golang-gin/service/tag_service"
	"github.com/unknwon/com"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	fmt.Println("---tag service", tagService)
	tags, err := tagService.GetAll()
	fmt.Println("----tags", tags)
	fmt.Println("----err", tags)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	// count, err := tagService.Count()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	// 	return
	// }

	// appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	// 	"lists": tags,
	// 	"total": count,
	// })
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
	})
}

// format data for add tag
type AddTagForm struct {
	Name      string `form:"name" json:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" json:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" json:"state" valid:"Range(0,1)"`
}

//@Summary Add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}
	var addTag AddTagForm
	if err := c.ShouldBindJSON(&addTag); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tag := tag_service.Tag{
		Name:      addTag.Name,
		CreatedBy: addTag.CreatedBy,
		State:     addTag.State,
	}
	// cehck exist tag
	exists, err := tag.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	// add tag
	err_add := tag.Add()
	if err_add != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	// name := c.Query("name")
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// fomat data for edit tag
type EditTagForm struct {
	ID         int    `form:"id" json:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" json:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" json:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" json:"state" valid:"Range(0,1)"`
}

func EditTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID must > 1")

	if valid.HasErrors() {
		fmt.Println(valid.Errors[0].Message)
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
