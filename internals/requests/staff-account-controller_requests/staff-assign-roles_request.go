package staffAccountControllerRequests

// StaffAssignRolesRequest — phân roles cho nhân viên.
type StaffAssignRolesRequest struct {
	Roles []string `json:"roles" validate:"omitempty,dive,required,oneof=DOCTOR"`
}
