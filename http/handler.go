package http

import (
	"fiber-mongo/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Handler struct {
	Svc domain.UserSvc
}

func NewHandler(svc domain.UserSvc) *Handler {
	return &Handler{
		Svc: svc,
	}
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("Aris is god!!")
	//ID := r.Context().Value("id").(string)
	//fmt.Println("pet id : ", ID)
	//petID, err := uuid.Parse(ID)
	//if err != nil {
	//	resp := Resp{
	//		Code: http.StatusBadRequest,
	//		Msg:  "invalid pet_id provided in url param",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//if petID == uuid.Nil {
	//	resp := Resp{
	//		Code: http.StatusBadRequest,
	//		Msg:  "please provide the pet id to retrieve",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//pet, err := h.Svc.Get(ID)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("error - fetching pet detail from db failed, err : %v", err))
	//	resp := Resp{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "fetching pet detail failed. please try again later",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//resp := Resp{
	//	Code: http.StatusOK,
	//	Msg:  "success",
	//	Data: pet,
	//}
	//respond(w, r, &resp)
}

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	return c.SendString("kist here")
	//query := r.Context().Value(QUERY).(*ListPetsQuery)
	//pets, err := h.Svc.List(query.Category)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("error - fetching pet details for given category from db failed, err : %v", err))
	//	resp := Resp{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "fetching all pets detail failed. please try again later",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//resp := Resp{
	//	Code: http.StatusOK,
	//	Msg:  "success",
	//	Data: pets,
	//}
	//respond(w, r, &resp)
}

func (h *Handler) AddUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	result, err := h.Svc.Create(&user)
	if err != nil {
		fmt.Println(fmt.Errorf("error - adding new user detail to db failed, err : %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": result.InsertedID})
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	objId, _ := primitive.ObjectIDFromHex(id)
	log.Println("objId", objId)
	err := h.Svc.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.SendString("deleted")
	//ID := r.Context().Value("id").(string)
	//fmt.Println("pet id : ", ID)
	//petID, err := uuid.Parse(ID)
	//if err != nil {
	//	resp := Resp{
	//		Code: http.StatusBadRequest,
	//		Msg:  "invalid pet_id provided in url param",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//if petID == uuid.Nil {
	//	resp := Resp{
	//		Code: http.StatusBadRequest,
	//		Msg:  "please provide the pet id to retrieve",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//err = h.Svc.Delete(petID)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("error - deleting pet record from db failed, err: %v", err))
	//	resp := Resp{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "unable to delete the pet details. please try again later",
	//	}
	//	respond(w, r, &resp)
	//	return
	//}
	//resp := Resp{
	//	Code: http.StatusOK,
	//	Msg:  "success",
	//}
	//respond(w, r, &resp)
}
