package permissions

// KYCPermissions returns KYC-domain permissions owned by kyc-service.
// Canonical keys are fine-grained (kyc.*); colon keys remain as legacy aliases.
func KYCPermissions() []Permission {
	return []Permission{
		// Canonical fine-grained
		{Key: KYCSubmissionListKey, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "List KYC Submissions", Description: "List user KYC submissions", Route: "GET /backoffice/kyc/submissions/list"},
		{Key: KYCSubmissionReadKey, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "Read KYC Submission", Description: "View user KYC submission details", Route: "GET /backoffice/kyc/submissions/:user_id"},
		{Key: KYCDocumentReviewKey, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "Review KYC Document", Description: "Approve or reject KYC documents", Route: "PATCH /backoffice/kyc/documents/:id/review"},
		{Key: KYCSubmissionReviewKey, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "Review KYC Submission", Description: "Approve, reject, or request more information for a KYC case", Route: "POST /backoffice/kyc/submissions/:user_id/review"},

		// Legacy colon aliases (no Route — ByRoute prefers canonical keys)
		{Key: KYCSubmissionList, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "List KYC Submissions (legacy)", Description: "Legacy alias for kyc.submission.list"},
		{Key: KYCSubmissionRead, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "Read KYC Submission (legacy)", Description: "Legacy alias for kyc.submission.read"},
		{Key: KYCDocumentReview, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "Review KYC Document (legacy)", Description: "Legacy alias for kyc.document.review"},
		{Key: KYCSubmissionReview, Service: ServiceKYC, Module: "kyc", Group: "KYC", Name: "Review KYC Submission (legacy)", Description: "Legacy alias for kyc.submission.review"},
	}
}
