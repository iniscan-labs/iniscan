package custom

import (
	"runtime"

	"github.com/iniscan-labs/iniscan/models"
	"github.com/mss-boot-io/mss-boot/pkg/migration"
	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1735266294177Migrate)
}

func _1735266294177Migrate(db *gorm.DB, version string) error {
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
			new(models.BalanceChange),
			new(models.Block),
			new(models.Cert),
			new(models.Config),
			new(models.Contract),
			new(models.ContractVerifyJob),
			new(models.EventInfo),
			new(models.FunctionInfo),
			new(models.InternalTransaction),
			new(models.OrderSpentChange),
			new(models.Provider),
			new(models.ProviderReward),
			new(models.Validator),
			new(models.ValidatorReward),
			new(models.VoteDetail),
			new(models.Punish),
			new(models.TokenListMaintain),
			new(models.TokenListStat),
			new(models.TokenPriceHistory),
			new(models.TokenTransfer),
			new(models.TransactionLogs),
			new(models.UbicOrders),
		)
		if err != nil {
			return err
		}

		return migration.Migrate.CreateVersion(tx, version)
	})
}
