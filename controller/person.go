package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Asliddin3/tz/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PersonController struct {
	*Handler
}

func (h *Handler) NewPersonController(api *gin.RouterGroup) {
	person := &PersonController{h}
	br := api.Group("person")
	{
		br.POST("/", person.CreatePerson)
		br.PUT("/:id", person.UpdatePerson)
		br.GET("/:id", person.GetByID)
		br.DELETE("/:id", person.DeletePerson)
		br.GET("", person.GetPerson)
	}
}

// @Summary		  Create new person
// @Description	   this api is create new person
// @Tags			Person
// @Accept			json
// @Produce			json
// @Param			data    body		models.PersonRequest	false	"data body"
// @Success			201		{object}	models.Person
// @Failure			400,409	{object}	Response
// @Failure			500		{object}	Response
// @Router			/api/person/ [POST]
func (h *PersonController) CreatePerson(c *gin.Context) {
	var body models.PersonRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		h.log.Error("failed to create person", err.Error())
		return
	}
	pubData, err := h.getPersonData(body.Name)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "failed to get person data")
		h.log.Errorf("failed to get person data %s", err)
		return
	}
	person := models.Person{
		Name:       body.Name,
		Surname:    body.Surname,
		Patronymic: body.Patronymic,
		Age:        pubData.Age,
		Gender:     pubData.Gender,
		Nation:     pubData.Nation,
	}
	err = h.db.Debug().Clauses(clause.Returning{}).Create(&person).Error
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Errorf("failed to create person %v", err)
		return
	}
	c.JSON(http.StatusOK, person)
}

// @Summary		  Update person
// @Description	   this api is Update person
// @Tags			Person
// @Accept			json
// @Produce			json
// @Param           id    path     string   true   "update id"
// @Param			data 	body		models.PersonUpdateRequest	true	"data body"
// @Success			201		{object}	models.Person
// @Failure			400,409	{object}	Response
// @Failure			500		{object}	Response
// @Router			/api/person/{id} [PUT]
func (h *PersonController) UpdatePerson(c *gin.Context) {
	inputId := c.Param("id")
	var body models.PersonUpdateRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		h.log.Error("failed to create person", err.Error())
		return
	}
	columns := map[string]interface{}{}

	id, err := strconv.ParseInt(inputId, 10, 64)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	person := models.Person{
		ID: int(id),
	}
	if body.Name != "" {
		columns["name"] = body.Name
	}
	if body.Surname != "" {
		columns["surname"] = body.Surname
	}
	if body.Patronymic != "" {
		columns["patronymic"] = body.Patronymic
	}
	if body.Gender != "" {
		columns["gender"] = body.Gender
	}
	if body.Age != 0 {
		columns["age"] = body.Age
	}
	if body.Nation != "" {
		columns["nation"] = body.Nation
	}
	err = h.db.Debug().Clauses(clause.Returning{}).Model(&person).
		Where("id=?", inputId).Updates(columns).Error
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed to save person", err.Error())
		return
	}
	c.JSON(http.StatusOK, person)
}

// @Summary		    Get person
// @Description	    this api is to get person by filter
// @Tags			Person
// @Accept			json
// @Produce			json
// @Param           data  query   		models.PersonFilter true "person filter"
// @Success			201		{object}	[]models.Person
// @Failure			400,409	{object}	Response
// @Failure			500		{object}	Response
// @Router			/api/person  [GET]
func (h *PersonController) GetPerson(c *gin.Context) {
	var filter models.PersonFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var persons []models.Person
	db := h.db.Debug().Model(&models.Person{})
	if filter.MultiSearch != "" {
		field := fmt.Sprintf("%%%s%%", filter.MultiSearch)
		db = db.Where(`LOWER(name) LIKE LOWER(?) OR LOWER(surname) LIKE LOWER(?)
		 OR LOWER(patronymic) LIKE LOWER(?)`, field, field, field)
	}
	if filter.Name != "" {
		db = db.Where("LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", filter.Name))
	}
	if filter.Age != 0 {
		db = db.Where("age=?", 0)
	}
	if filter.Nation != "" {
		db = db.Where("nation=?", filter.Nation)
	}
	if filter.Gender != "" {
		db = db.Where("gender=?", filter.Gender)
	}
	var count int
	err = db.Select("COUNT(*)").Scan(&count).Error
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed to get count of person", err.Error())
		return
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.PageSize == 0 {
		filter.PageSize = 10
	}
	err = db.Select("*").Find(&persons).Error
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.PersonFilterResponse{
		Person:   persons,
		Page:     filter.Page,
		PageSize: filter.PageSize,
		Count:    count,
	})
}

// @Summary		  Get person by id
// @Description	   this api is to get person by id
// @Tags			Person
// @Accept			json
// @Produce			json
// @Param           id    path     string   true   "id"
// @Success			201		{object}	models.Person
// @Failure			400,409	{object}	Response
// @Failure			500		{object}	Response
// @Router			/api/person/{id} [GET]
func (h *PersonController) GetByID(c *gin.Context) {
	personId := c.Param("id")
	var person models.Person
	err := h.db.Debug().First(&person, "id=?", personId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newResponse(c, http.StatusBadRequest, "not found person")
			return
		}
		newResponse(c, http.StatusBadRequest, err.Error())
		h.log.Error("failed to get person by id", err.Error())
		return
	}
	c.JSON(http.StatusOK, person)
}

// @Summary		  Delete person
// @Description	   this api is to delete person
// @Tags			Person
// @Accept			json
// @Produce			json
// @Param           id    path     int   true   "id"
// @Success			201		{object}	Response
// @Failure			400,409	{object}	Response
// @Failure			500		{object}	Response
// @Router			/api/person/{id} [DELETE]
func (h *PersonController) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	err := h.db.Delete(&models.Person{}, "id=?", id).Error
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		h.log.Error("failed to delete person", err.Error())
		return
	}
	newResponse(c, http.StatusOK, "success")
}

type PersonData struct {
	Age    int    `json:"age"`
	Nation string `json:"nation"`
	Gender string `json:"gender"`
}

func (h *Handler) getPersonData(name string) (*PersonData, error) {
	urls := []string{"https://api.agify.io/?name=", "https://api.genderize.io/?name=",
		"https://api.nationalize.io/?name="}
	person := PersonData{}
	for _, url := range urls {
		response, err := http.Get(url + name)
		if err != nil {
			h.log.Errorf("failed to get person from public api %v", err)
			return nil, err
		}
		defer response.Body.Close()

		res, err := io.ReadAll(response.Body)
		if err != nil {
			h.log.Errorf("Error reading response body:%v", err)
			return nil, err
		}
		data := map[string]interface{}{}
		err = json.Unmarshal(res, &data)
		if err != nil {
			h.log.Errorf("Error unmarshaling JSON:%v", err)
			return nil, err
		}
		if val, ok := data["age"]; ok {
			person.Age = int(val.(float64))
		} else if val, ok := data["gender"]; ok {
			person.Gender = val.(string)
		} else if val, ok := data["country"]; ok {
			arr, ok := val.([]CountriesRes)
			if len(arr) > 0 && ok {
				person.Nation = arr[0].CountryID
			}
		}
	}
	return &person, nil
}

type CountriesRes struct {
	CountryID string `json:"country_id"`
}
