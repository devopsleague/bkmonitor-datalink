// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package es

import (
	"context"
	"sync"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/log"
)

var storageMap map[string]Client

var storageLock *sync.RWMutex

// ReloadStorage
func ReloadStorage(infos map[string]*ESInfo) error {
	newStorageMap := make(map[string]Client)
	for key, info := range infos {
		client, err := NewClient(info)
		if err != nil {
			return err
		}
		newStorageMap[key] = client
	}
	storageLock.Lock()
	defer storageLock.Unlock()

	storageMap = newStorageMap
	for key := range newStorageMap {
		log.Debugf(context.TODO(), "reload storage:%s success", key)
	}
	return nil
}
