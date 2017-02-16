// Copyright 2016 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backupapi

import "time"

type ServiceStatus struct {
	// RecentBackup is status of the most recent backup created by
	// the backup service
	RecentBackup *BackupStatus `json:"recentBackup,omitempty"`

	// Backups is the totoal number of existing backups
	Backups int `json:"backups"`
}

type BackupStatus struct {
	// Size is the size of the backup.
	Size int64 `json:"size"`

	// Version is the version of the backup cluster.
	Version string `json:"version"`

	// TimeTook is the total time took to create the backup.
	TimeTook time.Duration `json:"timeTook"`
}
