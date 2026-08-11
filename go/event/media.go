package event

// Media event types (reserved). Producer: media-service.
// Media-service currently uses HTTP S2S without domain event fan-out; constants
// are reserved for future lifecycle events.
const (
	TypeMediaUploadCompletedV1 = "media.upload_completed.v1"
	TypeMediaFileDeletedV1     = "media.file_deleted.v1"
)
