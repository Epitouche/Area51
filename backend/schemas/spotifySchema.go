package schemas

type SpotifyAction string
const (
	SpotifyAddTrackAction SpotifyAction = "add_track_action"
)

type SpotifyReaction string
const (
	SpotifyAddTrackReaction SpotifyReaction = "add_track_reaction"
)

type SpotifyResponseToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type SpotifyUserInfo struct {
	Id           string `json:"id" gorm:"primaryKey"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name"`
}

type SpotifyActionOptions struct {
	Playlist string `json:"playlist"`
	NbSongs uint64 `json:"nbSongs"`
	IsOld   bool `json:"is_old"`
}

type SpotifyActionOptionsInfo struct {
	PlaylistId string `json:"playlist_id"`
	NbSongs string `json:"nbSongs"`
}

type SpotifyTracksInfos struct {
	Total uint64 `json:"total"`
}

type SpotifyPlaylistInfos struct {
	Tracks SpotifyTracksInfos `json:"tracks"`
}