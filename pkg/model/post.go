package model

import "time"

// Post is our main model for Posts
type Post struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
}

// PostDTO is our data transfer object for Post
type PostDTO struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// ToPost converts postDTO to post
func ToPost(postDTO *PostDTO) *Post {
	return &Post{
		Title: postDTO.Title,
		Body:  postDTO.Body,
	}
}

// ToPostDTO converts post to postDTO
func ToPostDTO(post *Post) *PostDTO {
	return &PostDTO{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
}

/* Example JSON
{
	"Title":"Dummy Title",
	"Body":"Dummy content",
}
*/
