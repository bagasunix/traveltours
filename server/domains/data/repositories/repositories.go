package repositories

import (
	groupmenu "github.com/bagasunix/traveltours/server/domains/data/repositories/group-menu"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/menu"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/permission"
	rolepermission "github.com/bagasunix/traveltours/server/domains/data/repositories/role-permission"
	submenu "github.com/bagasunix/traveltours/server/domains/data/repositories/sub-menu"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/tour"
	tourdificulty "github.com/bagasunix/traveltours/server/domains/data/repositories/tour-dificulty"
	tourreview "github.com/bagasunix/traveltours/server/domains/data/repositories/tour-review"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/user"
	userdetails "github.com/bagasunix/traveltours/server/domains/data/repositories/user-details"
	userreligi "github.com/bagasunix/traveltours/server/domains/data/repositories/user-religi"
	userrole "github.com/bagasunix/traveltours/server/domains/data/repositories/user-role"
	usersex "github.com/bagasunix/traveltours/server/domains/data/repositories/user-sex"
	"github.com/bagasunix/traveltours/server/domains/data/repositories/village"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repositories interface {
	GetTour() tour.Repository
	GetTourReview() tourreview.Repository
	GetTourDificulty() tourdificulty.Repository
	GetUser() user.Repository
	GetUserDetail() userdetails.Repository
	GetUserReligi() userreligi.Repository
	GetUserSex() usersex.Repository
	GetRole() userrole.Repository
	GetVillage() village.Repository
	GetRolePermission() rolepermission.Repository
	GetPermission() permission.Repository
	GetMenu() menu.Repository
	GetGroupMenu() groupmenu.Repository
	GetSubMenu() submenu.Repository
}

type repo struct {
	tour           tour.Repository
	tourReview     tourreview.Repository
	tourdificulty  tourdificulty.Repository
	user           user.Repository
	userDetail     userdetails.Repository
	userReligi     userreligi.Repository
	userSex        usersex.Repository
	role           userrole.Repository
	village        village.Repository
	rolepermission rolepermission.Repository
	permission     permission.Repository
	menu           menu.Repository
	groupmenu      groupmenu.Repository
	submenu        submenu.Repository
}

// GetGroupMenu implements Repositories
func (r *repo) GetGroupMenu() groupmenu.Repository {
	return r.groupmenu
}

// GetMenu implements Repositories
func (r *repo) GetMenu() menu.Repository {
	return r.menu
}

// GetSubMenu implements Repositories
func (r *repo) GetSubMenu() submenu.Repository {
	return r.submenu
}

// GetTourDificulty implements Repositories
func (r *repo) GetTourDificulty() tourdificulty.Repository {
	return r.tourdificulty
}

// GetPermission implements Repositories
func (r *repo) GetPermission() permission.Repository {
	return r.permission
}

// GetRolePermission implements Repositories
func (r *repo) GetRolePermission() rolepermission.Repository {
	return r.rolepermission
}

// GetRole implements Repositories
func (r *repo) GetRole() userrole.Repository {
	return r.role
}

// GetTour implements Repositories
func (r *repo) GetTour() tour.Repository {
	return r.tour
}

// GetTourReview implements Repositories
func (r *repo) GetTourReview() tourreview.Repository {
	return r.tourReview
}

// GetUser implements Repositories
func (r *repo) GetUser() user.Repository {
	return r.user
}

// GetUserDetail implements Repositories
func (r *repo) GetUserDetail() userdetails.Repository {
	return r.userDetail
}

// GetUserReligi implements Repositories
func (r *repo) GetUserReligi() userreligi.Repository {
	return r.userReligi
}

// GetUserSex implements Repositories
func (r *repo) GetUserSex() usersex.Repository {
	return r.userSex
}

// GetVillage implements Repositories
func (r *repo) GetVillage() village.Repository {
	return r.village
}

func New(logger *zap.Logger, db *gorm.DB) Repositories {
	rp := new(repo)
	rp.tour = tour.NewGorm(logger, db)
	rp.tourReview = tourreview.NewGorm(logger, db)
	rp.tourdificulty = tourdificulty.NewGorm(logger, db)
	rp.user = user.NewGorm(logger, db)
	rp.userDetail = userdetails.NewGorm(logger, db)
	rp.userReligi = userreligi.NewGorm(logger, db)
	rp.userSex = usersex.NewGorm(logger, db)
	rp.role = userrole.NewGorm(logger, db)
	rp.village = village.NewGorm(logger, db)
	rp.rolepermission = rolepermission.NewGorm(logger, db)
	rp.permission = permission.NewGorm(logger, db)
	rp.groupmenu = groupmenu.NewGorm(logger, db)
	rp.menu = menu.NewGorm(logger, db)
	rp.submenu = submenu.NewGorm(logger, db)
	return rp
}
