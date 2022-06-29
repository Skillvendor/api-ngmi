package httpServer

import (
	"api-ngmi/api/auth"
	"api-ngmi/api/report"
	"api-ngmi/api/s3"
	"api-ngmi/api/user"
	"api-ngmi/middleware"
	"net/http"

	"github.com/bmizerany/pat"
)

func InitRoutes(mux *pat.PatternServeMux) {
	// public

	// report
	mux.Get("/api/reports/:id", http.HandlerFunc(report.GetReport))
	mux.Get("/api/reports", http.HandlerFunc(report.GetPublishedReports))

	// user
	mux.Get("/api/user/:address", http.HandlerFunc(user.GetPublicUser))
	mux.Post("/api/user", http.HandlerFunc(user.CreateUser))

	// auth
	mux.Post("/api/auth/authentication", http.HandlerFunc(auth.Authentication))
	mux.Get("/api/auth/test", http.HandlerFunc(middleware.CheckJWTToken(auth.TestJWT, 1)))

	// private

	// report
	mux.Put("/api/reports/:id", http.HandlerFunc(report.UpdateReport))
	mux.Del("/api/reports/:id", http.HandlerFunc(report.DeleteReport))
	mux.Post("/api/reports", http.HandlerFunc(report.CreateReport))
	mux.Patch("/api/reports/publish/:id", http.HandlerFunc(report.ApproveReview))
	mux.Get("/api/reports/admin/all", http.HandlerFunc(middleware.CheckJWTToken(report.GetAllReports, 5)))

	// s3
	mux.Get("/api/signedUrl", http.HandlerFunc(middleware.CheckJWTToken(s3.GetSignedUploadUrl, 5)))

	http.Handle("/", mux)
}
