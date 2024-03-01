package model

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// EntityStatMixin 定义了用于实体状态跟踪的混入。
type EntityStatMixin struct {
	mixin.Schema
}

// DeletedNo 和 DeletedYes 分别代表未删除和已删除的状态值。
const (
	DeletedNo  uint = 0
	DeletedYes uint = 1
)

// Fields 方法定义了混入中包含的字段。
func (EntityStatMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().Comment("默认主键").Annotations(entsql.WithComments(true)),
		field.Time("created_at").Comment("创建时间").Annotations(entsql.WithComments(true)).Immutable().Default(time.Now),
		field.Time("updated_at").Comment("更新时间").Annotations(entsql.WithComments(true)).Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Comment("逻辑删除时间").Annotations(entsql.WithComments(true)).Optional(),
		field.Uint("_delete_").Comment("逻辑标识。0 未删除, 1 已删除").Annotations(entsql.WithComments(true)).
			Default(DeletedNo),
	}
}

// Hooks 定义通用的hook.
//
// 注意: 需要在对应的shema client初始化之前引入相关schema ent下的runtime包进行hook注册:
//
//	import(
//	 _ "rd.kakarot.com/project_a/app/app1/internal/data/schema1/ent/runtime"
//	)
//
// Hooks 方法定义了混入中包含的钩子函数。
func (EntityStatMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		// 当 _delete_ 被设置成已删除状态值时，自动更新 deleted_at 字段。
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				// 判断操作是否为更新操作
				if m.Op().Is(ent.OpUpdate) || m.Op().Is(ent.OpUpdate) {
					// 检查字段是否存在并且是 _delete_
					if dv, yes := m.Field("_delete_"); yes {
						// 如果 _delete_ 的值为 DeletedYes，则设置 deleted_at 字段为当前时间
						if dv.(uint) == DeletedYes {
							err := m.SetField("deleted_at", time.Now())
							if err != nil {
								// 如果设置字段的过程中发生错误，则抛出异常
								panic(err)
							}
						}
					}
				}
				// 继续执行下一个 Mutator
				return next.Mutate(ctx, m)
			})
		},
	}
}
