package v1

import (
	"fmt"
	"net/http"

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
