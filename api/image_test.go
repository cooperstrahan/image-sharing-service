package testing

import (
	"reflect"
	"testing"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/models"
)

func TestGetAllImages(t *testing.T) {
	t.Run("returns all images", func(t *testing.T) {
		got := models.GetAllImages()
		want := []models.Image{
			{
				ImageID:     2,
				FileName:    "dog2.jpg",
				Title:       "Right Hand Dog",
				Tags:        "Dog",
				Description: "Mans best freind",
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestGetImageById(t *testing.T) {
	t.Run("returns an image by its id", func(t *testing.T) {
		got, _ := models.GetImageById(1)
		want := models.Image{
			ImageID:     1,
			FileName:    "dog1.jpg",
			Title:       "Top Dog",
			Tags:        "Dog",
			Description: "The best dog you could ever get",
		}

		if got != &want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an image by its id", func(t *testing.T) {
		got, _ := models.GetImageById(2)
		want := models.Image{
			ImageID:     2,
			FileName:    "dog2.jpg",
			Title:       "Right Hand Dog",
			Tags:        "Dog",
			Description: "Mans best freind",
		}

		if got != &want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}

func TestUploadImage(t *testing.T) {
	t.Run("test upload functiton", func(t *testing.T) {
		image := models.Image{
			ImageID:     4,
			FileName:    "dog4.jpg",
			Title:       "Two good boys",
			Tags:        "Dog",
			Description: "A couple of good boys",
		}

		image.UploadImage()

		got, _ := models.GetImageById(image.ImageID)
		want := models.Image{
			ImageID:     4,
			FileName:    "dog4.jpg",
			Title:       "Two good boys",
			Tags:        "Dog",
			Description: "A couple of good boys",
		}

		if got != &want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func TestDeleteImage(t *testing.T) {
	t.Run("test delete function", func(t *testing.T) {
		models.DeleteImage(4)

		got, _ := models.GetImageById(4)
		want := models.Image{}

		if got != &want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
