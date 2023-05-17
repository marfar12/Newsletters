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

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var editor model.Editor
	err := util.UnmarshalRequest(r, &editor)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	signedInEditor, err := h.NewsletterService.SignIn(r.Context(), model.ToSvcEditor(editor), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	netEditor := model.ToNetEditor(signedInEditor)
	token := model.CreateToken(netEditor)

	util.WriteResponse(w, http.StatusCreated, token)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var editor model.Editor
	err := util.UnmarshalRequest(r, &editor)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	signedUpEditor, err := h.NewsletterService.SignUp(r.Context(), model.ToSvcEditor(editor), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, model.ToNetEditor(signedUpEditor))
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var subscription model.Subscription
	err := util.UnmarshalRequest(r, &subscription)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	subscriptionRes, err := h.NewsletterService.Subscribe(r.Context(), model.ToSvcSubscription(subscription), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, model.ToNetSubscription(subscriptionRes))
}

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	err := h.NewsletterService.Unsubscribe(r.Context(), util.GetIdFromURL(r), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, "Unsubscribed successfully")
}

func (h *Handler) Publish(w http.ResponseWriter, r *http.Request) {
	var issue model.Issue
	err := util.UnmarshalRequest(r, &issue)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.NewsletterService.Publish(r.Context(), model.ToSvcIssue(issue), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, "Issue published successfully")
}
