package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/yakuter/go-clean-code/pkg/model"
	"github.com/yakuter/go-clean-code/pkg/service"
)

// PostAPI ...
type PostAPI struct {
	PostService service.PostService
}

// NewPostAPI ...
func NewPostAPI(p service.PostService) PostAPI {
	return PostAPI{PostService: p}
}

// FindAllPosts ...
func (p PostAPI) FindAllPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := p.PostService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, posts)
	}
}

// FindByID ...
func (p PostAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if id is integer
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Find login by id from db
		post, err := p.PostService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToPostDTO(post))
	}
}

// CreatePost ...
func (p PostAPI) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var postDTO model.PostDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&postDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdPost, err := p.PostService.Save(model.ToPost(&postDTO))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToPostDTO(createdPost))
	}
}

// UpdatePost ...
func (p PostAPI) UpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var postDTO model.PostDTO
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&postDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		post, err := p.PostService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		post.Title = postDTO.Title
		post.Body = postDTO.Body
		updatedPost, err := p.PostService.Save(post)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToPostDTO(updatedPost))
	}
}

// DeletePost ...
func (p PostAPI) DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		post, err := p.PostService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = p.PostService.Delete(post.ID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		response := Response{
			Message: "Post deleted successfully!",
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}

// Migrate ...
func (p PostAPI) Migrate() {
	err := p.PostService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
