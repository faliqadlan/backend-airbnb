package image

import (
	"be/delivery/controllers/templates"
	"be/entities"
	"be/repository/database/image"
	"be/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/gommon/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/labstack/echo/v4"
)

type ImageController struct {
	repo image.Image
}

func New(repo image.Image) *ImageController {
	return &ImageController{
		repo: repo,
	}
}

func (ic *ImageController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		//default image

		image := CreateImageRequesFormat{}
		image.Room_uid = c.FormValue("room_uid")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		log.Info(src)

		defer src.Close()

		s, _ := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKIAS4KA3W5H4Z73S3NR",                     // id
				"XVGjvN4ApOPqNFH95wfmpM06PpQfqiXdDhGuBcFp", // secret
				""),
		})

		fileName, _ := utils.UploadFileToS3(s, src, file)


		// log.Info(fileName)
		// user := UserCreateRequest{}
		// image := entities.Image{}
		image.Url = "https://test-upload-s3-rogerdev.s3.ap-southeast-1.amazonaws.com/" + fileName

		res, err := ic.repo.Create(entities.Image{Room_uid: image.Room_uid, Url: image.Url})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, "error internal server error fo create new image", err))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "Success create new image", res))
	}
}
