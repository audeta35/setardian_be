package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"steradian_be/models"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SteradianHandler struct {
	DB *sql.DB
}

func InitSteradian(db *sql.DB) SteradianHandler {
	return SteradianHandler{
		DB: db,
	}
}

func (h SteradianHandler) FetchArticles(c echo.Context) (err error) {
	datas := make([]models.Article, 0)
	query := `SELECT * FROM article`

	rows, err := h.DB.Query(query)

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	defer rows.Close()

	for rows.Next() {
		var item models.Article
		err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Body,
		)

		if err != nil {
			resp := ErrorResponse{
				Message: err.Error(),
			}

			return c.JSON(http.StatusInternalServerError, resp)
		}

		datas = append(datas, item)
	}

	return c.JSON(http.StatusOK, datas)
}

func (h SteradianHandler) InsertArticles(c echo.Context) (err error) {
	var item models.Article
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `INSERT article SET title=?, body=?`

	dbRes, err := h.DB.Exec(query, item.Title, item.Body)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

func (h SteradianHandler) GetDetailArticle(c echo.Context) (err error) {
	articleID := c.Param("id")

	query := `SELECT * FROM article WHERE id=?`
	row := h.DB.QueryRow(query, articleID)
	var res models.Article
	err = row.Scan(
		&res.ID,
		&res.Title,
		&res.Body,
	)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h SteradianHandler) EditArticles(c echo.Context) (err error) {

	// articleId := c.Param("id")
	var item models.Article
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `UPDATE article SET title=?, body=? WHERE id=?`

	dbRes, err := h.DB.Exec(query, item.Title, item.Body, item.ID)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

func (h SteradianHandler) DeleteArticle(c echo.Context) (err error) {
	articleID := c.Param("id")

	query := `DELETE FROM article WHERE id=?`
	row := h.DB.QueryRow(query, articleID)
	var res models.Article
	err = row.Scan(
		&res.ID,
		&res.Title,
		&res.Body,
	)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, res)
}

// users CRUD

func (h SteradianHandler) LoginUser(c echo.Context) (err error) {
	var item models.UserAdminLogin

	query := `SELECT email, password FROM admin WHERE email=? AND password =?`
	row := h.DB.QueryRow(query, item.Email, item.Password)

	var res models.UserAdminData
	err = row.Scan(
		&res.UserID,
		&res.Email,
	)

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h SteradianHandler) RegisterUser(c echo.Context) (err error) {
	var item models.UserRegister
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `INSERT user SET email=?, password=?, phoneNumber=?, city=?, zip=?, message=?, username=?, address=?`

	dbRes, err := h.DB.Exec(query, item.Email, item.Password, item.PhoneNumber, item.City, item.Zip, item.Message, item.UserName, item.Address)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

// admin CRUD

func (h SteradianHandler) LoginAdmin(c echo.Context) (err error) {
	var item models.UserAdminLogin

	query := `SELECT email, password FROM admin WHERE email=? AND password =?`
	row := h.DB.QueryRow(query, item.Email, item.Password)

	var res models.UserAdminData
	err = row.Scan(
		&res.UserID,
		&res.Email,
	)

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h SteradianHandler) RegisterAdmin(c echo.Context) (err error) {
	var item models.AdminRegister
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `INSERT admin SET email=?, password=?`

	dbRes, err := h.DB.Exec(query, item.Email, item.Password)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

// order CRUD

func (h SteradianHandler) OrderFetch(c echo.Context) (err error) {
	datas := make([]models.Orders, 0)
	query := `SELECT * FROM orders`

	rows, err := h.DB.Query(query)

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	defer rows.Close()

	for rows.Next() {
		var res models.Orders
		err := rows.Scan(
			&res.PickUpLoc,
			&res.DropOffLoc,
			&res.PickUpDate,
			&res.DropOffDate,
			&res.PickUpTime,
			&res.CarId,
			&res.UserId,
			&res.AdminId,
			&res.ID,
		)

		if err != nil {
			resp := ErrorResponse{
				Message: err.Error(),
			}

			return c.JSON(http.StatusInternalServerError, resp)
		}

		datas = append(datas, res)
	}

	return c.JSON(http.StatusOK, datas)
}

func (h SteradianHandler) OrderAdd(c echo.Context) (err error) {
	var item models.Orders
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `INSERT order SET pickUpLoc=?, dropOffLoc=?, pickUpDate=?, dropOffDate=?, pickUpTime=?, carId=?, userId=?, adminId=?`

	dbRes, err := h.DB.Exec(query, item.PickUpLoc, item.DropOffLoc, item.PickUpDate, item.DropOffDate, item.PickUpTime, item.CarId, item.UserId, item.AdminId)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

func (h SteradianHandler) OrderEdit(c echo.Context) (err error) {

	// articleId := c.Param("id")
	var item models.Orders
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `UPDATE orders SET pickUpLoc=?, dropOffLoc=?, pickUpDate=?, dropOffDate=?, pickUpTime=?, carId=?, userId=?, adminId=? WHERE id=?`

	dbRes, err := h.DB.Exec(query, item.PickUpLoc, item.DropOffLoc, item.PickUpDate, item.DropOffDate, item.PickUpTime, item.CarId, item.UserId, item.AdminId, item.ID)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

func (h SteradianHandler) OrderDelete(c echo.Context) (err error) {
	orderID := c.Param("id")

	query := `DELETE FROM orders WHERE id=?`
	row := h.DB.QueryRow(query, orderID)
	var res models.Orders
	err = row.Scan(
		&res.PickUpLoc,
		&res.DropOffLoc,
		&res.PickUpDate,
		&res.DropOffDate,
		&res.PickUpTime,
		&res.CarId,
		&res.UserId,
		&res.AdminId,
		&res.ID,
	)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, res)
}

// cars CRUD

func (h SteradianHandler) CarsFetch(c echo.Context) (err error) {
	datas := make([]models.Cars, 0)
	query := `SELECT * FROM cars`

	rows, err := h.DB.Query(query)

	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	defer rows.Close()

	for rows.Next() {
		var res models.Cars
		err := rows.Scan(
			&res.CarId,
			&res.Name,
			&res.CarType,
			&res.Rating,
			&res.Fuel,
			&res.Image,
			&res.HourRate,
			&res.DayRate,
			&res.MonthRate,
		)

		if err != nil {
			resp := ErrorResponse{
				Message: err.Error(),
			}

			return c.JSON(http.StatusInternalServerError, resp)
		}

		datas = append(datas, res)
	}

	return c.JSON(http.StatusOK, datas)
}

func (h SteradianHandler) CarsAdd(c echo.Context) (err error) {
	var item models.Cars
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `INSERT cars SET name=?, carType=?, rating=?, fuel=?, image=?, hourRate=?, dayRate=?, monthRate=?`

	dbRes, err := h.DB.Exec(query,
		item.Name,
		item.CarType,
		item.Rating,
		item.Fuel,
		item.Image,
		item.HourRate,
		item.DayRate,
		item.MonthRate)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

func (h SteradianHandler) CarsEdit(c echo.Context) (err error) {

	// articleId := c.Param("id")
	var item models.Cars
	err = c.Bind(&item)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	query := `UPDATE orders SET name=?, carType=?, rating=?, fuel=?, image=?, hourRate=?, dayRate=?, monthRate=? WHERE id=?`

	dbRes, err := h.DB.Exec(query,
		item.Name,
		item.CarType,
		item.Rating,
		item.Fuel,
		item.Image,
		item.HourRate,
		item.DayRate,
		item.MonthRate)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	insertedID, err := dbRes.LastInsertId()
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	item.ID = fmt.Sprintf("%d", insertedID)
	return c.JSON(http.StatusCreated, item)
}

func (h SteradianHandler) CarsDelete(c echo.Context) (err error) {
	carsId := c.Param("id")

	query := `DELETE FROM cars WHERE id=?`
	row := h.DB.QueryRow(query, carsId)
	var res models.Cars
	err = row.Scan(
		&res.CarId,
		&res.Name,
		&res.CarType,
		&res.Rating,
		&res.Fuel,
		&res.Image,
		&res.HourRate,
		&res.DayRate,
		&res.MonthRate,
	)
	if err != nil {
		resp := ErrorResponse{
			Message: err.Error(),
		}
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, res)
}
