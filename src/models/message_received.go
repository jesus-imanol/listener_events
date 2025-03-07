package models
type ReceivedMessage struct {
    Id              int32  `json:"id"`
    FullName        string `json:"full_name"`
    Email           string `json:"email"`
    PasswordHash    string `json:"password_hash"`
    Gender          string `json:"gender"`
    MatchPreference string `json:"match_preference"`
    City            string `json:"city"`
    State           string `json:"state"`
    Interests       string `json:"interests"`
    StatusMessage   string `json:"status_message"`
    ProfilePicture  string `json:"profile_picture"`
}