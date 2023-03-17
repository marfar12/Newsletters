package transport

import (
	"net/http"

	"newsletter/transport/model"
	"newsletter/transport/util"

	"github.com/go-chi/chi/v5"
)

func getIdFromURL(r *http.Request) string {
	id := chi.URLParam(r, "id")

	return id
}

func (h *Handler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletter model.Newsletter
	err := util.UnmarshalRequest(r, &newsletter)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.Service.CreateNewsletter(r.Context(), model.ToSvcNewsletter(newsletter), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, newsletter)
}

func (h *Handler) GetNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletter, err := h.Service.GetNewsletter(r.Context(), getIdFromURL(r), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, model.ToNetNewsletter(newsletter))
}

func (h *Handler) ListNewsletters(w http.ResponseWriter, r *http.Request) {
	newsletters := h.Service.ListNewsletters(r.Context(), h.DB)

	util.WriteResponse(w, http.StatusOK, model.ToNetNewsletters(newsletters))
}

func (h *Handler) UpdateNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletter model.Newsletter
	err := util.UnmarshalRequest(r, &newsletter)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	newNewsletter, err := h.Service.UpdateNewsletter(r.Context(), getIdFromURL(r), model.ToSvcNewsletter(newsletter), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newNewsletter)
}

func (h *Handler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	err := h.Service.DeleteNewsletter(r.Context(), getIdFromURL(r), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusNoContent, nil)
}
