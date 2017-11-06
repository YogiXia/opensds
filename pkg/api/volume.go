// Copyright (c) 2016 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

/*
This module implements a entry into the OpenSDS northbound service.

*/

package api

import (
	"encoding/json"
	"fmt"

	log "github.com/golang/glog"

	"github.com/opensds/opensds/pkg/controller"
	"github.com/opensds/opensds/pkg/db"
	"github.com/opensds/opensds/pkg/model"
	"github.com/opensds/opensds/pkg/utils"

	"github.com/astaxie/beego"
)

type VolumePortal struct {
	beego.Controller
}

func (this *VolumePortal) CreateVolume() {
	var volume = model.VolumeSpec{
		BaseModel: &model.BaseModel{},
	}

	// Unmarshal the request body
	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&volume); err != nil {
		reason := fmt.Sprintf("Parse volume request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Call global controller variable to handle create volume request.
	result, err := controller.Brain.CreateVolume(&volume)
	if err != nil {
		reason := fmt.Sprintf("Create volume failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume created result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusAccepted)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumePortal) ListVolumes() {
	// Call db api module to handle list volumes request.
	result, err := db.C.ListVolumes()
	if err != nil {
		reason := fmt.Sprintf("List volumes failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volumes listed result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumePortal) GetVolume() {
	id := this.Ctx.Input.Param(":volumeId")

	// Call db api module to handle get volume request.
	result, err := db.C.GetVolume(id)
	if err != nil {
		reason := fmt.Sprintf("Get volume failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume showed result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumePortal) UpdateVolume() {
	this.Ctx.Output.SetStatus(StatusNotImplemented)
	return
}

func (this *VolumePortal) DeleteVolume() {
	var volume = model.VolumeSpec{
		BaseModel: &model.BaseModel{},
	}
	volId := this.Ctx.Input.Param(":volumeId")

	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&volume); err != nil {
		reason := fmt.Sprintf("Parse volume request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}
	volume.Id = volId

	// Call global controller variable to handle delete volume request.
	result := controller.Brain.DeleteVolume(&volume)
	if result.Status != "Success" {
		reason := fmt.Sprintf("Delete volume failed: %s", result.GetError())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume deleted result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusAccepted)
	this.Ctx.Output.Body(body)
	return
}

type VolumeAttachmentPortal struct {
	beego.Controller
}

func (this *VolumeAttachmentPortal) CreateVolumeAttachment() {
	var attachment = model.VolumeAttachmentSpec{
		BaseModel: &model.BaseModel{},
	}

	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&attachment); err != nil {
		reason := fmt.Sprintf("Parse volume attachment request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Call global controller variable to handle create volume attachment request.
	result, err := controller.Brain.CreateVolumeAttachment(&attachment)
	if err != nil {
		reason := fmt.Sprintf("Create volume attachment failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume attachment created result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusAccepted)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeAttachmentPortal) ListVolumeAttachments() {
	volId := this.GetString("volumeId")

	result, err := db.C.ListVolumeAttachments(volId)
	if err != nil {
		reason := fmt.Sprintf("List volume attachments failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume attachments listed result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeAttachmentPortal) GetVolumeAttachment() {
	id := this.Ctx.Input.Param(":attachmentId")
	volId := this.GetString("volumeId")

	result, err := db.C.GetVolumeAttachment(volId, id)
	if err != nil {
		reason := fmt.Sprintf("Get volume attachment failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume attachment showed result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeAttachmentPortal) UpdateVolumeAttachment() {
	// TODO Please note that this function has been disabled in controller
	// module.
	var attachment = model.VolumeAttachmentSpec{
		BaseModel: &model.BaseModel{},
	}
	id := this.Ctx.Input.Param(":attachmentId")

	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&attachment); err != nil {
		reason := fmt.Sprintf("Parse volume attachment request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}
	attachment.Id = id

	result, err := controller.Brain.UpdateVolumeAttachment(&attachment)
	if err != nil {
		reason := fmt.Sprintf("Update volume attachment failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume attachment updated result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeAttachmentPortal) DeleteVolumeAttachment() {
	var attachment = model.VolumeAttachmentSpec{
		BaseModel: &model.BaseModel{},
	}
	id := this.Ctx.Input.Param(":attachmentId")

	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&attachment); err != nil {
		reason := fmt.Sprintf("Parse volume attachment request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}
	attachment.Id = id

	// Call global controller variable to handle delete volume attachment request.
	result := controller.Brain.DeleteVolumeAttachment(&attachment)
	if result.Status != "Success" {
		reason := fmt.Sprintf("Delete volume attachment failed: %s", result.GetError())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume attachment deleted result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusAccepted)
	this.Ctx.Output.Body(body)
	return
}

type VolumeSnapshotPortal struct {
	beego.Controller
}

func (this *VolumeSnapshotPortal) CreateVolumeSnapshot() {
	var snapshot = model.VolumeSnapshotSpec{
		BaseModel: &model.BaseModel{},
	}

	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&snapshot); err != nil {
		reason := fmt.Sprintf("Parse volume snapshot request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Call global controller variable to handle create volume snapshot request.
	result, err := controller.Brain.CreateVolumeSnapshot(&snapshot)
	if err != nil {
		reason := fmt.Sprintf("Create volume snapshot failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume snapshot created result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusAccepted)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeSnapshotPortal) ListVolumeSnapshots() {
	result, err := db.C.ListVolumeSnapshots()
	if err != nil {
		reason := fmt.Sprintf("List volume snapshots failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume snapshots listed result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeSnapshotPortal) GetVolumeSnapshot() {
	id := this.Ctx.Input.Param(":snapshotId")

	result, err := db.C.GetVolumeSnapshot(id)
	if err != nil {
		reason := fmt.Sprintf("Get volume snapshot failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume snapshot showed result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusOK)
	this.Ctx.Output.Body(body)
	return
}

func (this *VolumeSnapshotPortal) UpdateVolumeSnapshot() {
	this.Ctx.Output.SetStatus(StatusNotImplemented)
}

func (this *VolumeSnapshotPortal) DeleteVolumeSnapshot() {
	var snapshot = model.VolumeSnapshotSpec{
		BaseModel: &model.BaseModel{},
	}
	id := this.Ctx.Input.Param(":snapshotId")

	if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&snapshot); err != nil {
		reason := fmt.Sprintf("Parse volume snapshot request body failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusInternalServerError)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}
	snapshot.Id = id

	// Call global controller variable to handle delete volume snapshot request.
	result := controller.Brain.DeleteVolumeSnapshot(&snapshot)
	if result.Status != "Success" {
		reason := fmt.Sprintf("Delete volume snapshot failed: %s", result.GetError())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	// Marshal the result.
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Marshal volume snapshot deleted result failed: %s", err.Error())
		this.Ctx.Output.SetStatus(StatusBadRequest)
		this.Ctx.Output.Body(utils.ErrorStatus(this.Ctx.Output.Status, reason))
		log.Error(reason)
		return
	}

	this.Ctx.Output.SetStatus(StatusAccepted)
	this.Ctx.Output.Body(body)
	return
}
