package testing

import (
	"testing"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/models"
)

// " '6', 'dog1.jpg', 'Doggo', 'dog', 'The happiest dog I have ever seen', '2021-07-27 23:13:04', '2021-07-27 23:13:04', NULL "
// " '7', 'dog2.jpg', 'Stoked dog', 'dogs', 'The most stoked dog bouncing around!', '2021-07-27 23:14:46', '2021-07-27 23:14:46', NULL"
// " '8', 'dog4.jpg', 'Two dogs', 'dogs', 'A couple of buddies palling around', '2021-07-27 23:15:52', '2021-07-27 23:15:52', NULL"
// " '9', 'cat1.jpg', 'Cute Cat', 'cats', 'A very cute grey cat.', '2021-07-27 23:16:28', '2021-07-27 23:16:28', NULL"
// " '10', 'dog3.jpg', 'Sweet puppy', 'dogs', 'A very sweet pup!', '2021-07-27 23:17:22', '2021-07-27 23:17:22', NULL "
// " '12', 'prancer.jpg', 'Prancer', 'dogs', 'Prancer, Coco\'s dog', '2021-07-28 02:31:14', '2021-07-28 02:31:14', NULL "
// " '13', 'IMG_8645.jpeg', 'Mug', 'this is a mug', 'Description', '2021-07-28 02:36:00', '2021-07-28 02:36:00', NULL "

func TestGetAllImages(t *testing.T) {
	t.Run("returns all images", func(t *testing.T) {
		got := models.GetAllImages()
		want := []models.Image{
			{
				ImageID:     6,
				FileName:    "dog1.jpg",
				Title:       "Doggo",
				Tags:        "dog",
				Description: "The happiest dog I have ever seen",
			},
			{
				ImageID:     7,
				FileName:    "dog2.jpg",
				Title:       "Stoked dog",
				Tags:        "dogs",
				Description: "The most stoked dog bouncing around!",
			},
			{
				ImageID:     8,
				FileName:    "dog4.jpg",
				Title:       "Two dogs",
				Tags:        "dogs",
				Description: "A couple of buddies palling around",
			},
			{
				ImageID:     9,
				FileName:    "cat1.jpg",
				Title:       "Cute Cat",
				Tags:        "cats",
				Description: "A very cute grey cat.",
			},

			{
				ImageID:     10,
				FileName:    "dog3.jpg",
				Title:       "Sweet puppy",
				Tags:        "dogs",
				Description: "A very sweet pup!",
			},
			{
				ImageID:     12,
				FileName:    "prancer.jpg",
				Title:       "Prancer",
				Tags:        "dogs",
				Description: "Prancer, Coco's dog",
			},
			{
				ImageID:     13,
				FileName:    "IMG_8645.jpeg",
				Title:       "Mug",
				Tags:        "this is a mug",
				Description: "Description",
			},
		}

		for i, s := range got {
			checkImageEquality(t, &s, &want[i])
		}

	})
}

func TestGetImageById(t *testing.T) {
	t.Run("returns an image by its id", func(t *testing.T) {
		got, _ := models.GetImageById(13)
		want := models.Image{
			ImageID:     13,
			FileName:    "IMG_8645.jpeg",
			Title:       "Mug",
			Tags:        "this is a mug",
			Description: "Description",
		}

		checkImageEquality(t, got, &want)

	})
}

func TestUploadImage(t *testing.T) {
	t.Run("test upload functiton", func(t *testing.T) {
		image := models.Image{
			FileName:    "dog4.jpg",
			Title:       "Two good boys",
			Tags:        "Dog",
			Description: "A couple of good boys",
		}

		image.UploadImage()

		newId := models.GetRecentId()

		got, _ := models.GetImageById(newId)
		want := models.Image{
			ImageID:     newId,
			FileName:    "dog4.jpg",
			Title:       "Two good boys",
			Tags:        "Dog",
			Description: "A couple of good boys",
		}

		checkImageEquality(t, got, &want)
	})

}

func TestDeleteImage(t *testing.T) {
	t.Run("test delete function", func(t *testing.T) {
		newId := models.GetRecentId()
		models.DeleteImage(newId)

		got, _ := models.GetImageById(newId)
		want := models.Image{}

		checkImageEquality(t, got, &want)

	})
}

func checkImageEquality(t testing.TB, a, b *models.Image) {
	t.Helper()
	if a.ImageID != b.ImageID {
		t.Errorf("Did not get equivalent Image IDs, got %d wanted %d", a.ImageID, b.ImageID)
	}
	if a.FileName != b.FileName {
		t.Errorf("Did not get equivalent FileNames, got %s wanted %s", a.FileName, b.FileName)
	}
	if a.Title != b.Title {
		t.Errorf("Did not get equivalent Titles, got %s wanted %s", a.Title, b.Title)
	}
	if a.Tags != b.Tags {
		t.Errorf("Did not get equivalent Tags, got %s wanted %s", a.Tags, b.Tags)
	}
	if a.Description != b.Description {
		t.Errorf("Did not get equivalent Descriptions, got %s wanted %s", a.Description, b.Description)
	}
}
