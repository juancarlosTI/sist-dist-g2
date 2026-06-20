package usecases_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	commands "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
// 	usecases "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/usecases"
// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
// 	svc "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token"
// 	oidc "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/federation/oidc"
// 	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/common"
// )

// type mockProvider struct {
// 	tokenResp *oidc.TokenResponse
// 	claims    *oidc.NormalizedClaims
// }

// func (m *mockProvider) Name() string { return "google" }
// func (m *mockProvider) ExchangeCode(ctx context.Context, code string, redirectURI string) (*oidc.TokenResponse, error) {
// 	if m.tokenResp == nil {
// 		return nil, errors.New("no token")
// 	}
// 	return m.tokenResp, nil
// }
// func (m *mockProvider) VerifyIDToken(ctx context.Context, idToken string) (*oidc.NormalizedClaims, error) {
// 	if m.claims == nil {
// 		return nil, errors.New("invalid")
// 	}
// 	return m.claims, nil
// }

// type mockUserRepo struct {
// 	byID    map[string]*identidade.User
// 	byEmail map[string]*identidade.User
// 	saved   []*identidade.User
// }

// func (m *mockUserRepo) Salvar(ctx context.Context, user *identidade.User) error {
// 	m.saved = append(m.saved, user)
// 	return nil
// }
// func (m *mockUserRepo) FindByID(ctx context.Context, id string) (*identidade.User, error) {
// 	if u, ok := m.byID[id]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("not found")
// }
// func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*identidade.User, error) {
// 	if u, ok := m.byEmail[email]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("not found")
// }

// type mockExternalRepo struct {
// 	byKey map[string]*identidade.ExternalIdentity
// 	saved []*identidade.ExternalIdentity
// }

// func key(provider, pid string) string { return provider + "|" + pid }
// func (m *mockExternalRepo) Salvar(ctx context.Context, ei *identidade.ExternalIdentity) error {
// 	m.saved = append(m.saved, ei)
// 	m.byKey[key(ei.Provider(), ei.ProviderUserID())] = ei
// 	return nil
// }
// func (m *mockExternalRepo) FindByProviderAndProviderID(ctx context.Context, provider string, providerUserID string) (*identidade.ExternalIdentity, error) {
// 	if ei, ok := m.byKey[key(provider, providerUserID)]; ok {
// 		return ei, nil
// 	}
// 	return nil, errors.New("not found")
// }
// func (m *mockExternalRepo) FindByUserID(ctx context.Context, userID string) ([]*identidade.ExternalIdentity, error) {
// 	var res []*identidade.ExternalIdentity
// 	for _, v := range m.byKey {
// 		if v.UserID() == userID {
// 			res = append(res, v)
// 		}
// 	}
// 	return res, nil
// }
// func (m *mockExternalRepo) Delete(ctx context.Context, provider string, providerUserID string) error {
// 	delete(m.byKey, key(provider, providerUserID))
// 	return nil
// }

// type mockAccessSvc struct{ tok string }

// func (m *mockAccessSvc) Generate(at *token.AccessToken) (string, error) { return m.tok, nil }
// func (m *mockAccessSvc) Validate(ctx context.Context, token string) (*token.AccessToken, error) {
// 	return nil, nil
// }

// type mockRefreshSvc struct{ rt *svc.RefreshToken }

// func (m *mockRefreshSvc) Create(ctx context.Context, userID string) (*svc.RefreshToken, error) {
// 	return m.rt, nil
// }
// func (m *mockRefreshSvc) Rotate(ctx context.Context, oldValue string) (*svc.RefreshToken, error) {
// 	return nil, nil
// }
// func (m *mockRefreshSvc) Revoke(ctx context.Context, value string) error   { return nil }
// func (m *mockRefreshSvc) Validate(ctx context.Context, value string) error { return nil }

// func TestOIDCLogin_ExistingExternalIdentity(t *testing.T) {
// 	prov := &mockProvider{tokenResp: &oidc.TokenResponse{IDToken: "id"}, claims: &oidc.NormalizedClaims{Subject: "sub-1", Email: "e@x.com"}}
// 	providers := oidc.NewProviderRegistry(prov)

// 	user, _ := identidade.RegistrarUsuario("e@x.com", shared_context.EventContext{})
// 	userRepo := &mockUserRepo{byID: map[string]*identidade.User{user.ID(): user}, byEmail: map[string]*identidade.User{}}

// 	external, _ := identidade.NovoExternalIdentity(user.ID(), "google", "sub-1")
// 	extRepo := &mockExternalRepo{byKey: map[string]*identidade.ExternalIdentity{key("google", "sub-1"): external}}

// 	access := &mockAccessSvc{tok: "access-token"}
// 	refresh := &mockRefreshSvc{rt: &svc.RefreshToken{Value: "refresh-token"}}

// 	uc := usecases.NewOIDCLoginUseCase(providers, access, refresh, userRepo, extRepo)

// 	at, rt, uid, err := uc.Execute(context.Background(), commands.LoginComOIDCCommand{Provider: "google", Code: "c", RedirectURI: "r"})
// 	if err != nil {
// 		t.Fatalf("unexpected err: %v", err)
// 	}
// 	if at != "access-token" || rt != "refresh-token" || uid != user.ID() {
// 		t.Fatalf("unexpected result: %s %s %s (expected user id %s)", at, rt, uid, user.ID())
// 	}
// }
