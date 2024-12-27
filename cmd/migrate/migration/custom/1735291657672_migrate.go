package custom

import (
	"runtime"

	"github.com/iniscan-labs/iniscan/models"
	"github.com/mss-boot-io/mss-boot/pkg/migration"
	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1735291657672Migrate)
}

func _1735291657672Migrate(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// TODO: here to write the content to be changed

		// TODO: e.g. modify table field, please delete this code during use
		//err := tx.Migrator().RenameColumn(&models.SysConfig{}, "config_id", "id")
		//if err != nil {
		// 	return err
		//}

		// TODO: e.g. add table structure, please delete this code during use
		err := tx.Migrator().AutoMigrate(
			new(models.Account),
			new(models.AccountBalance),
			new(models.Block),
			new(models.Contract),
			new(models.ContractVerifyJob),
			new(models.EventInfo),
			new(models.FunctionInfo),
			new(models.InternalTransaction),
			new(models.OrderSpentChange),
			new(models.TokenTransfer),
			new(models.Transaction),
			new(models.TransactionLogs),
		)
		if err != nil {
			return err
		}

		return migration.Migrate.CreateVersion(tx, version)
	})
}
