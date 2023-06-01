package reearthcmsapiwebhook

import (
	"encoding/json"
	"io"
	"net/http"
)

type MiddlewareConfig struct {
	Secret []byte
	Logger func(format string, v ...any)
}

func Middleware(config MiddlewareConfig) func(http.Handler) http.Handler {
	if config.Logger == nil {
		config.Logger = func(format string, v ...any) {}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				jsonResp(w, http.StatusUnprocessableEntity, map[string]string{"error": "unprocessable entity"})
				return
			}

			sig := r.Header.Get(SignatureHeader)
			config.Logger("webhook: received: sig=%s", sig)
			if !validateSignature(sig, body, config.Secret) {
				jsonResp(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
				return
			}

			config.Logger("webhook: body: %s", body)

			p := &Payload{}
			if err := json.Unmarshal(body, p); err != nil {
				jsonResp(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
				return
			}

			p.Body = body
			p.Sig = sig
			ctx := AttachPayload(r.Context(), p)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func jsonResp(w http.ResponseWriter, code int, msg any) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(msg)
}
