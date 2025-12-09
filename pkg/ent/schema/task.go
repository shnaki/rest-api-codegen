package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Uint64("user_id").
			Comment("タスク所有者のユーザーID。").
			// Edge定義でUnique()を使っているので、ここではOptional()のままでOK。
			Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("tasks").
			Unique().
			// このエッジが user_id フィールドを使うことをentに指示する。
			Field("user_id"),
	}
}

// Mixin は再利用可能なスキーマを注入する。
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
