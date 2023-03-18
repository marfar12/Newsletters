package transport

import (
	"net/http"

	"newsletter/transport/model"
	"newsletter/transport/util"
)

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

	id, err := h.NewsletterService.CreateNewsletter(r.Context(), model.ToSvcNewsletter(newsletter), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, id)
}

func (h *Handler) GetNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletter, err := h.NewsletterService.GetNewsletter(r.Context(), util.GetIdFromURL(r), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, model.ToNetNewsletter(newsletter))
}

func (h *Handler) ListNewsletters(w http.ResponseWriter, r *http.Request) {
	newsletters, err := h.NewsletterService.ListNewsletters(r.Context(), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}

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

	newNewsletter, err := h.NewsletterService.UpdateNewsletter(r.Context(), util.GetIdFromURL(r), model.ToSvcNewsletter(newsletter), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newNewsletter)
}

func (h *Handler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	err := h.NewsletterService.DeleteNewsletter(r.Context(), util.GetIdFromURL(r), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusNoContent, nil)
}
