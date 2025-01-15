package schemas

type LogoutFromService struct {
	ServiceName string `json:"service_name" binding:"required"`
}
