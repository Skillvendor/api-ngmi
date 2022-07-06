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
	mux.Patch("/api/reports/reject/:id", http.HandlerFunc(report.RejectReview))
	mux.Get("/api/reports/admin/all", http.HandlerFunc(report.GetAllReports))
	mux.Get("/api/reports/admin/:id", http.HandlerFunc(report.GetReportAdmin))

	// s3
	mux.Get("/api/assets/upload", http.HandlerFunc(s3.GetSignedUploadUrlAssets))
	mux.Get("/api/reports/upload", http.HandlerFunc(s3.GetSignedUploadUrlReports))
	mux.Get("/api/reports/read/:key", http.HandlerFunc(s3.GetSignedDownloadUrlReports))

	http.Handle("/", mux)
}
