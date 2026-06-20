package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/middlewares"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/response"
	httpresponse "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/response"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/handlers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/usecases"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/config"
)

type AuthController struct {
	registroComEmailHandler   *handlers.RegistroComEmailHandler
	loginComEmailHandler      *handlers.LoginComEmailHandler
	loginComGoogleHandler     *usecases.OIDCLoginUseCase
	meHandler                 *handlers.MeHandler
	validateRefreshHandler    *handlers.ValidateRefreshTokenHandler
	updateRefreshTokenHandler *handlers.UpdateRefreshTokenHandler
	revokeRefreshTokenHandler *handlers.RevokeTokenHandler
	logoutHandler             *handlers.LogoutHandler
	config                    *config.Config
}

func NewAuthController(
	registro_com_email *handlers.RegistroComEmailHandler,
	login_com_email *handlers.LoginComEmailHandler,
	login_com_google *usecases.OIDCLoginUseCase,
	me *handlers.MeHandler,
	logout *handlers.LogoutHandler,
	validate_refresh_token *handlers.ValidateRefreshTokenHandler,
	update_refresh_token *handlers.UpdateRefreshTokenHandler,
	revoke_refresh_token *handlers.RevokeTokenHandler,
	config *config.Config,

) *AuthController {
	return &AuthController{
		registroComEmailHandler:   registro_com_email,
		loginComEmailHandler:      login_com_email,
		loginComGoogleHandler:     login_com_google,
		meHandler:                 me,
		validateRefreshHandler:    validate_refresh_token,
		updateRefreshTokenHandler: update_refresh_token,
		revokeRefreshTokenHandler: revoke_refresh_token,
		logoutHandler:             logout,
		config:                    config,
	}
}

// @Summary Registrar usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body commands.RegistroComEmailCommand true "Dados do usuário"
// @Success 201 "ok"
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/register [post]
func (c *AuthController) RegistroComEmail(w http.ResponseWriter, r *http.Request) {

	reqCtx, _ := middlewares.GetRequestContext(r)
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB

	cmd := commands.RegistroComEmailCommand{}
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, "invalid body", "invalid_request")
		return
	}

	result := c.registroComEmailHandler.Handle(r.Context(), reqCtx, cmd)
	if result != nil {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, result.Error(), "registration_failed")
		return
	}

	httpresponse.WriteJSON(w, http.StatusCreated, result)
}

// @Summary Login com email
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body commands.LoginComEmailCommand true "Credenciais"
// @Success 200 {object} dto.TokenResponseDTO
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/login [post]
func (c *AuthController) LoginComEmail(w http.ResponseWriter, r *http.Request) {

	authCtx, _ := middlewares.GetAuthContext(r)

	cmd := commands.LoginComEmailCommand{}
	log.Println("Body: ", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, "invalid body", "invalid_request")
		return
	}

	result, err := c.loginComEmailHandler.Handle(r.Context(), authCtx, cmd)
	if err != nil {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, err.Error(), "invalid_credentials")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// @Summary Retorna usuário autenticado
// @Description Retorna os dados do usuário autenticado via JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.MeResponseDTO
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/me [get]
func (c *AuthController) Me(
	w http.ResponseWriter,
	r *http.Request,
) {

	authCtx, ok :=
		middlewares.GetAuthContext(r)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := c.meHandler.Handle(r.Context(), authCtx)
	if err != nil {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, err.Error(), "invalid_credentials")
		return
	}
	log.Println("Result da REQ ME: ", result)

	response.WriteJSON(
		w,
		http.StatusOK,
		result,
	)
}

// @Summary Iniciar login com Google
// @Tags Auth OIDC
// @Produce plain
// @Success 307 "Redirect para Google"
// @Router /auth/login/google [get]
func (c *AuthController) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	redirectURI := "https://localhost/api/v1/auth-service/auth/google/callback"

	clientID := c.config.OIDCGoogleClientID // ideal: vir de config

	url := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid%%20email%%20profile&access_type=offline&prompt=consent",
		clientID,
		url.QueryEscape(redirectURI),
	)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// @Summary Callback do Google OAuth
// @Tags Auth OIDC
// @Produce plain
// @Param code query string true "Authorization code"
// @Success 307 "Redirect para frontend com tokens"
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/google/callback [get]
func (c *AuthController) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	reqCtx, _ := middlewares.GetRequestContext(r)

	code := r.URL.Query().Get("code")
	if code == "" {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, "missing code", "invalid_request")
		return
	}

	redirectURI := "https://localhost/api/v1/auth-service/auth/google/callback"

	cmd := commands.LoginComOIDCCommand{
		Provider:    "google",
		Code:        code,
		RedirectURI: redirectURI,
	}

	access, refresh, _, err := c.loginComGoogleHandler.Execute(r.Context(), reqCtx, cmd)
	if err != nil {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, err.Error(), "oidc_failed")
		return
	}

	// 🔥 Redireciona pro frontend com tokens
	frontendRedirect := fmt.Sprintf(
		"https://localhost/?access=%s&refresh=%s",
		url.QueryEscape(access),
		url.QueryEscape(refresh),
	)

	http.Redirect(w, r, frontendRedirect, http.StatusTemporaryRedirect)
}

// @Summary Validar access token (Nginx)
// @Tags Auth Internal
// @Produce plain
// @Success 200 "Headers com dados do usuário"
// @Failure 401
// @Router /auth/validate [get]
func (c *AuthController) Validate(w http.ResponseWriter, r *http.Request) {

	authCtx, ok := middlewares.GetAuthContext(r)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println("CTX:", r.Context().Value(middlewares.AuthContextKey))
	w.Header().Set("X-User-ID", authCtx.Autor.ID)

	if len(authCtx.Roles.Tipo) > 0 {
		w.Header().Set("X-User-Role", string(authCtx.Roles.Tipo))
	}

	w.Header().Set("X-User-Canal", authCtx.Origem.Canal.String())
	w.Header().Set("X-User-Sistema", authCtx.Origem.Sistema.String())

	w.WriteHeader(http.StatusOK)
}

// @Summary Renovar tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body commands.RefreshTokenPayload true "Refresh token"
// @Success 200 {object} dto.TokenResponseDTO
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/token/refresh [post]
func (c *AuthController) Refresh(w http.ResponseWriter, r *http.Request) {

	authCtx, ok := middlewares.GetAuthContext(r)
	if !ok {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}

	cmd := commands.RefreshTokenPayload{}
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, "invalid body", "invalid_request")
		return
	}

	result, err := c.updateRefreshTokenHandler.Handle(r.Context(), authCtx, cmd)
	if err != nil {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, err.Error(), "refresh_failed")
		return
	}

	json.NewEncoder(w).Encode(result)
}

// @Summary Revogar refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body commands.RefreshTokenPayload true "Refresh token"
// @Success 204
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/token/revoke [post]
func (c *AuthController) Revoke(w http.ResponseWriter, r *http.Request) {
	authCtx, ok := middlewares.GetAuthContext(r)
	if !ok {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}

	cmd := commands.RefreshTokenPayload{}
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, "invalid body", "invalid_request")
		return
	}

	if err := c.revokeRefreshTokenHandler.Handle(r.Context(), authCtx, cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusInternalServerError, err.Error(), "revoke_failed")
		return
	}

	response.WriteNoContent(w)
}

// @Summary Logout do usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body commands.RefreshTokenPayload true "Logout payload"
// @Success 204
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/logout [post]
func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	authCtx, ok := middlewares.GetAuthContext(r)
	if !ok {
		httpresponse.WriteJSONError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}

	cmd := commands.RefreshTokenPayload{}
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusBadRequest, "invalid body", "invalid_request")
		return
	}

	if err := c.logoutHandler.Handle(r.Context(), authCtx, cmd); err != nil {
		httpresponse.WriteJSONError(w, http.StatusInternalServerError, err.Error(), "logout_failed")
		return
	}

	response.WriteNoContent(w)
}
