package httpServer

import (
	"api-ngmi/api/auth"
	"api-ngmi/api/report"
	"api-ngmi/api/s3"
	"api-ngmi/api/user"
	"api-ngmi/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	// public

	// report
	mux.HandleFunc("/api/report/all", report.GetPublishedReports)
	mux.HandleFunc("/api/report/show", report.GetReport)

	// user
	mux.HandleFunc("/api/user/create", user.CreateUser)
	mux.HandleFunc("/api/user/show", user.GetPublicUser)

	// auth
	mux.HandleFunc("/api/auth/authentication", auth.Authentication)
	mux.HandleFunc("/api/auth/test", middleware.CheckJWTToken(auth.TestJWT, 1))

	// private

	// report
	mux.HandleFunc("/api/report/create", report.CreateReport)
	mux.HandleFunc("/api/report/update", middleware.CheckJWTToken(report.UpdateReport, 5))
	mux.HandleFunc("/api/report/delete", middleware.CheckJWTToken(report.DeleteReport, 6))
	mux.HandleFunc("/api/report/approveReview", middleware.CheckJWTToken(report.ApproveReview, 6))
	mux.HandleFunc("/api/report/admin/all", middleware.CheckJWTToken(report.GetAllReports, 5))

	// s3
	mux.HandleFunc("/api/signedUrl", middleware.CheckJWTToken(s3.GetSignedUploadUrl, 5))
}
