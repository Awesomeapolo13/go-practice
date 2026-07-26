package verify

import (
	"go/validation-api/configs"
	"go/validation-api/pkg/files"
	"go/validation-api/pkg/request"
	"go/validation-api/pkg/response"
	"hash/fnv"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
)

type VerificationHandlerDeps struct {
	*configs.Config
}

type VerificationHandler struct {
	*configs.Config
}

func NewVerificationHandler(router *http.ServeMux, deps VerificationHandlerDeps) {
	handler := &VerificationHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *VerificationHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[SendVerificationRequest](&w, r)
		if err != nil {
			return
		}

		emailConfirmationConf := h.Config.EmailConfirmation

		hasher := fnv.New32a()
		hasher.Write([]byte(body.Email))
		hashNumber := hasher.Sum32()

		hash := strconv.FormatUint(uint64(hashNumber), 36)
		varification := NewVerification(body.Email, hash)
		link := "http://localhost:8081/verify/" + hash

		e := &email.Email{
			To:      []string{body.Email},
			From:    "Email verification <" + emailConfirmationConf.Email + ">",
			Subject: "Verify your email address",
			Text:    []byte("Click on the link to verify your e-mail address in your browser"),
			HTML:    []byte("<a href=\"" + link + "\">Click here</a>"),
		}
		err = e.Send("smtp.yandex.ru:465", smtp.PlainAuth("", emailConfirmationConf.Email, emailConfirmationConf.Password, "smtp.yandex.ru"))
		if err != nil {
			return
		}

		db := files.NewJsonDB("database.json")
		verifications, _ := files.GetCollection[Varification](db, "email_verification")
		verifications = append(verifications, *varification)
		err = files.SetCollection(db, "email_verification", verifications)
		if err != nil {
			return
		}

		res := SendVerificationResponse{
			Success: true,
		}
		// Тут будет отправка письма
		response.Json(w, res, http.StatusOK)
	}
}

func (h *VerificationHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := VerificationRequest{Hash: r.PathValue("hash")}
		if err := request.IsValid(payload); err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		db := files.NewJsonDB("database.json")
		verifications, _ := files.GetCollection[Varification](db, "email_verification")

		index := -1
		for i, v := range verifications {
			if v.Hash == payload.Hash {
				index = i
				break
			}
		}

		if index == -1 {
			res := VerificationResponse{
				Success: false,
				Message: "Could not find an email",
			}
			response.Json(w, res, http.StatusNotFound)
			return
		}

		verifications = append(verifications[:index], verifications[index+1:]...)
		if err := files.SetCollection(db, "email_verification", verifications); err != nil {
			return
		}

		res := VerificationResponse{
			Success: true,
			Message: "The email verified successfully.",
		}
		response.Json(w, res, http.StatusOK)
	}
}
