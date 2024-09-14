package main

import (
	"fmt"
	"strings"
)

// Define bitmask indices for
// in bitmask which index denotes which type of permission.
const (
	LogsViewIndex               = 0
	LogsDeleteIndex             = 1
	EggchangerManageIndex       = 2
	ControlConsoleIndex         = 3
	ControlStartIndex           = 4
	ControlStopIndex            = 5
	ControlRestartIndex         = 6
	UserCreateIndex             = 7
	UserReadIndex               = 8
	UserUpdateIndex             = 9
	UserDeleteIndex             = 10
	FileCreateIndex             = 11
	FileReadIndex               = 12
	FileReadContentIndex        = 13
	FileUpdateIndex             = 14
	FileDeleteIndex             = 15
	FileArchiveIndex            = 16
	FileSFTPIndex               = 17
	FileDownloadIndex           = 18
	BackupCreateIndex           = 19
	BackupReadIndex             = 20
	BackupDeleteIndex           = 21
	BackupDownloadIndex         = 22
	BackupRestoreIndex          = 23
	DatabaseBackupCreateIndex   = 24
	DatabaseBackupReadIndex     = 25
	DatabaseBackupUpdateIndex   = 26
	DatabaseBackupDeleteIndex   = 27
	DatabaseBackupDownloadIndex = 28
	DatabaseBackupRestoreIndex  = 29
	AllocationReadIndex         = 30
	AllocationCreateIndex       = 31
	AllocationUpdateIndex       = 32
	AllocationDeleteIndex       = 33
	StartupReadIndex            = 34
	StartupUpdateIndex          = 35
	StartupDockerImageIndex     = 36
	DatabaseCreateIndex         = 37
	DatabaseReadIndex           = 38
	DatabaseUpdateIndex         = 39
	DatabaseDeleteIndex         = 40
	DatabaseViewPasswordIndex   = 41
	ServerImporterAccessIndex   = 42
	ScheduleCreateIndex         = 43
	ScheduleReadIndex           = 44
	ScheduleUpdateIndex         = 45
	ScheduleDeleteIndex         = 46
	SettingsRenameIndex         = 47
	SettingsReinstallIndex      = 48
	SettingsUpdateIndex         = 49
	SubdomainManageIndex        = 50
	ModsManageIndex             = 51
	TransferManageIndex         = 52
	PluginsManageIndex          = 53
	ActivityReadIndex           = 54
)

type Permissions struct {
	Logs           LogsPermission           `json:"logs"`
	Eggchanger     EggchangerPermission     `json:"eggchanger"`
	Control        ControlPermission        `json:"control"`
	User           UserPermission           `json:"user"`
	File           FilePermission           `json:"file"`
	Backup         BackupPermission         `json:"backup"`
	DatabaseBackup DatabaseBackupPermission `json:"database_backup"`
	Allocation     AllocationPermission     `json:"allocation"`
	Startup        StartupPermission        `json:"startup"`
	Database       DatabasePermission       `json:"database"`
	ServerImporter ServerImporterPermission `json:"serverimporter"`
	Schedule       SchedulePermission       `json:"schedule"`
	Settings       SettingsPermission       `json:"settings"`
	Subdomain      SubdomainPermission      `json:"subdomain"`
	Mods           ModsPermission           `json:"mods"`
	Transfer       TransferPermission       `json:"transfer"`
	Plugins        PluginsPermission        `json:"plugins"`
	Activity       ActivityPermission       `json:"activity"`
}

type StaffPermission struct {
	Manage bool `json:"manage"`
}

type StartupPermission struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type LogsPermission struct {
	View   bool `json:"view"`
	Delete bool `json:"delete"`
}

type EggchangerPermission struct {
	Manage bool `json:"manage"`
}

type ControlPermission struct {
	Console bool `json:"console"`
	Start   bool `json:"start"`
	Stop    bool `json:"stop"`
	Restart bool `json:"restart"`
}

type UserPermission struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type FilePermission struct {
	Create      bool `json:"create"`
	Read        bool `json:"read"`
	ReadContent bool `json:"read_content"`
	Update      bool `json:"update"`
	Delete      bool `json:"delete"`
	Archive     bool `json:"archive"`
	SFTP        bool `json:"sftp"`
	Download    bool `json:"download"`
}

type BackupPermission struct {
	Create   bool `json:"create"`
	Read     bool `json:"read"`
	Delete   bool `json:"delete"`
	Download bool `json:"download"`
	Restore  bool `json:"restore"`
}

type DatabaseBackupPermission struct {
	Create   bool `json:"create"`
	Read     bool `json:"read"`
	Update   bool `json:"update"`
	Delete   bool `json:"delete"`
	Download bool `json:"download"`
	Restore  bool `json:"restore"`
}

type AllocationPermission struct {
	Read   bool `json:"read"`
	Create bool `json:"create"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type DatabasePermission struct {
	Create       bool `json:"create"`
	Read         bool `json:"read"`
	Update       bool `json:"update"`
	Delete       bool `json:"delete"`
	ViewPassword bool `json:"viewpassword"`
}

type ServerImporterPermission struct {
	Access bool `json:"access"`
}

type SchedulePermission struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type SettingsPermission struct {
	Rename    bool `json:"rename"`
	Reinstall bool `json:"reinstall"`
	Update    bool `json:"update"`
}

type SubdomainPermission struct {
	Manage bool `json:"manage"`
}

type ModsPermission struct {
	Manage bool `json:"manage"`
}

type TransferPermission struct {
	Manage bool `json:"manage"`
}

type PluginsPermission struct {
	Manage bool `json:"manage"`
}

type ActivityPermission struct {
	Read bool `json:"read"`
}

func setBitMask(bitsmask []string, index int, value bool) {
	bitsmask[index] = BoolToBit(value)
}

func CreateBitmaskForPermissions(permissions Permissions) []string {
	var bitmask []string = make([]string, 60)

	fmt.Println("length of bitmask", len(bitmask))

	// Set bits for Logs permissions
	setBitMask(bitmask, LogsViewIndex, permissions.Logs.View)
	setBitMask(bitmask, LogsDeleteIndex, permissions.Logs.Delete)

	// Set bits for Eggchanger permissions
	setBitMask(bitmask, EggchangerManageIndex, permissions.Eggchanger.Manage)

	// Set bits for Control permissions
	setBitMask(bitmask, ControlConsoleIndex, permissions.Control.Console)
	setBitMask(bitmask, ControlStartIndex, permissions.Control.Start)
	setBitMask(bitmask, ControlStopIndex, permissions.Control.Stop)
	setBitMask(bitmask, ControlRestartIndex, permissions.Control.Restart)

	// Set bits for User permissions
	setBitMask(bitmask, UserCreateIndex, permissions.User.Create)
	setBitMask(bitmask, UserReadIndex, permissions.User.Read)
	setBitMask(bitmask, UserUpdateIndex, permissions.User.Update)
	setBitMask(bitmask, UserDeleteIndex, permissions.User.Delete)

	// Set bits for File permissions
	setBitMask(bitmask, FileCreateIndex, permissions.File.Create)
	setBitMask(bitmask, FileReadIndex, permissions.File.Read)
	setBitMask(bitmask, FileReadContentIndex, permissions.File.ReadContent)
	setBitMask(bitmask, FileUpdateIndex, permissions.File.Update)
	setBitMask(bitmask, FileDeleteIndex, permissions.File.Delete)
	setBitMask(bitmask, FileArchiveIndex, permissions.File.Archive)
	setBitMask(bitmask, FileSFTPIndex, permissions.File.SFTP)
	setBitMask(bitmask, FileDownloadIndex, permissions.File.Download)

	// Set bits for Backup permissions
	setBitMask(bitmask, BackupCreateIndex, permissions.Backup.Create)
	setBitMask(bitmask, BackupReadIndex, permissions.Backup.Read)
	setBitMask(bitmask, BackupDeleteIndex, permissions.Backup.Delete)
	setBitMask(bitmask, BackupDownloadIndex, permissions.Backup.Download)
	setBitMask(bitmask, BackupRestoreIndex, permissions.Backup.Restore)

	// Set bits for Database Backup permissions
	setBitMask(bitmask, DatabaseBackupCreateIndex, permissions.DatabaseBackup.Create)
	setBitMask(bitmask, DatabaseBackupReadIndex, permissions.DatabaseBackup.Read)
	setBitMask(bitmask, DatabaseBackupUpdateIndex, permissions.DatabaseBackup.Update)
	setBitMask(bitmask, DatabaseBackupDeleteIndex, permissions.DatabaseBackup.Delete)
	setBitMask(bitmask, DatabaseBackupDownloadIndex, permissions.DatabaseBackup.Download)
	setBitMask(bitmask, DatabaseBackupRestoreIndex, permissions.DatabaseBackup.Restore)

	// Set bits for Allocation permissions
	setBitMask(bitmask, AllocationReadIndex, permissions.Allocation.Read)
	setBitMask(bitmask, AllocationCreateIndex, permissions.Allocation.Create)
	setBitMask(bitmask, AllocationUpdateIndex, permissions.Allocation.Update)
	setBitMask(bitmask, AllocationDeleteIndex, permissions.Allocation.Delete)

	// Set bits for Startup permissions
	setBitMask(bitmask, StartupReadIndex, permissions.Startup.Read)
	setBitMask(bitmask, StartupUpdateIndex, permissions.Startup.Update)
	setBitMask(bitmask, StartupDockerImageIndex, permissions.Startup.Create) // Assuming DockerImage is part of Create for now

	// Set bits for Database permissions
	setBitMask(bitmask, DatabaseCreateIndex, permissions.Database.Create)
	setBitMask(bitmask, DatabaseReadIndex, permissions.Database.Read)
	setBitMask(bitmask, DatabaseUpdateIndex, permissions.Database.Update)
	setBitMask(bitmask, DatabaseDeleteIndex, permissions.Database.Delete)
	setBitMask(bitmask, DatabaseViewPasswordIndex, permissions.Database.ViewPassword)

	// Set bits for Server Importer permissions
	setBitMask(bitmask, ServerImporterAccessIndex, permissions.ServerImporter.Access)

	// Set bits for Schedule permissions
	setBitMask(bitmask, ScheduleCreateIndex, permissions.Schedule.Create)
	setBitMask(bitmask, ScheduleReadIndex, permissions.Schedule.Read)
	setBitMask(bitmask, ScheduleUpdateIndex, permissions.Schedule.Update)
	setBitMask(bitmask, ScheduleDeleteIndex, permissions.Schedule.Delete)

	// Set bits for Settings permissions
	setBitMask(bitmask, SettingsRenameIndex, permissions.Settings.Rename)
	setBitMask(bitmask, SettingsReinstallIndex, permissions.Settings.Reinstall)
	setBitMask(bitmask, SettingsUpdateIndex, permissions.Settings.Update)

	// Set bits for Subdomain permissions
	setBitMask(bitmask, SubdomainManageIndex, permissions.Subdomain.Manage)

	// Set bits for Mods permissions
	setBitMask(bitmask, ModsManageIndex, permissions.Mods.Manage)

	// Set bits for Transfer permissions
	setBitMask(bitmask, TransferManageIndex, permissions.Transfer.Manage)

	// Set bits for Plugins permissions
	setBitMask(bitmask, PluginsManageIndex, permissions.Plugins.Manage)

	// Set bits for Activity permissions
	setBitMask(bitmask, ActivityReadIndex, permissions.Activity.Read)

	// For reserved bits, mark them as false for now
	// It can be used for future permissions added
	setBitMask(bitmask, 55, false)
	setBitMask(bitmask, 56, false)
	setBitMask(bitmask, 57, false)
	setBitMask(bitmask, 58, false)
	setBitMask(bitmask, 59, false)

	return bitmask
}

// func GetPermissionsFromBitmask(bitmask []string) Permissions {
// 	permissions := Permissions{}

// 	// check every permissions index and set the corresponding permission

// 	// logs permissions
// 	permissions.Logs.View = BitToBool(bitmask[LogsViewIndex])
// 	permissions.Logs.Delete = BitToBool(bitmask[LogsDeleteIndex])

// 	permissions.Eggchanger.Manage = BitToBool(bitmask[EggchangerManageIndex])
// }

func GetPermissionsFromBitmask(bitmask []string) Permissions {
	permissions := Permissions{}

	// Logs permissions
	permissions.Logs.View = BitToBool(bitmask[LogsViewIndex])
	permissions.Logs.Delete = BitToBool(bitmask[LogsDeleteIndex])

	// Eggchanger permissions
	permissions.Eggchanger.Manage = BitToBool(bitmask[EggchangerManageIndex])

	// Control permissions
	permissions.Control.Console = BitToBool(bitmask[ControlConsoleIndex])
	permissions.Control.Start = BitToBool(bitmask[ControlStartIndex])
	permissions.Control.Stop = BitToBool(bitmask[ControlStopIndex])
	permissions.Control.Restart = BitToBool(bitmask[ControlRestartIndex])

	// User permissions
	permissions.User.Create = BitToBool(bitmask[UserCreateIndex])
	permissions.User.Read = BitToBool(bitmask[UserReadIndex])
	permissions.User.Update = BitToBool(bitmask[UserUpdateIndex])
	permissions.User.Delete = BitToBool(bitmask[UserDeleteIndex])

	// File permissions
	permissions.File.Create = BitToBool(bitmask[FileCreateIndex])
	permissions.File.Read = BitToBool(bitmask[FileReadIndex])
	permissions.File.ReadContent = BitToBool(bitmask[FileReadContentIndex])
	permissions.File.Update = BitToBool(bitmask[FileUpdateIndex])
	permissions.File.Delete = BitToBool(bitmask[FileDeleteIndex])
	permissions.File.Archive = BitToBool(bitmask[FileArchiveIndex])
	permissions.File.SFTP = BitToBool(bitmask[FileSFTPIndex])
	permissions.File.Download = BitToBool(bitmask[FileDownloadIndex])

	// Backup permissions
	permissions.Backup.Create = BitToBool(bitmask[BackupCreateIndex])
	permissions.Backup.Read = BitToBool(bitmask[BackupReadIndex])
	permissions.Backup.Delete = BitToBool(bitmask[BackupDeleteIndex])
	permissions.Backup.Download = BitToBool(bitmask[BackupDownloadIndex])
	permissions.Backup.Restore = BitToBool(bitmask[BackupRestoreIndex])

	// Database Backup permissions
	permissions.DatabaseBackup.Create = BitToBool(bitmask[DatabaseBackupCreateIndex])
	permissions.DatabaseBackup.Read = BitToBool(bitmask[DatabaseBackupReadIndex])
	permissions.DatabaseBackup.Update = BitToBool(bitmask[DatabaseBackupUpdateIndex])
	permissions.DatabaseBackup.Delete = BitToBool(bitmask[DatabaseBackupDeleteIndex])
	permissions.DatabaseBackup.Download = BitToBool(bitmask[DatabaseBackupDownloadIndex])
	permissions.DatabaseBackup.Restore = BitToBool(bitmask[DatabaseBackupRestoreIndex])

	// Allocation permissions
	permissions.Allocation.Read = BitToBool(bitmask[AllocationReadIndex])
	permissions.Allocation.Create = BitToBool(bitmask[AllocationCreateIndex])
	permissions.Allocation.Update = BitToBool(bitmask[AllocationUpdateIndex])
	permissions.Allocation.Delete = BitToBool(bitmask[AllocationDeleteIndex])

	// Startup permissions
	permissions.Startup.Read = BitToBool(bitmask[StartupReadIndex])
	permissions.Startup.Update = BitToBool(bitmask[StartupUpdateIndex])
	permissions.Startup.Create = BitToBool(bitmask[StartupDockerImageIndex])

	// Database permissions
	permissions.Database.Create = BitToBool(bitmask[DatabaseCreateIndex])
	permissions.Database.Read = BitToBool(bitmask[DatabaseReadIndex])
	permissions.Database.Update = BitToBool(bitmask[DatabaseUpdateIndex])
	permissions.Database.Delete = BitToBool(bitmask[DatabaseDeleteIndex])
	permissions.Database.ViewPassword = BitToBool(bitmask[DatabaseViewPasswordIndex])

	// Server Importer permissions
	permissions.ServerImporter.Access = BitToBool(bitmask[ServerImporterAccessIndex])

	// Schedule permissions
	permissions.Schedule.Create = BitToBool(bitmask[ScheduleCreateIndex])
	permissions.Schedule.Read = BitToBool(bitmask[ScheduleReadIndex])
	permissions.Schedule.Update = BitToBool(bitmask[ScheduleUpdateIndex])
	permissions.Schedule.Delete = BitToBool(bitmask[ScheduleDeleteIndex])

	// Settings permissions
	permissions.Settings.Rename = BitToBool(bitmask[SettingsRenameIndex])
	permissions.Settings.Reinstall = BitToBool(bitmask[SettingsReinstallIndex])
	permissions.Settings.Update = BitToBool(bitmask[SettingsUpdateIndex])

	// Subdomain permissions
	permissions.Subdomain.Manage = BitToBool(bitmask[SubdomainManageIndex])

	// Mods permissions
	permissions.Mods.Manage = BitToBool(bitmask[ModsManageIndex])

	// Transfer permissions
	permissions.Transfer.Manage = BitToBool(bitmask[TransferManageIndex])

	// Plugins permissions
	permissions.Plugins.Manage = BitToBool(bitmask[PluginsManageIndex])

	// Activity permissions
	permissions.Activity.Read = BitToBool(bitmask[ActivityReadIndex])

	return permissions
}

func BoolToBit(b bool) string {
	if b {
		return "1"
	}

	return "0"
}

func BitToBool(bit string) bool {
	if bit == "1" {
		return true
	}

	return false
}

func main() {
	// userId := ""
	// serverId := ""

	Permissions := Permissions{
		Logs: LogsPermission{
			View:   true,
			Delete: true,
		},
		Eggchanger: EggchangerPermission{
			Manage: true,
		},
		Control: ControlPermission{
			Console: true,
			Start:   false,
			Stop:    false,
			Restart: true,
		},
	}

	bitmask := CreateBitmaskForPermissions(Permissions)

	str := strings.Join(bitmask, "")
	fmt.Println(str)

	PermissionsFromBitmask := GetPermissionsFromBitmask(strings.Split(str, ""))
	fmt.Println(PermissionsFromBitmask)
}
