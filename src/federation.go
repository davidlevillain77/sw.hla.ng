// ****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2025 MASA Group
//
// ****************************************************************************

package hla

// Service for federation management
type Federation interface {
	ConnectAndJoin(connectionString string, federateName string, federationName string, fomFile string, siteId uint16, applicationId uint16) bool

	PublishObjectClass(className string, attributesName []string)

	SubscribeObjectClass(lassName string, attributesName []string)

	PublishInteractionClass(className string)

	SubscribeInteractionClass(className string)
}
