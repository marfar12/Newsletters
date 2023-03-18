package transport

import (
	"net/http"

	"newsletter/transport/model"
	"newsletter/transport/util"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var editor model.Editor
	err := util.UnmarshalRequest(r, &editor)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	signedInEditor, err := h.EditorService.SignIn(r.Context(), model.ToSvcEditor(editor), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, model.ToNetEditor(signedInEditor))
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var editor model.Editor
	err := util.UnmarshalRequest(r, &editor)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	signedUpEditor, err := h.EditorService.SignUp(r.Context(), model.ToSvcEditor(editor), h.DB)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, model.ToNetEditor(signedUpEditor))
}
