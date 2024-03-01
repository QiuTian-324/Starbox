package schema

import (
	"BuzzBox/pkg/model"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User 表示用户实体的模式定义。
type User struct {
	ent.Schema
}

// Annotations 返回用户模式的注解。
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// 将表名设置为 "user"。
		entsql.Annotation{
			Table: "user",
		},
	}
}

// Indexes 返回用户模式的索引定义。
func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 创建一个名为 "mobile" 的唯一索引。
		index.Fields("mobile").
			Unique(),
	}
}

// Mixin 返回用户模式的混合器定义。
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// 使用实体状态混合器。
		model.EntityStatMixin{},
	}
}

// Fields 返回用户模式的字段定义。
func (User) Fields() []ent.Field {
	return []ent.Field{
		// 用户名字段，类型为字符串。
		field.String("username").Comment("用户名").Annotations(entsql.WithComments(true)),
		// 手机号字段，类型为字符串。
		field.String("mobile").Comment("手机号").Annotations(entsql.WithComments(true)),
		// 头像URL字段，类型为字符串。
		field.String("avatar").Comment("头像url").Annotations(entsql.WithComments(true)),
		// 密码字段，类型为字符串。
		field.String("password").Comment("密码").Annotations(entsql.WithComments(true)),
	}
}

// Edges 返回用户模式的边定义。
func (User) Edges() []ent.Edge {
	return nil
}
