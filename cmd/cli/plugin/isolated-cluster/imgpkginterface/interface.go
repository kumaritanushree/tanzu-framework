// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package imgpkginterface ImgPkgClient defines functions to pull/push/List images
package imgpkginterface

type ImgpkgClient interface {
	CopyImageFromTar(sourceImageName string, destImageRepo string, customImageRepoCertificate string, insecureconnection bool) error
	CopyImageToTar(sourceImageName string, destImageRepo string, customImageRepoCertificate string, insecureconnection bool) error
	PullImage(sourceImageName string, destDir string) error
	GetImageTagList(sourceImageName string) []string
}
