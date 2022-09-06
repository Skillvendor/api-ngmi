package httpServer

import (
	adminAuthCtrl "api-ngmi/api/admin/auth"
	adminReportCtrl "api-ngmi/api/admin/report"
	commonS3Ctrl "api-ngmi/api/common/s3"
	mainAuthCtrl "api-ngmi/api/main/auth"
	mainReportCtrl "api-ngmi/api/main/report"
	mainUserCtrl "api-ngmi/api/main/user"
	"api-ngmi/middleware"
	"net/http"

	"github.com/bmizerany/pat"
)

func InitRoutes(mux *pat.PatternServeMux) {
	// public

	// report
	mux.Get("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainReportCtrl.Show)))
	mux.Get("/api/reports", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainReportCtrl.Index)))

	// user
	mux.Get("/api/user/:address", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainUserCtrl.Show)))
	mux.Post("/api/user", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainUserCtrl.Create)))

	// auth
	mux.Post("/api/auth/authentication", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainAuthCtrl.Authentication)))
	mux.Get("/api/auth/test", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(mainAuthCtrl.TestJWT, 1))))

	// private

	// s3
	mux.Get("/api/upload/asset", http.HandlerFunc(middleware.ApplyStandardMiddlewares(commonS3Ctrl.GetSignedUploadUrlAssets)))
	mux.Get("/api/upload/report", http.HandlerFunc(middleware.ApplyStandardMiddlewares(commonS3Ctrl.GetSignedUploadUrlReports)))
	mux.Get("/api/upload/report/read/:key", http.HandlerFunc(middleware.ApplyStandardMiddlewares(commonS3Ctrl.GetSignedDownloadUrlReports)))

	// admin

	// report
	mux.Put("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Update)))
	mux.Del("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Delete)))
	mux.Post("/api/reports", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Create)))
	mux.Patch("/api/reports/publish/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Publish)))
	mux.Patch("/api/reports/reject/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Unpublish)))
	mux.Get("/api/reports/admin/all", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Index)))
	mux.Get("/api/reports/admin/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminReportCtrl.Show)))

	// auth
	mux.Post("/api/auth/admin/authentication", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminAuthCtrl.Authentication)))

	http.Handle("/", mux)
}
