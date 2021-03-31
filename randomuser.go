package randomuser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const endpoint = "https://randomuser.me/api/1.3/?noinfo"

// RandomuserScheme is the all randomuser scheme response
type RandomuserScheme struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City        string `json:"city"`
			State       string `json:"state"`
			Country     string `json:"country"`
			Postcode    int    `json:"postcode"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
			Timezone struct {
				Offset      string `json:"offset"`
				Description string `json:"description"`
			} `json:"timezone"`
		} `json:"location"`
		Email string `json:"email"`
		Login struct {
			UUID     string `json:"uuid"`
			Username string `json:"username"`
			Password string `json:"password"`
			Salt     string `json:"salt"`
			Md5      string `json:"md5"`
			Sha1     string `json:"sha1"`
			Sha256   string `json:"sha256"`
		} `json:"login"`
		Dob struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"dob"`
		Registered struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"registered"`
		Phone string `json:"phone"`
		Cell  string `json:"cell"`
		ID    struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"id"`
		Picture struct {
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Thumbnail string `json:"thumbnail"`
		} `json:"picture"`
		Nat string `json:"nat"`
	} `json:"results"`
}

// NewRandomUserJson instance the API request
func NewUsers(results int) (*RandomuserScheme, error) {
	var r RandomuserScheme

	res, err := http.Get(fmt.Sprintf("%s&nat=es&results=%d", endpoint, results))
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error not 200 OK: %s", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		return nil, fmt.Errorf("JSON type error: %s", err)
	}

	return &r, nil
}

// FullName is a method return composite of Title, First and Last
func (r *RandomuserScheme) FullName(id int) string {
	name := &r.Results[id].Name

	return fmt.Sprintf("%s %s %s", name.Title, name.First, name.Last)
}

// Picture return path image according to large, medium, thumbnail
func (r *RandomuserScheme) Picture(id int, typeSize string) string {
	picture := &r.Results[id].Picture

	switch typeSize {
	case "large":
		return picture.Large
	case "medium":
		return picture.Medium
	case "thumbnail":
		return picture.Thumbnail
	default:
		return picture.Medium
	}
}

// Gender return gender by id
func (r *RandomuserScheme) Gender(id int) string {
	return r.Results[id].Gender
}

// Age return age by id
func (r *RandomuserScheme) Age(id int) int {
	return r.Results[id].Dob.Age
}

// Summary return a summary map
func (r *RandomuserScheme) Summary(id int) map[string]string {
	summary := make(map[string]string)

	summary["name"] = r.FullName(id)
	summary["gender"] = string(r.Gender(id))
	summary["picture"] = r.Picture(id, "medium")
	summary["age"] = string(r.Age(id))

	return summary
}
