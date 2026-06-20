📦 1. Identity Federation (OIDC)
[ ] Implementar interface OIDCProvider
[ ] Criar provider Google (ExchangeCode + VerifyIDToken)
[ ] Implementar validação de ID Token (iss, aud, exp)
[ ] Implementar validação de assinatura via JWKS
[ ] Criar mecanismo de cache de JWKS (com TTL)
[ ] Implementar Provider Registry (suporte a múltiplos providers)
[ ] Normalizar claims (sub, email, name → modelo interno)
[ ] Implementar tratamento de erro para falhas OIDC
👤 2. Identity Core (Usuário + Identidade externa)
[ ] Criar entidade User
[ ] Criar entidade ExternalIdentity
[ ] Definir regra de unicidade (provider + provider_user_id)
[ ] Implementar UserRepository
[ ] Implementar ExternalIdentityRepository
[ ] Criar fluxo de criação de usuário via OIDC
[ ] Implementar idempotência no login federado
[ ] Criar relacionamento User ↔ ExternalIdentity
🔐 3. Auth-Service (Token interno)
[ ] Implementar TokenService (JWT)
[ ] Definir claims padrão (sub, iss, roles, exp)
[ ] Implementar assinatura JWT (RS256 recomendado)
[ ] Configurar tempo de expiração do token
[ ] Implementar refresh token (opcional)
[ ] Criar middleware de validação JWT
[ ] Garantir que apenas tokens internos são aceitos nas APIs
🔄 4. Use Case principal (Login OIDC)
[ ] Criar OIDCLoginUseCase
[ ] Implementar fluxo ExchangeCode → ID Token
[ ] Implementar validação de ID Token
[ ] Implementar busca de identidade externa
[ ] Criar usuário se não existir
[ ] Gerar token interno
[ ] Retornar resposta padronizada (AccessToken + UserID)
🌐 5. Interface HTTP
[ ] Criar endpoint POST /auth/oidc/login
[ ] Validar payload (provider, code)
[ ] Integrar handler com OIDCLoginUseCase
[ ] Padronizar resposta HTTP
[ ] Implementar tratamento de erros (401, 400, 500)
🔐 6. Segurança (crítico)
[ ] Forçar HTTPS em todos os endpoints
[ ] Validar issuer (iss) do ID Token
[ ] Validar audience (aud)
[ ] Validar expiração (exp)
[ ] Implementar PKCE no frontend
[ ] Proteger secrets (client_secret, keys)
[ ] Rotacionar chaves de assinatura JWT
⚖️ 7. LGPD Compliance
[ ] Mapear dados coletados (data inventory)
[ ] Classificar dados (identificadores, pessoais, sensíveis)
[ ] Minimizar dados armazenados (persistir apenas necessário)
[ ] Implementar endpoint DELETE /users/{id}
[ ] Implementar exclusão de ExternalIdentity
[ ] Definir política de retenção de dados
[ ] Criar logs de auditoria (sem dados sensíveis)
[ ] Garantir anonimização em logs
🧾 8. Auditoria e rastreabilidade
[ ] Criar tabela audit_logs
[ ] Registrar eventos de login (success/failure)
[ ] Registrar criação de usuário
[ ] Registrar vínculo com provider externo
[ ] Implementar correlation_id por request
[ ] Garantir rastreabilidade sem expor dados sensíveis
⚙️ 9. Infraestrutura
[ ] Implementar cache para JWKS
[ ] Configurar timeout para chamadas externas (OIDC)
[ ] Implementar retry policy para OIDC
[ ] Configurar variáveis de ambiente (client_id, issuer)
[ ] Preparar configuração multi-provider
🧠 10. Arquitetura e organização
[ ] Criar módulo federation (OIDC)
[ ] Criar separação clara entre domain/application/infrastructure
[ ] Garantir desacoplamento entre providers e domínio
[ ] Criar interfaces para facilitar testes
[ ] Implementar testes unitários para use cases
🛡️ 11. Preparação para Zero Trust
[ ] Garantir que todos serviços validam JWT
[ ] Implementar autenticação entre serviços (service-to-service)
[ ] Evitar confiança implícita na rede interna
[ ] Definir escopos/roles no token
🏢 12. Preparação para SOC 2
[ ] Implementar controle de acesso baseado em roles
[ ] Garantir logging estruturado
[ ] Documentar fluxos de autenticação
[ ] Implementar monitoramento de falhas de login
🔐 13. Preparação para ISO 27001
[ ] Criar política de gestão de credenciais
[ ] Definir controle de acesso formal
[ ] Implementar gestão de riscos (mesmo que simples)
[ ] Documentar arquitetura de segurança