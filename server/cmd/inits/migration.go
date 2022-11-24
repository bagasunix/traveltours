package inits

import (
	"github.com/bagasunix/traveltours/pkg/errors"
	"github.com/bagasunix/traveltours/server/domains/data/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetTables() (tables []interface{}) {

	tables = append(tables, models.NewCountryBuilder().Build())
	tables = append(tables, models.NewProvinceBuilder().Build())
	tables = append(tables, models.NewCityBuilder().Build())
	tables = append(tables, models.NewSubDistrictBuilder().Build())
	tables = append(tables, models.NewVillageBuilder().Build())

	tables = append(tables, models.NewRoleBuilder().Build())
	tables = append(tables, models.NewPermissionBuilder().Build())
	tables = append(tables, models.NewRolePermissionBuilder().Build())

	tables = append(tables, models.NewUserReligiBuilder().Build())
	tables = append(tables, models.NewUserSexBuilder().Build())
	tables = append(tables, models.NewUserBuilder().Build())
	tables = append(tables, models.NewUserDetailsBuilder().Build())

	tables = append(tables, models.NewImageBuilder().Build())
	tables = append(tables, models.NewTourBuilder().Build())
	tables = append(tables, models.NewTourDifficultyBuilder().Build())
	tables = append(tables, models.NewTourReviewBuilder().Build())
	tables = append(tables, models.NewLocationTourBuilder().Build())

	tables = append(tables, models.NewGroupMenuBuilder().Build())
	tables = append(tables, models.NewMenuBuilder().Build())
	tables = append(tables, models.NewSubMenuBuilder().Build())
	tables = append(tables, models.NewAccessMenuBuilder().Build())

	return tables
}
func Migrate(logs *zap.Logger, db *gorm.DB) {
	errors.HandlerWithOSExit(logs, db.AutoMigrate(GetTables()...), "AutoMigrate")
}
