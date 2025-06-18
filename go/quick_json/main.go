package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// * create struct for JSON
type PostData struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// * function to get JSON and update struct pointer
func getPost(page int) (*PostData, error) {
	var data PostData
	r, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + fmt.Sprint(page))
	if err != nil {
		fmt.Println("Error fetching post data:", err)
		return nil, err
	}

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding post body:", err)
		return nil, err
	}

	return &data, nil
}

func getAllPosts() {
	var postData []PostData

	// ! get post and verify get logic
	// * get multiple posts

	for i := 1; i < 10; i++ {
		data, err := getPost(i)
		if err != nil {
			fmt.Println("Failed to get post:", err)
		}
		postData = append(postData, *data)
	}

	fmt.Print(postData)
}

// * logic around get
func main() {
	getAllPosts()
}
