package httpServer

import (
	"api-ngmi/api/admin"
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
	mux.Get("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.GetReport)))
	mux.Get("/api/reports", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.GetPublishedReports)))

	// user
	mux.Get("/api/user/:address", http.HandlerFunc(middleware.ApplyStandardMiddlewares(user.Show)))
	mux.Post("/api/user", http.HandlerFunc(middleware.ApplyStandardMiddlewares(user.Create)))

	// auth
	mux.Post("/api/auth/authentication", http.HandlerFunc(middleware.ApplyStandardMiddlewares(auth.Authentication)))
	mux.Get("/api/auth/test", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(auth.TestJWT, 1))))

	// private

	// report
	mux.Put("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.UpdateReport)))
	mux.Del("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.DeleteReport)))
	mux.Post("/api/reports", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.CreateReport)))
	mux.Patch("/api/reports/publish/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.ApproveReview)))
	mux.Patch("/api/reports/reject/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.RejectReview)))
	mux.Get("/api/reports/admin/all", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.GetAllReports)))
	mux.Get("/api/reports/admin/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(report.GetReportAdmin)))

	// s3
	mux.Get("/api/upload/asset", http.HandlerFunc(middleware.ApplyStandardMiddlewares(s3.GetSignedUploadUrlAssets)))
	mux.Get("/api/upload/report", http.HandlerFunc(middleware.ApplyStandardMiddlewares(s3.GetSignedUploadUrlReports)))
	mux.Get("/api/upload/report/read/:key", http.HandlerFunc(middleware.ApplyStandardMiddlewares(s3.GetSignedDownloadUrlReports)))

	// admin
	mux.Post("/api/auth/admin/authentication", http.HandlerFunc(middleware.ApplyStandardMiddlewares(admin.Authentication)))

	http.Handle("/", mux)
}
