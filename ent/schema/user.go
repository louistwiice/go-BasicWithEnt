package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Match(regexp.MustCompile(`[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)).Unique(),
		field.String("username").Match(regexp.MustCompile(`^[a-zA-Z0-9_.]*$`)).Unique(),
		field.String("first_name").MaxLen(100),
		field.String("last_name").MaxLen(255),
		field.Text("password"),
		field.Bool("is_active").Default(false),
		field.Bool("is_staff").Default(false),
		field.Bool("is_superuser").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("last_authentication_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
