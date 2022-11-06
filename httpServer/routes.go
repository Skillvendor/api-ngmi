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
	mux.Get("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckJWTToken(mainReportCtrl.Show, -1))))
	mux.Get("/api/reports", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckJWTToken(mainReportCtrl.Index, -1))))

	// user
	mux.Post("/api/user", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainUserCtrl.Create)))

	// auth
	mux.Post("/api/auth/authentication", http.HandlerFunc(middleware.ApplyStandardMiddlewares(mainAuthCtrl.Authentication)))
	mux.Get("/api/auth/test", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(mainAuthCtrl.TestJWT, 1))))

	// admin

	// auth
	mux.Post("/api/auth/admin/authentication", http.HandlerFunc(middleware.ApplyStandardMiddlewares(adminAuthCtrl.Authentication)))

	// private

	// s3
	mux.Get("/api/upload/asset", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(commonS3Ctrl.GetSignedUploadUrlAssets, 6))))
	mux.Get("/api/upload/report", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(commonS3Ctrl.GetSignedUploadUrlReports, 6))))
	mux.Get("/api/upload/report/read/:key", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(commonS3Ctrl.GetSignedDownloadUrlReports, 6))))
	mux.Get("/api/report/read/:key", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckJWTToken(commonS3Ctrl.GetSignedDownloadUrlReports, 3))))

	// user
	mux.Get("/api/user", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckJWTToken(mainUserCtrl.Show, 0))))
	mux.Post("/api/user/refreshAccessLevelCache", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckJWTToken(mainUserCtrl.ResetAccessLevelCache, 0))))

	// admin

	// report
	mux.Put("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Update, 6))))
	mux.Del("/api/reports/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Delete, 7))))
	mux.Post("/api/reports", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Create, 6))))
	mux.Patch("/api/reports/publish/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Publish, 7))))
	mux.Patch("/api/reports/reject/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Unpublish, 7))))
	mux.Get("/api/reports/admin/all", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Index, 6))))
	mux.Get("/api/reports/admin/:id", http.HandlerFunc(middleware.ApplyStandardMiddlewares(middleware.CheckAdminJWTToken(adminReportCtrl.Show, 6))))

	http.Handle("/", mux)
}
