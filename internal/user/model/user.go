package model

// profile terkait jurnal dan publikasi
// pr, supabase harus sama dengan json name
type User struct {
	Id        int    `json:"id,omitempty"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"passwordHash"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type UserCredential struct {
	Username string `json:"username" validate:"required,max=30"` // big int di postgresql
	Password string `json:"password" validate:"required"`
}
