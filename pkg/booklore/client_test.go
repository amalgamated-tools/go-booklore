package booklore

import (
	"context"
	"fmt"
	"os"
	"testing"
)

// Generated test functions for API endpoints
// These tests hit the real API endpoint configured in environment variables

// TestCatchall_GET tests the GET /api/kobo/{token} endpoint
func TestCatchall_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for catchAll
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCatchall2_PUT tests the PUT /api/kobo/{token} endpoint
func TestCatchall2_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for catchAll_2
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCatchall1_POST tests the POST /api/kobo/{token} endpoint
func TestCatchall1_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for catchAll_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCatchall3_DELETE tests the DELETE /api/kobo/{token} endpoint
func TestCatchall3_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for catchAll_3
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCatchall4_PATCH tests the PATCH /api/kobo/{token} endpoint
func TestCatchall4_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for catchAll_4
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetemailrecipient_GET tests the GET /api/v2/email/recipients/{id} endpoint
func TestGetemailrecipient_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getEmailRecipient
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateemailrecipient_PUT tests the PUT /api/v2/email/recipients/{id} endpoint
func TestUpdateemailrecipient_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateEmailRecipient
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteemailrecipient_DELETE tests the DELETE /api/v2/email/recipients/{id} endpoint
func TestDeleteemailrecipient_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteEmailRecipient
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetemailprovider_GET tests the GET /api/v2/email/providers/{id} endpoint
func TestGetemailprovider_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getEmailProvider
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateemailprovider_PUT tests the PUT /api/v2/email/providers/{id} endpoint
func TestUpdateemailprovider_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateEmailProvider
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteemailprovider_DELETE tests the DELETE /api/v2/email/providers/{id} endpoint
func TestDeleteemailprovider_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteEmailProvider
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetnotebyid_GET tests the GET /api/v2/book-notes/{noteId} endpoint
func TestGetnotebyid_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getNoteById
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatenote_PUT tests the PUT /api/v2/book-notes/{noteId} endpoint
func TestUpdatenote_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateNote
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletenote_DELETE tests the DELETE /api/v2/book-notes/{noteId} endpoint
func TestDeletenote_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteNote
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetuser_GET tests the GET /api/v1/users/{id} endpoint
func TestGetuser_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateuser_PUT tests the PUT /api/v1/users/{id} endpoint
func TestUpdateuser_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteuser_DELETE tests the DELETE /api/v1/users/{id} endpoint
func TestDeleteuser_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateusersetting_PUT tests the PUT /api/v1/users/{id}/settings endpoint
func TestUpdateusersetting_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateUserSetting
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestChangeuserpassword_PUT tests the PUT /api/v1/users/change-user-password endpoint
func TestChangeuserpassword_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for changeUserPassword
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestChangepassword_PUT tests the PUT /api/v1/users/change-password endpoint
func TestChangepassword_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for changePassword
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetshelf_GET tests the GET /api/v1/shelves/{shelfId} endpoint
func TestGetshelf_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateshelf_PUT tests the PUT /api/v1/shelves/{shelfId} endpoint
func TestUpdateshelf_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteshelf_DELETE tests the DELETE /api/v1/shelves/{shelfId} endpoint
func TestDeleteshelf_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetappsettings_GET tests the GET /api/v1/settings endpoint
func TestGetappsettings_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAppSettings
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatesettings_PUT tests the PUT /api/v1/settings endpoint
func TestUpdatesettings_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateSettings
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetlibrary_GET tests the GET /api/v1/libraries/{libraryId} endpoint
func TestGetlibrary_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatelibrary_PUT tests the PUT /api/v1/libraries/{libraryId} endpoint
func TestUpdatelibrary_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletelibrary_DELETE tests the DELETE /api/v1/libraries/{libraryId} endpoint
func TestDeletelibrary_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRescanlibrary_PUT tests the PUT /api/v1/libraries/{libraryId}/refresh endpoint
func TestRescanlibrary_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for rescanLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcurrentuser_GET tests the GET /api/v1/koreader-users/me endpoint
func TestGetcurrentuser_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getCurrentUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpsertcurrentuser_PUT tests the PUT /api/v1/koreader-users/me endpoint
func TestUpsertcurrentuser_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for upsertCurrentUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsettings_GET tests the GET /api/v1/kobo-settings endpoint
func TestGetsettings_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSettings
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatesettings1_PUT tests the PUT /api/v1/kobo-settings endpoint
func TestUpdatesettings1_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateSettings_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateorupdatetoken_PUT tests the PUT /api/v1/kobo-settings/token endpoint
func TestCreateorupdatetoken_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createOrUpdateToken
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsettings1_GET tests the GET /api/v1/hardcover-sync-settings endpoint
func TestGetsettings1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSettings_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatesettings2_PUT tests the PUT /api/v1/hardcover-sync-settings endpoint
func TestUpdatesettings2_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateSettings_2
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookviewersettings_GET tests the GET /api/v1/books/{bookId}/viewer-setting endpoint
func TestGetbookviewersettings_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookViewerSettings
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatebookviewersettings_PUT tests the PUT /api/v1/books/{bookId}/viewer-setting endpoint
func TestUpdatebookviewersettings_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateBookViewerSettings
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatemetadata_PUT tests the PUT /api/v1/books/{bookId}/metadata endpoint
func TestUpdatemetadata_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatepersonalrating_PUT tests the PUT /api/v1/books/personal-rating endpoint
func TestUpdatepersonalrating_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updatePersonalRating
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestTogglefieldlocks_PUT tests the PUT /api/v1/books/metadata/toggle-field-locks endpoint
func TestTogglefieldlocks_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for toggleFieldLocks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestToggleallmetadata_PUT tests the PUT /api/v1/books/metadata/toggle-all-lock endpoint
func TestToggleallmetadata_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for toggleAllMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestBulkeditmetadata_PUT tests the PUT /api/v1/books/bulk-edit-metadata endpoint
func TestBulkeditmetadata_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for bulkEditMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookmarkbyid_GET tests the GET /api/v1/bookmarks/{bookmarkId} endpoint
func TestGetbookmarkbyid_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookmarkById
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatebookmark_PUT tests the PUT /api/v1/bookmarks/{bookmarkId} endpoint
func TestUpdatebookmark_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateBookmark
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletebookmark_DELETE tests the DELETE /api/v1/bookmarks/{bookmarkId} endpoint
func TestDeletebookmark_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteBookmark
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetannotationbyid_GET tests the GET /api/v1/annotations/{annotationId} endpoint
func TestGetannotationbyid_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAnnotationById
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateannotation_PUT tests the PUT /api/v1/annotations/{annotationId} endpoint
func TestUpdateannotation_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateAnnotation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteannotation_DELETE tests the DELETE /api/v1/annotations/{annotationId} endpoint
func TestDeleteannotation_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteAnnotation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatestatus_PUT tests the PUT /api/mobile/v1/books/{bookId}/status endpoint
func TestUpdatestatus_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateStatus
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdaterating_PUT tests the PUT /api/mobile/v1/books/{bookId}/rating endpoint
func TestUpdaterating_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateRating
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateprogress_PUT tests the PUT /api/koreader/syncs/progress endpoint
func TestUpdateprogress_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateProgress
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetstate_GET tests the GET /api/kobo/{token}/v1/library/{bookId}/state endpoint
func TestGetstate_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getState
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatestate_PUT tests the PUT /api/kobo/{token}/v1/library/{bookId}/state endpoint
func TestUpdatestate_PUT(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateState
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetusers_GET tests the GET /api/v2/opds-users endpoint
func TestGetusers_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getUsers
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateuser_POST tests the POST /api/v2/opds-users endpoint
func TestCreateuser_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetemailrecipients_GET tests the GET /api/v2/email/recipients endpoint
func TestGetemailrecipients_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getEmailRecipients
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateemailrecipient_POST tests the POST /api/v2/email/recipients endpoint
func TestCreateemailrecipient_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createEmailRecipient
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetemailproviders_GET tests the GET /api/v2/email/providers endpoint
func TestGetemailproviders_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getEmailProviders
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateemailprovider_POST tests the POST /api/v2/email/providers endpoint
func TestCreateemailprovider_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createEmailProvider
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSendemail_POST tests the POST /api/v2/email/book endpoint
func TestSendemail_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for sendEmail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestEmailbookquick_POST tests the POST /api/v2/email/book/{bookId} endpoint
func TestEmailbookquick_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for emailBookQuick
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreatenote_POST tests the POST /api/v2/book-notes endpoint
func TestCreatenote_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createNote
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestStarttask_POST tests the POST /api/v1/tasks/start endpoint
func TestStarttask_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for startTask
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetallshelves_GET tests the GET /api/v1/shelves endpoint
func TestGetallshelves_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllShelves
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateshelf_POST tests the POST /api/v1/shelves endpoint
func TestCreateshelf_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSetupfirstuser_POST tests the POST /api/v1/setup endpoint
func TestSetupfirstuser_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for setupFirstUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRefreshreviews_POST tests the POST /api/v1/reviews/book/{bookId}/refresh endpoint
func TestRefreshreviews_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for refreshReviews
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRecordreadingsession_POST tests the POST /api/v1/reading-sessions endpoint
func TestRecordreadingsession_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for recordReadingSession
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetlibraries_GET tests the GET /api/v1/libraries endpoint
func TestGetlibraries_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getLibraries
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreatelibrary_POST tests the POST /api/v1/libraries endpoint
func TestCreatelibrary_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestScanlibrarypaths_POST tests the POST /api/v1/libraries/scan endpoint
func TestScanlibrarypaths_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for scanLibraryPaths
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGeticonnames_GET tests the GET /api/v1/icons endpoint
func TestGeticonnames_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getIconNames
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSavesvgicon_POST tests the POST /api/v1/icons endpoint
func TestSavesvgicon_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for saveSvgIcon
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSavebatchsvgicons_POST tests the POST /api/v1/icons/batch endpoint
func TestSavebatchsvgicons_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for saveBatchSvgIcons
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUploadfile_POST tests the POST /api/v1/files/upload endpoint
func TestUploadfile_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for uploadFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUploadfile1_POST tests the POST /api/v1/files/upload/bookdrop endpoint
func TestUploadfile1_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for uploadFile_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestMovefiles_POST tests the POST /api/v1/files/move endpoint
func TestMovefiles_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for moveFiles
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUploadfont_POST tests the POST /api/v1/custom-fonts/upload endpoint
func TestUploadfont_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for uploadFont
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestAttachbookfiles_POST tests the POST /api/v1/books/{targetBookId}/attach-file endpoint
func TestAttachbookfiles_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for attachBookFiles
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRegeneratecovers_POST tests the POST /api/v1/books/{bookId}/regenerate-cover endpoint
func TestRegeneratecovers_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for regenerateCovers
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetmetadatalist_POST tests the POST /api/v1/books/{bookId}/metadata/prospective endpoint
func TestGetmetadatalist_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getMetadataList
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetimages_POST tests the POST /api/v1/books/{bookId}/metadata/covers endpoint
func TestGetimages_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getImages
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUploadcoverfromfile_POST tests the POST /api/v1/books/{bookId}/metadata/cover/upload endpoint
func TestUploadcoverfromfile_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for uploadCoverFromFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUploadcoverfromurl_POST tests the POST /api/v1/books/{bookId}/metadata/cover/from-url endpoint
func TestUploadcoverfromurl_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for uploadCoverFromUrl
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGeneratecustomcover_POST tests the POST /api/v1/books/{bookId}/generate-custom-cover endpoint
func TestGeneratecustomcover_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for generateCustomCover
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetadditionalfiles_GET tests the GET /api/v1/books/{bookId}/files endpoint
func TestGetadditionalfiles_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAdditionalFiles
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUploadadditionalfile_POST tests the POST /api/v1/books/{bookId}/files endpoint
func TestUploadadditionalfile_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for uploadAdditionalFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatereadstatus_POST tests the POST /api/v1/books/status endpoint
func TestUpdatereadstatus_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateReadStatus
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestAddbooktoshelf_POST tests the POST /api/v1/books/shelves endpoint
func TestAddbooktoshelf_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for addBookToShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestResetprogress_POST tests the POST /api/v1/books/reset-progress endpoint
func TestResetprogress_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for resetProgress
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestResetpersonalrating_POST tests the POST /api/v1/books/reset-personal-rating endpoint
func TestResetpersonalrating_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for resetPersonalRating
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRegeneratecovers1_POST tests the POST /api/v1/books/regenerate-covers endpoint
func TestRegeneratecovers1_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for regenerateCovers_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestAddbooktoprogress_POST tests the POST /api/v1/books/progress endpoint
func TestAddbooktoprogress_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for addBookToProgress
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreatephysicalbook_POST tests the POST /api/v1/books/physical endpoint
func TestCreatephysicalbook_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createPhysicalBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRecalculatematchscores_POST tests the POST /api/v1/books/metadata/recalculate-match-scores endpoint
func TestRecalculatematchscores_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for recalculateMatchScores
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletemetadata_POST tests the POST /api/v1/books/metadata/manage/delete endpoint
func TestDeletemetadata_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestMergemetadata_POST tests the POST /api/v1/books/metadata/manage/consolidate endpoint
func TestMergemetadata_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for mergeMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestBulkuploadcover_POST tests the POST /api/v1/books/bulk-upload-cover endpoint
func TestBulkuploadcover_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for bulkUploadCover
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRegeneratecoversforbooks_POST tests the POST /api/v1/books/bulk-regenerate-covers endpoint
func TestRegeneratecoversforbooks_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for regenerateCoversForBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGeneratecustomcoversforbooks_POST tests the POST /api/v1/books/bulk-generate-custom-covers endpoint
func TestGeneratecustomcoversforbooks_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for generateCustomCoversForBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreatebookmark_POST tests the POST /api/v1/bookmarks endpoint
func TestCreatebookmark_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createBookmark
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRescanbookdrop_POST tests the POST /api/v1/bookdrop/rescan endpoint
func TestRescanbookdrop_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for rescanBookdrop
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestFinalizeimport_POST tests the POST /api/v1/bookdrop/imports/finalize endpoint
func TestFinalizeimport_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for finalizeImport
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestExtractfrompattern_POST tests the POST /api/v1/bookdrop/files/extract-pattern endpoint
func TestExtractfrompattern_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for extractFromPattern
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDiscardselectedfiles_POST tests the POST /api/v1/bookdrop/files/discard endpoint
func TestDiscardselectedfiles_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for discardSelectedFiles
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestBulkeditmetadata1_POST tests the POST /api/v1/bookdrop/files/bulk-edit endpoint
func TestBulkeditmetadata1_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for bulkEditMetadata_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreatenote1_POST tests the POST /api/v1/book-notes endpoint
func TestCreatenote1_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createNote_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRegisteruser_POST tests the POST /api/v1/auth/register endpoint
func TestRegisteruser_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for registerUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestRefreshtoken_POST tests the POST /api/v1/auth/refresh endpoint
func TestRefreshtoken_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for refreshToken
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestHandleoidccallback_POST tests the POST /api/v1/auth/mobile/oidc/callback endpoint
func TestHandleoidccallback_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for handleOidcCallback
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestLoginuser_POST tests the POST /api/v1/auth/login endpoint
func TestLoginuser_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for loginUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateannotation_POST tests the POST /api/v1/annotations endpoint
func TestCreateannotation_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createAnnotation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateproposalstatus_POST tests the POST /api/metadata/tasks/{taskId}/proposals/{proposalId}/status endpoint
func TestUpdateproposalstatus_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateProposalStatus
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetallforuser_GET tests the GET /api/magic-shelves endpoint
func TestGetallforuser_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllForUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateupdateshelf_POST tests the POST /api/magic-shelves endpoint
func TestCreateupdateshelf_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createUpdateShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCreateuser1_POST tests the POST /api/koreader/users/create endpoint
func TestCreateuser1_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for createUser_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestAuthenticatedevice_POST tests the POST /api/kobo/{token}/v1/auth/device endpoint
func TestAuthenticatedevice_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for authenticateDevice
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGettests_POST tests the POST /api/kobo/{token}/v1/analytics/gettests endpoint
func TestGettests_POST(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getTests
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteuser1_DELETE tests the DELETE /api/v2/opds-users/{id} endpoint
func TestDeleteuser1_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteUser_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdateuser1_PATCH tests the PATCH /api/v2/opds-users/{id} endpoint
func TestUpdateuser1_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateUser_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSetdefaultemailrecipient_PATCH tests the PATCH /api/v2/email/recipients/{id}/set-default endpoint
func TestSetdefaultemailrecipient_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for setDefaultEmailRecipient
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSetdefaultemailprovider_PATCH tests the PATCH /api/v2/email/providers/{id}/set-default endpoint
func TestSetdefaultemailprovider_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for setDefaultEmailProvider
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestPatchcronconfig_PATCH tests the PATCH /api/v1/tasks/{taskType}/cron endpoint
func TestPatchcronconfig_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for patchCronConfig
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSetfilenamingpattern_PATCH tests the PATCH /api/v1/libraries/{libraryId}/file-naming-pattern endpoint
func TestSetfilenamingpattern_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for setFileNamingPattern
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestUpdatesyncenabled_PATCH tests the PATCH /api/v1/koreader-users/me/sync endpoint
func TestUpdatesyncenabled_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for updateSyncEnabled
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestTogglesyncprogresswithbooklore_PATCH tests the PATCH /api/v1/koreader-users/me/sync-progress-with-booklore endpoint
func TestTogglesyncprogresswithbooklore_PATCH(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for toggleSyncProgressWithBooklore
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcurrentuser1_GET tests the GET /komga/api/v2/users/me endpoint
func TestGetcurrentuser1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getCurrentUser_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetallseries_GET tests the GET /komga/api/v1/series endpoint
func TestGetallseries_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllSeries
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetseries_GET tests the GET /komga/api/v1/series/{seriesId} endpoint
func TestGetseries_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSeries
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetseriesthumbnail_GET tests the GET /komga/api/v1/series/{seriesId}/thumbnail endpoint
func TestGetseriesthumbnail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSeriesThumbnail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetseriesbooks_GET tests the GET /komga/api/v1/series/{seriesId}/books endpoint
func TestGetseriesbooks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSeriesBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetalllibraries_GET tests the GET /komga/api/v1/libraries endpoint
func TestGetalllibraries_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllLibraries
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetlibrary1_GET tests the GET /komga/api/v1/libraries/{libraryId} endpoint
func TestGetlibrary1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getLibrary_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcollections_GET tests the GET /komga/api/v1/collections endpoint
func TestGetcollections_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getCollections
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetallbooks_GET tests the GET /komga/api/v1/books endpoint
func TestGetallbooks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbook_GET tests the GET /komga/api/v1/books/{bookId} endpoint
func TestGetbook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookthumbnail_GET tests the GET /komga/api/v1/books/{bookId}/thumbnail endpoint
func TestGetbookthumbnail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookThumbnail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookpages_GET tests the GET /komga/api/v1/books/{bookId}/pages endpoint
func TestGetbookpages_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookPages
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookpage_GET tests the GET /komga/api/v1/books/{bookId}/pages/{pageNumber} endpoint
func TestGetbookpage_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookPage
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDownloadbook_GET tests the GET /komga/api/v1/books/{bookId}/file endpoint
func TestDownloadbook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for downloadBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetnotesforbook_GET tests the GET /api/v2/book-notes/book/{bookId} endpoint
func TestGetnotesforbook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getNotesForBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetversioninfo_GET tests the GET /api/v1/version endpoint
func TestGetversioninfo_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getVersionInfo
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetchangelogsincecurrent_GET tests the GET /api/v1/version/changelog endpoint
func TestGetchangelogsincecurrent_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getChangelogSinceCurrent
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetallusers_GET tests the GET /api/v1/users endpoint
func TestGetallusers_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllUsers
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetmyself_GET tests the GET /api/v1/users/me endpoint
func TestGetmyself_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getMyself
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGettimelineforweek_GET tests the GET /api/v1/user-stats/timeline endpoint
func TestGettimelineforweek_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getTimelineForWeek
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetreadingspeedforyear_GET tests the GET /api/v1/user-stats/speed endpoint
func TestGetreadingspeedforyear_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getReadingSpeedForYear
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetpeakreadinghours_GET tests the GET /api/v1/user-stats/peak-hours endpoint
func TestGetpeakreadinghours_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getPeakReadingHours
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetheatmapforyear_GET tests the GET /api/v1/user-stats/heatmap endpoint
func TestGetheatmapforyear_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getHeatmapForYear
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetheatmapformonth_GET tests the GET /api/v1/user-stats/heatmap/monthly endpoint
func TestGetheatmapformonth_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getHeatmapForMonth
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetgenrestatistics_GET tests the GET /api/v1/user-stats/genres endpoint
func TestGetgenrestatistics_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getGenreStatistics
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetfavoritereadingdays_GET tests the GET /api/v1/user-stats/favorite-days endpoint
func TestGetfavoritereadingdays_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getFavoriteReadingDays
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcompletiontimeline_GET tests the GET /api/v1/user-stats/completion-timeline endpoint
func TestGetcompletiontimeline_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getCompletionTimeline
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookcompletionheatmap_GET tests the GET /api/v1/user-stats/book-completion-heatmap endpoint
func TestGetbookcompletionheatmap_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookCompletionHeatmap
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetavailabletasks_GET tests the GET /api/v1/tasks endpoint
func TestGetavailabletasks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAvailableTasks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetlatesttasksforeachtype_GET tests the GET /api/v1/tasks/last endpoint
func TestGetlatesttasksforeachtype_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getLatestTasksForEachType
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetshelfbooks_GET tests the GET /api/v1/shelves/{shelfId}/books endpoint
func TestGetshelfbooks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getShelfBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsetupstatus_GET tests the GET /api/v1/setup/status endpoint
func TestGetsetupstatus_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSetupStatus
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestListbybook_GET tests the GET /api/v1/reviews/book/{bookId} endpoint
func TestListbybook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for listByBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteallbybookid_DELETE tests the DELETE /api/v1/reviews/book/{bookId} endpoint
func TestDeleteallbybookid_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteAllByBookId
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetreadingsessionsforbook_GET tests the GET /api/v1/reading-sessions/book/{bookId} endpoint
func TestGetreadingsessionsforbook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getReadingSessionsForBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetpublicsettings_GET tests the GET /api/v1/public-settings endpoint
func TestGetpublicsettings_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getPublicSettings
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestListpages_GET tests the GET /api/v1/pdf/{bookId}/pages endpoint
func TestListpages_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for listPages
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookinfo_GET tests the GET /api/v1/pdf/{bookId}/info endpoint
func TestGetbookinfo_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookInfo
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetfolders_GET tests the GET /api/v1/path endpoint
func TestGetfolders_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getFolders
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetrootcatalog_GET tests the GET /api/v1/opds endpoint
func TestGetrootcatalog_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getRootCatalog
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDownloadbook1_GET tests the GET /api/v1/opds/{bookId}/download endpoint
func TestDownloadbook1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for downloadBook_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookcover_GET tests the GET /api/v1/opds/{bookId}/cover endpoint
func TestGetbookcover_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookCover
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsurprisefeed_GET tests the GET /api/v1/opds/surprise endpoint
func TestGetsurprisefeed_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSurpriseFeed
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetshelvesnavigation_GET tests the GET /api/v1/opds/shelves endpoint
func TestGetshelvesnavigation_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getShelvesNavigation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetseriesnavigation_GET tests the GET /api/v1/opds/series endpoint
func TestGetseriesnavigation_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSeriesNavigation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsearchdescription_GET tests the GET /api/v1/opds/search.opds endpoint
func TestGetsearchdescription_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSearchDescription
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetrecentbooks_GET tests the GET /api/v1/opds/recent endpoint
func TestGetrecentbooks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getRecentBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetmagicshelvesnavigation_GET tests the GET /api/v1/opds/magic-shelves endpoint
func TestGetmagicshelvesnavigation_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getMagicShelvesNavigation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetlibrariesnavigation_GET tests the GET /api/v1/opds/libraries endpoint
func TestGetlibrariesnavigation_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getLibrariesNavigation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcatalog_GET tests the GET /api/v1/opds/catalog endpoint
func TestGetcatalog_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getCatalog
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetauthorsnavigation_GET tests the GET /api/v1/opds/authors endpoint
func TestGetauthorsnavigation_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAuthorsNavigation
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookdropcover_GET tests the GET /api/v1/media/bookdrop/{bookdropId}/cover endpoint
func TestGetbookdropcover_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookdropCover
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookthumbnail1_GET tests the GET /api/v1/media/book/{bookId}/thumbnail endpoint
func TestGetbookthumbnail1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookThumbnail_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetpdfpage_GET tests the GET /api/v1/media/book/{bookId}/pdf/pages/{pageNumber} endpoint
func TestGetpdfpage_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getPdfPage
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookcover1_GET tests the GET /api/v1/media/book/{bookId}/cover endpoint
func TestGetbookcover1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookCover_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcbxpage_GET tests the GET /api/v1/media/book/{bookId}/cbx/pages/{pageNumber} endpoint
func TestGetcbxpage_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getCbxPage
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbooks_GET tests the GET /api/v1/libraries/{libraryId}/book endpoint
func TestGetbooks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbook1_GET tests the GET /api/v1/libraries/{libraryId}/book/{bookId} endpoint
func TestGetbook1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBook_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsvgiconcontent_GET tests the GET /api/v1/icons/{svgName}/content endpoint
func TestGetsvgiconcontent_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSvgIconContent
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetalliconscontent_GET tests the GET /api/v1/icons/all/content endpoint
func TestGetalliconscontent_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAllIconsContent
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetping_GET tests the GET /api/v1/healthcheck endpoint
func TestGetping_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getPing
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookinfo1_GET tests the GET /api/v1/epub/{bookId}/info endpoint
func TestGetbookinfo1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookInfo_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetfile_GET tests the GET /api/v1/epub/{bookId}/file endpoint
func TestGetfile_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetuserfonts_GET tests the GET /api/v1/custom-fonts endpoint
func TestGetuserfonts_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getUserFonts
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetfontfile_GET tests the GET /api/v1/custom-fonts/{fontId}/file endpoint
func TestGetfontfile_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getFontFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestListpages1_GET tests the GET /api/v1/cbx/{bookId}/pages endpoint
func TestListpages1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for listPages_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetpageinfo_GET tests the GET /api/v1/cbx/{bookId}/page-info endpoint
func TestGetpageinfo_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getPageInfo
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbooks1_GET tests the GET /api/v1/books endpoint
func TestGetbooks1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBooks_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletebooks_DELETE tests the DELETE /api/v1/books endpoint
func TestDeletebooks_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetrecommendations_GET tests the GET /api/v1/books/{id}/recommendations endpoint
func TestGetrecommendations_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getRecommendations
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbook2_GET tests the GET /api/v1/books/{bookId} endpoint
func TestGetbook2_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBook_2
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDownloadadditionalfile_GET tests the GET /api/v1/books/{bookId}/files/{fileId}/download endpoint
func TestDownloadadditionalfile_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for downloadAdditionalFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDownloadbook2_GET tests the GET /api/v1/books/{bookId}/download endpoint
func TestDownloadbook2_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for downloadBook_2
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDownloadallbookfiles_GET tests the GET /api/v1/books/{bookId}/download-all endpoint
func TestDownloadallbookfiles_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for downloadAllBookFiles
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookcontent_GET tests the GET /api/v1/books/{bookId}/content endpoint
func TestGetbookcontent_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookContent
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcomicinfometadata_GET tests the GET /api/v1/books/{bookId}/cbx/metadata/comicinfo endpoint
func TestGetcomicinfometadata_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getComicInfoMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbooksbyids_GET tests the GET /api/v1/books/batch endpoint
func TestGetbooksbyids_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBooksByIds
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookmarksforbook_GET tests the GET /api/v1/bookmarks/book/{bookId} endpoint
func TestGetbookmarksforbook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookmarksForBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetsummary_GET tests the GET /api/v1/bookdrop/notification endpoint
func TestGetsummary_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getSummary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetfilesbystatus_GET tests the GET /api/v1/bookdrop/files endpoint
func TestGetfilesbystatus_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getFilesByStatus
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetnotesforbook1_GET tests the GET /api/v1/book-notes/book/{bookId} endpoint
func TestGetnotesforbook1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getNotesForBook_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetauthorsbybookid_GET tests the GET /api/v1/authors/book/{bookId} endpoint
func TestGetauthorsbybookid_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAuthorsByBookId
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestLoginremote_GET tests the GET /api/v1/auth/remote endpoint
func TestLoginremote_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for loginRemote
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestHandleoidcredirect_GET tests the GET /api/v1/auth/mobile/oidc/redirect endpoint
func TestHandleoidcredirect_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for handleOidcRedirect
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestStreamtrack_GET tests the GET /api/v1/audiobook/{bookId}/track/{trackIndex}/stream endpoint
func TestStreamtrack_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for streamTrack
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestStreamaudiobook_GET tests the GET /api/v1/audiobook/{bookId}/stream endpoint
func TestStreamaudiobook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for streamAudiobook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetaudiobookinfo_GET tests the GET /api/v1/audiobook/{bookId}/info endpoint
func TestGetaudiobookinfo_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAudiobookInfo
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetembeddedcover_GET tests the GET /api/v1/audiobook/{bookId}/cover endpoint
func TestGetembeddedcover_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getEmbeddedCover
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetannotationsforbook_GET tests the GET /api/v1/annotations/book/{bookId} endpoint
func TestGetannotationsforbook_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getAnnotationsForBook
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetshelves_GET tests the GET /api/mobile/v1/shelves endpoint
func TestGetshelves_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getShelves
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetmagicshelves_GET tests the GET /api/mobile/v1/shelves/magic endpoint
func TestGetmagicshelves_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getMagicShelves
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetlibraries1_GET tests the GET /api/mobile/v1/libraries endpoint
func TestGetlibraries1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getLibraries_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbooks2_GET tests the GET /api/mobile/v1/books endpoint
func TestGetbooks2_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBooks_2
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookdetail_GET tests the GET /api/mobile/v1/books/{bookId} endpoint
func TestGetbookdetail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookDetail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSearchbooks_GET tests the GET /api/mobile/v1/books/search endpoint
func TestSearchbooks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for searchBooks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetrecentlyadded_GET tests the GET /api/mobile/v1/books/recently-added endpoint
func TestGetrecentlyadded_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getRecentlyAdded
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbooksbymagicshelf_GET tests the GET /api/mobile/v1/books/magic-shelf/{magicShelfId} endpoint
func TestGetbooksbymagicshelf_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBooksByMagicShelf
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetcontinuereading_GET tests the GET /api/mobile/v1/books/continue-reading endpoint
func TestGetcontinuereading_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getContinueReading
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGettaskwithproposals_GET tests the GET /api/metadata/tasks/{taskId} endpoint
func TestGettaskwithproposals_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getTaskWithProposals
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletetask_DELETE tests the DELETE /api/metadata/tasks/{taskId} endpoint
func TestDeletetask_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteTask
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetactivetasks_GET tests the GET /api/metadata/tasks/active endpoint
func TestGetactivetasks_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getActiveTasks
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetshelf1_GET tests the GET /api/magic-shelves/{id} endpoint
func TestGetshelf1_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getShelf_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteshelf1_DELETE tests the DELETE /api/magic-shelves/{id} endpoint
func TestDeleteshelf1_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteShelf_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestAuthorizeuser_GET tests the GET /api/koreader/users/auth endpoint
func TestAuthorizeuser_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for authorizeUser
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetprogress_GET tests the GET /api/koreader/syncs/progress/{bookHash} endpoint
func TestGetprogress_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getProgress
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetbookmetadata_GET tests the GET /api/kobo/{token}/v1/library/{bookId}/metadata endpoint
func TestGetbookmetadata_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getBookMetadata
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestSynclibrary_GET tests the GET /api/kobo/{token}/v1/library/sync endpoint
func TestSynclibrary_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for syncLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestInitialization_GET tests the GET /api/kobo/{token}/v1/initialization endpoint
func TestInitialization_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for initialization
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetversionedgreythumbnail_GET tests the GET /api/kobo/{token}/v1/books/{imageId}/{version}/thumbnail/{width}/{height}/{quality}/{isGreyscale}/image.jpg endpoint
func TestGetversionedgreythumbnail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getVersionedGreyThumbnail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetversionedthumbnail_GET tests the GET /api/kobo/{token}/v1/books/{imageId}/{version}/thumbnail/{width}/{height}/false/image.jpg endpoint
func TestGetversionedthumbnail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getVersionedThumbnail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetgreythumbnail_GET tests the GET /api/kobo/{token}/v1/books/{imageId}/thumbnail/{width}/{height}/{quality}/{isGreyscale}/image.jpg endpoint
func TestGetgreythumbnail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getGreyThumbnail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestGetthumbnail_GET tests the GET /api/kobo/{token}/v1/books/{imageId}/thumbnail/{width}/{height}/false/image.jpg endpoint
func TestGetthumbnail_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for getThumbnail
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDownloadbook3_GET tests the GET /api/kobo/{token}/v1/books/{bookId}/download endpoint
func TestDownloadbook3_GET(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for downloadBook_3
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestCanceltask_DELETE tests the DELETE /api/v1/tasks/{taskId}/cancel endpoint
func TestCanceltask_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for cancelTask
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDelete_DELETE tests the DELETE /api/v1/reviews/{id} endpoint
func TestDelete_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for delete
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletesvgicon_DELETE tests the DELETE /api/v1/icons/{svgName} endpoint
func TestDeletesvgicon_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteSvgIcon
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletefont_DELETE tests the DELETE /api/v1/custom-fonts/{fontId} endpoint
func TestDeletefont_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteFont
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeleteadditionalfile_DELETE tests the DELETE /api/v1/books/{bookId}/files/{fileId} endpoint
func TestDeleteadditionalfile_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteAdditionalFile
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletenote1_DELETE tests the DELETE /api/v1/book-notes/{noteId} endpoint
func TestDeletenote1_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteNote_1
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// TestDeletebookfromlibrary_DELETE tests the DELETE /api/kobo/{token}/v1/library/{bookId} endpoint
func TestDeletebookfromlibrary_DELETE(t *testing.T) {
	ctx := context.Background()
	client, err := setupTestClient()
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// TODO: Set up parameters for deleteBookFromLibrary
	// TODO: Add assertions to validate response

	_ = client
	_ = ctx
}

// setupTestClient creates a test client using environment variables
func setupTestClient() (*ClientWithResponses, error) {
	server := os.Getenv("BOOKLORE_SERVER")
	if server == "" {
		return nil, fmt.Errorf("BOOKLORE_SERVER environment variable not set")
	}

	client, err := NewClientWithResponses(server)
	if err != nil {
		return nil, err
	}

	return client, nil
}
