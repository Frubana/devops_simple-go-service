package main
import (
    "net/http/httptest"
    "net/http"
    "testing"
)

func TestHealthCheck(t *testing.T) {
    router:= RouteBuilder()
    req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
    rw :=  httptest.NewRecorder()
    router.ServeHTTP(rw, req)
    resp := rw.Result()
    if resp.StatusCode != 200 {
        t.Errorf("IntMin(2, -2) = %d; want -2", resp.StatusCode)
    }
}